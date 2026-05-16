package propraven

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"time"
)

// helper to construct a valid signature for a body/secret/timestamp combo.
func sign(t *testing.T, body []byte, secret string, ts int64) string {
	t.Helper()
	mac := hmac.New(sha256.New, []byte(secret))
	fmt.Fprintf(mac, "%d.", ts)
	mac.Write(body)
	return fmt.Sprintf("t=%d,v1=%s", ts, hex.EncodeToString(mac.Sum(nil)))
}

func TestVerifyWebhook_Valid(t *testing.T) {
	secret := "whsec_test"
	body := []byte(`{"event_type":"parcel.sold"}`)
	ts := time.Now().UnixMilli()
	header := sign(t, body, secret, ts)

	if err := VerifyWebhook(header, body, secret); err != nil {
		t.Fatalf("expected valid signature, got error: %v", err)
	}
}

func TestVerifyWebhook_WrongSecret(t *testing.T) {
	body := []byte(`{}`)
	ts := time.Now().UnixMilli()
	header := sign(t, body, "right-secret", ts)

	err := VerifyWebhook(header, body, "wrong-secret")
	if err == nil || !strings.Contains(err.Error(), "mismatch") {
		t.Fatalf("expected signature mismatch, got: %v", err)
	}
}

func TestVerifyWebhook_TamperedBody(t *testing.T) {
	secret := "whsec_test"
	body := []byte(`{"original":true}`)
	ts := time.Now().UnixMilli()
	header := sign(t, body, secret, ts)

	tampered := []byte(`{"original":false}`)
	err := VerifyWebhook(header, tampered, secret)
	if err == nil {
		t.Fatal("expected tampered body to fail verification")
	}
}

func TestVerifyWebhook_StaleTimestamp(t *testing.T) {
	secret := "whsec_test"
	body := []byte(`{}`)
	stale := time.Now().Add(-10 * time.Minute).UnixMilli()
	header := sign(t, body, secret, stale)

	err := VerifyWebhook(header, body, secret)
	if err == nil || !strings.Contains(err.Error(), "MaxClockSkew") {
		t.Fatalf("expected clock-skew rejection, got: %v", err)
	}
}

func TestVerifyWebhook_CustomSkew(t *testing.T) {
	secret := "whsec_test"
	body := []byte(`{}`)
	stale := time.Now().Add(-10 * time.Minute).UnixMilli()
	header := sign(t, body, secret, stale)

	// Allow 15 minutes — should pass.
	err := VerifyWebhook(header, body, secret, VerifyOptions{MaxClockSkew: 15 * time.Minute})
	if err != nil {
		t.Fatalf("expected pass with widened skew, got: %v", err)
	}
}

func TestVerifyWebhook_MalformedHeader(t *testing.T) {
	cases := []string{
		"",
		"garbage",
		"t=12345",
		"v1=abc",
		"t=,v1=abc",
		"t=12345,v1=",
		"t=notanumber,v1=abc",
	}
	for _, h := range cases {
		t.Run(h, func(t *testing.T) {
			if err := VerifyWebhook(h, []byte("{}"), "secret"); err == nil {
				t.Fatalf("expected error for header %q", h)
			}
		})
	}
}

func TestParseSignatureHeader_OrderInsensitive(t *testing.T) {
	// PropRaven canonically emits t=...,v1=... in that order, but verifiers
	// should accept either ordering. Belt + suspenders.
	header := "v1=" + strings.Repeat("a", 64) + ",t=" + fmt.Sprintf("%d", time.Now().UnixMilli())
	_, _, err := parseSignatureHeader(header)
	if err != nil {
		t.Fatalf("expected order-insensitive parsing, got: %v", err)
	}
}
