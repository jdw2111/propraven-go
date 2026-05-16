package propraven

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// WebhooksService binds to /v1/webhooks.
type WebhooksService struct {
	client *Client
}

// CreateWebhookRequest is the body for [WebhooksService.Create].
type CreateWebhookRequest struct {
	URL         string   `json:"url"`
	EventTypes  []string `json:"event_types"`
	FilterKind  string   `json:"filter_kind"`
	FilterValue any      `json:"filter_value"`
}

// CreatedWebhook is the response from [WebhooksService.Create]. The Secret is
// returned exactly once and must be stored by the caller — there is no way to
// retrieve it later. Use [VerifyWebhook] with this Secret to authenticate
// incoming payloads.
type CreatedWebhook struct {
	Webhook
	Secret string `json:"secret"`
}

// Create registers a new outbound webhook endpoint. The returned [CreatedWebhook.Secret]
// is shown once — persist it before discarding the response.
func (s *WebhooksService) Create(ctx context.Context, req CreateWebhookRequest) (*CreatedWebhook, error) {
	var w CreatedWebhook
	if err := s.client.do(ctx, "POST", "/v1/webhooks", nil, req, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

// List returns all webhook endpoints owned by the calling API key.
func (s *WebhooksService) List(ctx context.Context) ([]Webhook, error) {
	var page struct {
		Data []Webhook `json:"data"`
	}
	if err := s.client.do(ctx, "GET", "/v1/webhooks", nil, nil, &page); err != nil {
		return nil, err
	}
	return page.Data, nil
}

// Delete removes a webhook endpoint by id. Already-queued deliveries continue
// to retry until they succeed or exhaust their retry budget.
func (s *WebhooksService) Delete(ctx context.Context, id string) error {
	return s.client.do(ctx, "DELETE", "/v1/webhooks/"+id, nil, nil, nil)
}

// VerifyOptions configures [VerifyWebhook]. Defaults are sensible — set
// MaxClockSkew only if your clock is known to drift significantly from PropRaven's.
type VerifyOptions struct {
	// MaxClockSkew is the maximum allowed difference between the signature
	// timestamp and the verifier's wall clock. Defaults to 5 minutes (matching
	// the PropRaven dispatch retry window).
	MaxClockSkew time.Duration
}

// VerifyWebhook authenticates an inbound webhook payload. signatureHeader is the
// raw value of the X-PropRaven-Signature header; rawBody is the unmodified
// request body bytes. Returns nil on success, an error describing the mismatch
// on failure. Performs constant-time signature comparison.
//
// Expected header shape: t=<unix_ms>,v1=<hex_hmac>. HMAC is over
// "<unix_ms>.<rawBody>" with the secret returned by [WebhooksService.Create].
func VerifyWebhook(signatureHeader string, rawBody []byte, secret string, opts ...VerifyOptions) error {
	if secret == "" {
		return errors.New("propraven: empty webhook secret")
	}
	t, sig, err := parseSignatureHeader(signatureHeader)
	if err != nil {
		return err
	}

	o := VerifyOptions{MaxClockSkew: 5 * time.Minute}
	if len(opts) > 0 && opts[0].MaxClockSkew > 0 {
		o.MaxClockSkew = opts[0].MaxClockSkew
	}
	skew := time.Since(time.UnixMilli(t)).Abs()
	if skew > o.MaxClockSkew {
		return fmt.Errorf("propraven: signature timestamp %v outside MaxClockSkew %v", skew, o.MaxClockSkew)
	}

	mac := hmac.New(sha256.New, []byte(secret))
	fmt.Fprintf(mac, "%d.", t)
	mac.Write(rawBody)
	expected := mac.Sum(nil)

	got, err := hex.DecodeString(sig)
	if err != nil {
		return fmt.Errorf("propraven: signature is not hex: %w", err)
	}
	if !hmac.Equal(got, expected) {
		return errors.New("propraven: signature mismatch")
	}
	return nil
}

func parseSignatureHeader(h string) (timestampMs int64, sigHex string, err error) {
	h = strings.TrimSpace(h)
	if h == "" {
		return 0, "", errors.New("propraven: empty X-PropRaven-Signature header")
	}
	for _, part := range strings.Split(h, ",") {
		k, v, ok := strings.Cut(strings.TrimSpace(part), "=")
		if !ok {
			continue
		}
		switch k {
		case "t":
			t, perr := strconv.ParseInt(v, 10, 64)
			if perr != nil {
				return 0, "", fmt.Errorf("propraven: bad t: %w", perr)
			}
			timestampMs = t
		case "v1":
			sigHex = v
		}
	}
	if timestampMs == 0 || sigHex == "" {
		return 0, "", errors.New("propraven: signature header missing t or v1")
	}
	return timestampMs, sigHex, nil
}
