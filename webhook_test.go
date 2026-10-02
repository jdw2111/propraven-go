package propraven

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"testing"
	"time"
)

const (
	whSecret = "whsec_test"
	whBody   = `{"type":"parcel.sold","id":"evt_1"}`
	whTS     = int64(1700000000000)
	// HMAC_SHA256("whsec_test", "1700000000000.{...}") from Python's hmac module.
	whHex = "5b5571cbfa4a4d4a2671f01bba9d4e2e60a3738433375ff005f2377d1f26a874"
)

func whNow() WebhookOption { return WithWebhookNow(time.UnixMilli(whTS)) }

func TestWebhookReferenceVector(t *testing.T) {
	mac := hmac.New(sha256.New, []byte(whSecret))
	mac.Write([]byte("1700000000000." + whBody))
	if got := hex.EncodeToString(mac.Sum(nil)); got != whHex {
		t.Fatalf("computed %s, reference %s", got, whHex)
	}
	if got := SignWebhook([]byte(whBody), whSecret, time.UnixMilli(whTS)); got != "t=1700000000000,v1="+whHex {
		t.Fatalf("SignWebhook = %s", got)
	}
}

func TestVerifyWebhookValid(t *testing.T) {
	ev, err := VerifyWebhook([]byte(whBody), "t=1700000000000,v1="+whHex, whSecret, whNow())
	if err != nil {
		t.Fatal(err)
	}
	if string(ev) != whBody {
		t.Fatalf("event = %s", ev)
	}
	// Within tolerance on either side.
	if _, err := VerifyWebhook([]byte(whBody), "t=1700000000000,v1="+whHex, whSecret,
		WithWebhookNow(time.UnixMilli(whTS+299_000))); err != nil {
		t.Fatal(err)
	}
}

func TestVerifyWebhookBadSignature(t *testing.T) {
	bad := strings.Repeat("0", 64)
	_, err := VerifyWebhook([]byte(whBody), "t=1700000000000,v1="+bad, whSecret, whNow())
	if !errors.Is(err, ErrWebhookSignature) {
		t.Fatalf("err = %v", err)
	}
	// Tampered body.
	_, err = VerifyWebhook([]byte(`{"type":"parcel.sold","id":"evt_2"}`), "t=1700000000000,v1="+whHex, whSecret, whNow())
	if !errors.Is(err, ErrWebhookSignature) {
		t.Fatalf("tampered body err = %v", err)
	}
	// Wrong secret, and a secret that is hex-decoded or stripped is wrong too.
	for _, s := range []string{"whsec_other", "test", ""} {
		if _, err := VerifyWebhook([]byte(whBody), "t=1700000000000,v1="+whHex, s, whNow()); !errors.Is(err, ErrWebhookSignature) {
			t.Fatalf("secret %q err = %v", s, err)
		}
	}
}

func TestVerifyWebhookExpired(t *testing.T) {
	for _, delta := range []int64{301_000, -301_000} {
		_, err := VerifyWebhook([]byte(whBody), "t=1700000000000,v1="+whHex, whSecret,
			WithWebhookNow(time.UnixMilli(whTS+delta)))
		if !errors.Is(err, ErrWebhookSignature) || !strings.Contains(err.Error(), "tolerance") {
			t.Fatalf("delta %d err = %v", delta, err)
		}
	}
	// A wider tolerance accepts it.
	if _, err := VerifyWebhook([]byte(whBody), "t=1700000000000,v1="+whHex, whSecret,
		WithWebhookNow(time.UnixMilli(whTS+600_000)), WithWebhookTolerance(15*time.Minute)); err != nil {
		t.Fatal(err)
	}
}

func TestVerifyWebhookMalformedHeader(t *testing.T) {
	for _, h := range []string{"", "garbage", "t=abc,v1=" + whHex, "v1=" + whHex, "t=1700000000000", "t=1700000000000,v1="} {
		if _, err := VerifyWebhook([]byte(whBody), h, whSecret, whNow()); !errors.Is(err, ErrWebhookSignature) {
			t.Fatalf("header %q err = %v", h, err)
		}
	}
}

func TestVerifyWebhookMultipleV1(t *testing.T) {
	old := strings.Repeat("ab", 32)
	h := "t=1700000000000,v1=" + old + ",v1=" + whHex
	if _, err := VerifyWebhook([]byte(whBody), h, whSecret, whNow()); err != nil {
		t.Fatalf("rotation header rejected: %v", err)
	}
	h = "t=1700000000000, v1=" + whHex + ", v1=nothex"
	if _, err := VerifyWebhook([]byte(whBody), h, whSecret, whNow()); err != nil {
		t.Fatalf("spaced header rejected: %v", err)
	}
}
