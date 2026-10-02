package propraven

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// WebhookSignatureHeader is the header carrying a delivery's signature.
const WebhookSignatureHeader = "X-PropRaven-Signature"

// DefaultWebhookTolerance is the default allowed clock difference between
// the signature timestamp and now.
const DefaultWebhookTolerance = 300 * time.Second

// ErrWebhookSignature is wrapped by every error [VerifyWebhook] returns, so
// errors.Is(err, propraven.ErrWebhookSignature) identifies a rejected
// delivery.
var ErrWebhookSignature = errors.New("propraven: webhook signature verification failed")

type webhookConfig struct {
	tolerance time.Duration
	now       func() time.Time
}

// WebhookOption configures [VerifyWebhook].
type WebhookOption func(*webhookConfig)

// WithWebhookTolerance sets the allowed |now - t| (default 300s). Zero or a
// negative value disables the timestamp check.
func WithWebhookTolerance(d time.Duration) WebhookOption {
	return func(c *webhookConfig) { c.tolerance = d }
}

// WithWebhookNow overrides the clock (for tests).
func WithWebhookNow(now time.Time) WebhookOption {
	return func(c *webhookConfig) { c.now = func() time.Time { return now } }
}

// VerifyWebhook authenticates a webhook delivery and returns its JSON event.
//
//   - payload: the raw request body, unmodified.
//   - signature: the X-PropRaven-Signature header, "t=<unix_ms>,v1=<hex>".
//     Several v1= entries are accepted (secret rotation).
//   - secret: the webhook secret exactly as issued (starts with "whsec_").
//
// The signed message is "<t>.<payload>" with HMAC-SHA256 keyed by the secret
// string; comparison is constant time. Deliveries whose timestamp is more
// than the tolerance (default 300s) from now are rejected.
//
// Every failure wraps [ErrWebhookSignature].
func VerifyWebhook(payload []byte, signature, secret string, opts ...WebhookOption) (json.RawMessage, error) {
	cfg := webhookConfig{tolerance: DefaultWebhookTolerance, now: time.Now}
	for _, o := range opts {
		if o != nil {
			o(&cfg)
		}
	}
	if secret == "" {
		return nil, fmt.Errorf("%w: empty webhook secret", ErrWebhookSignature)
	}
	ts, sigs, err := parseSignatureHeader(signature)
	if err != nil {
		return nil, err
	}
	if cfg.tolerance > 0 {
		diff := cfg.now().UnixMilli() - ts
		if diff < 0 {
			diff = -diff
		}
		if diff > cfg.tolerance.Milliseconds() {
			return nil, fmt.Errorf("%w: timestamp outside the %s tolerance", ErrWebhookSignature, cfg.tolerance)
		}
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(strconv.FormatInt(ts, 10)))
	mac.Write([]byte("."))
	mac.Write(payload)
	expected := mac.Sum(nil)
	ok := false
	for _, s := range sigs {
		got, err := hex.DecodeString(s)
		if err != nil {
			continue
		}
		if hmac.Equal(got, expected) {
			ok = true
		}
	}
	if !ok {
		return nil, fmt.Errorf("%w: no matching v1 signature", ErrWebhookSignature)
	}
	if !json.Valid(payload) {
		return nil, fmt.Errorf("%w: payload is not valid JSON", ErrWebhookSignature)
	}
	return json.RawMessage(append([]byte(nil), payload...)), nil
}

// SignWebhook computes the X-PropRaven-Signature value for payload at time
// t, matching the server's signer. Useful for testing your handler.
func SignWebhook(payload []byte, secret string, t time.Time) string {
	ts := t.UnixMilli()
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(strconv.FormatInt(ts, 10)))
	mac.Write([]byte("."))
	mac.Write(payload)
	return "t=" + strconv.FormatInt(ts, 10) + ",v1=" + hex.EncodeToString(mac.Sum(nil))
}

func parseSignatureHeader(h string) (int64, []string, error) {
	h = strings.TrimSpace(h)
	if h == "" {
		return 0, nil, fmt.Errorf("%w: missing %s header", ErrWebhookSignature, WebhookSignatureHeader)
	}
	var (
		ts    int64
		hasTS bool
		sigs  []string
	)
	for _, part := range strings.Split(h, ",") {
		k, v, ok := strings.Cut(strings.TrimSpace(part), "=")
		if !ok {
			return 0, nil, fmt.Errorf("%w: malformed signature header", ErrWebhookSignature)
		}
		switch strings.TrimSpace(k) {
		case "t":
			n, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
			if err != nil || n <= 0 {
				return 0, nil, fmt.Errorf("%w: malformed timestamp", ErrWebhookSignature)
			}
			ts, hasTS = n, true
		case "v1":
			if s := strings.TrimSpace(v); s != "" {
				sigs = append(sigs, s)
			}
		}
	}
	if !hasTS {
		return 0, nil, fmt.Errorf("%w: signature header has no t=", ErrWebhookSignature)
	}
	if len(sigs) == 0 {
		return 0, nil, fmt.Errorf("%w: signature header has no v1=", ErrWebhookSignature)
	}
	return ts, sigs, nil
}
