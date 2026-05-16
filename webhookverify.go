// Hand-maintained companion to the Stainless-generated SDK.
//
// Webhook signature verification is not currently part of the generated
// surface — every PropRaven SDK ships a constant-time verifier so customers
// can authenticate inbound deliveries without rolling crypto themselves.

package propraven

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// VerifyWebhookOptions configures [VerifyWebhook]. MaxClockSkew defaults to
// 5 minutes — matching the PropRaven dispatch retry window.
type VerifyWebhookOptions struct {
	MaxClockSkew time.Duration
}

// VerifyWebhook authenticates an inbound webhook payload.
//
//	signatureHeader: raw value of the X-PropRaven-Signature header, formatted
//	    as "t=<unix_ms>,v1=<hex_hmac>".
//	rawBody: unmodified request body bytes.
//	secret: the webhook secret returned at endpoint creation time
//	    (CreatedWebhookSecret on V1WebhookNewEndpointResponse).
//
// Returns nil on success and an error describing the mismatch on failure.
// Performs constant-time signature comparison and enforces a 5-minute replay
// window by default. Override via [VerifyWebhookOptions].
func VerifyWebhook(signatureHeader string, rawBody []byte, secret string, opts ...VerifyWebhookOptions) error {
	if secret == "" {
		return errors.New("propraven: empty webhook secret")
	}
	t, sig, err := parseWebhookSignatureHeader(signatureHeader)
	if err != nil {
		return err
	}

	o := VerifyWebhookOptions{MaxClockSkew: 5 * time.Minute}
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

func parseWebhookSignatureHeader(h string) (timestampMs int64, sigHex string, err error) {
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
			tt, perr := strconv.ParseInt(v, 10, 64)
			if perr != nil {
				return 0, "", fmt.Errorf("propraven: bad t: %w", perr)
			}
			timestampMs = tt
		case "v1":
			sigHex = v
		}
	}
	if timestampMs == 0 || sigHex == "" {
		return 0, "", errors.New("propraven: signature header missing t or v1")
	}
	return timestampMs, sigHex, nil
}
