package propraven

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestErrorClassesFromProblemBody(t *testing.T) {
	cases := []struct {
		status int
		class  string
		pred   func(error) bool
	}{
		{400, "BadRequestError", IsBadRequest},
		{401, "AuthenticationError", IsAuthentication},
		{402, "PaymentRequiredError", IsPaymentRequired},
		{403, "PermissionDeniedError", IsPermissionDenied},
		{404, "NotFoundError", IsNotFound},
		{405, "MethodNotAllowedError", nil},
		{409, "ConflictError", IsConflict},
		{413, "PayloadTooLargeError", nil},
		{422, "UnprocessableEntityError", nil},
		{429, "RateLimitError", IsRateLimited},
		{500, "InternalServerError", IsServerError},
		{502, "InternalServerError", IsServerError},
		{503, "ServiceUnavailableError", IsServerError},
		{504, "GatewayTimeoutError", IsServerError},
		{418, "APIError", nil},
	}
	for _, tc := range cases {
		body := `{"type":"https://api.propraven.com/errors/x","title":"T","status":` + itoa(tc.status) +
			`,"detail":"limit: must be >= 1","code":"invalid_parameter","errors":[{"param":"limit","message":"must be >= 1"}],"request_id":"iad1::abc"}`
		env := newEnv(t, jsonReply(tc.status, body, "Content-Type", "application/problem+json"), WithMaxRetries(0))
		_, err := env.client.Parcels.Get(context.Background(), "p1", nil)
		var e *Error
		if !errors.As(err, &e) {
			t.Fatalf("%d: err = %T %v", tc.status, err, err)
		}
		if e.StatusCode != tc.status || e.Class() != tc.class {
			t.Errorf("%d: status %d class %s, want %s", tc.status, e.StatusCode, e.Class(), tc.class)
		}
		if e.Code != "invalid_parameter" || e.Detail != "limit: must be >= 1" || e.Title != "T" ||
			e.Type != "https://api.propraven.com/errors/x" || e.RequestID != "iad1::abc" {
			t.Errorf("%d: parsed %+v", tc.status, e)
		}
		if len(e.Errors) != 1 || e.Errors[0].Param != "limit" || e.Errors[0].Message != "must be >= 1" {
			t.Errorf("%d: errors %+v", tc.status, e.Errors)
		}
		if want := itoa(tc.status) + " invalid_parameter: limit: must be >= 1"; e.Error() != want {
			t.Errorf("%d: message %q, want %q", tc.status, e.Error(), want)
		}
		if tc.pred != nil && !tc.pred(err) {
			t.Errorf("%d: predicate false", tc.status)
		}
		if e.Header == nil || len(e.Body) == 0 {
			t.Errorf("%d: header/body not kept", tc.status)
		}
		if AsError(err) != e {
			t.Errorf("%d: AsError mismatch", tc.status)
		}
	}
}

func TestErrorFromLegacyBody(t *testing.T) {
	env := newEnv(t, jsonReply(404, `{"error":"Parcel not found"}`))
	_, err := env.client.Parcels.Get(context.Background(), "nope", nil)
	if !IsNotFound(err) {
		t.Fatalf("err = %v", err)
	}
	e := AsError(err)
	if e.Detail != "Parcel not found" || e.Error() != "404: Parcel not found" || e.Class() != "NotFoundError" {
		t.Fatalf("parsed %+v / %q", e, e.Error())
	}
	// Non-JSON bodies keep the text as the detail.
	env = newEnv(t, jsonReply(400, "bad things", "Content-Type", "text/plain"))
	_, err = env.client.Parcels.Get(context.Background(), "x", nil)
	if e := AsError(err); e == nil || e.Detail != "bad things" || !IsBadRequest(err) {
		t.Fatalf("err = %v", err)
	}
}

func TestErrorFromX402Body(t *testing.T) {
	body := `{"x402Version":1,"error":"X-PAYMENT header is required","accepts":[{"scheme":"exact","network":"base","maxAmountRequired":"5000000","resource":"https://propraven.com/api/v1/parcels/p1/report"}]}`
	env := newEnv(t, jsonReply(402, body))
	_, err := env.client.Parcels.Report(context.Background(), "p1", nil)
	if !IsPaymentRequired(err) {
		t.Fatalf("err = %v", err)
	}
	e := AsError(err)
	if e.Class() != "PaymentRequiredError" || e.Detail != "X-PAYMENT header is required" || e.Code != "payment_required" {
		t.Fatalf("parsed %+v", e)
	}
	if len(e.Accepts) != 1 || e.Accepts[0]["maxAmountRequired"] != "5000000" {
		t.Fatalf("accepts %+v", e.Accepts)
	}
	if env.rec.count() != 1 {
		t.Fatalf("402 must not be retried; attempts = %d", env.rec.count())
	}
	// Problem-shaped 402s have an empty (non-nil) Accepts.
	env = newEnv(t, jsonReply(402, `{"type":"t","title":"Payment Required","status":402,"detail":"Monthly cap reached","code":"monthly_cap_reached"}`, "Retry-After", "86400"))
	_, err = env.client.Deals.Absentee(context.Background(), nil)
	e = AsError(err)
	if e == nil || e.Accepts == nil || len(e.Accepts) != 0 || e.Code != "monthly_cap_reached" {
		t.Fatalf("parsed %+v", e)
	}
	if e.RetryAfter == nil || *e.RetryAfter != 86400 {
		t.Fatalf("RetryAfter = %v", e.RetryAfter)
	}
	if env.rec.count() != 1 {
		t.Fatalf("attempts = %d", env.rec.count())
	}
}

func TestRetryOn429WithRetryAfter(t *testing.T) {
	env := newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		if attempt == 0 {
			jsonReply(429, `{"type":"t","title":"Too Many","status":429,"detail":"slow down","code":"rate_limit_exceeded"}`, "Retry-After", "3")(w, r, attempt)
			return
		}
		jsonReply(200, `{}`)(w, r, attempt)
	})
	if _, err := env.client.Freshness.Get(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	if env.rec.count() != 2 || len(env.waits) != 1 || env.waits[0] != 3*time.Second {
		t.Fatalf("attempts %d waits %v", env.rec.count(), env.waits)
	}
}

func TestRetryAfterHTTPDateAndRateLimitReset(t *testing.T) {
	now := time.Date(2026, 10, 2, 12, 0, 0, 0, time.UTC)
	env := newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		switch attempt {
		case 0:
			jsonReply(429, `{}`, "Retry-After", now.Add(5*time.Second).Format(http.TimeFormat))(w, r, attempt)
		case 1:
			jsonReply(429, `{}`, "X-RateLimit-Remaining", "0", "X-RateLimit-Reset", itoa(int(now.Unix())+7))(w, r, attempt)
		default:
			jsonReply(200, `{}`)(w, r, attempt)
		}
	})
	env.client.core.now = func() time.Time { return now }
	if _, err := env.client.Freshness.Get(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	if len(env.waits) != 2 || env.waits[0] != 5*time.Second || env.waits[1] != 7*time.Second {
		t.Fatalf("waits = %v", env.waits)
	}
}

func TestRetryOn503ForPOST(t *testing.T) {
	env := newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		if attempt < 2 {
			jsonReply(503, `{"type":"t","title":"Unavailable","status":503,"detail":"warming","code":"service_unavailable"}`)(w, r, attempt)
			return
		}
		jsonReply(200, `{}`)(w, r, attempt)
	})
	if _, err := env.client.Lookup.Batch(context.Background(), &LookupBatchParams{Queries: []string{"x"}}); err != nil {
		t.Fatal(err)
	}
	if env.rec.count() != 3 {
		t.Fatalf("attempts = %d", env.rec.count())
	}
	// Exponential backoff without Retry-After: 0.5s, 1s (jitter pinned to 0).
	if len(env.waits) != 2 || env.waits[0] != 500*time.Millisecond || env.waits[1] != time.Second {
		t.Fatalf("waits = %v", env.waits)
	}
}

func TestBackoffJitterBounds(t *testing.T) {
	core := &clientCore{jitter: func() float64 { return -1 }}
	if d := core.backoff(1); d != 750*time.Millisecond {
		t.Fatalf("low jitter backoff = %v", d)
	}
	core.jitter = func() float64 { return 0.999999 }
	if d := core.backoff(1); d < 1249*time.Millisecond || d > 1250*time.Millisecond {
		t.Fatalf("high jitter backoff = %v", d)
	}
	core.jitter = func() float64 { return 0 }
	if d := core.backoff(20); d != 60*time.Second {
		t.Fatalf("backoff cap = %v", d)
	}
}

func TestNoRetryOn400And402(t *testing.T) {
	for _, status := range []int{400, 401, 402, 404, 409, 422} {
		env := newEnv(t, jsonReply(status, `{"error":"no"}`))
		_, err := env.client.Freshness.Get(context.Background(), nil)
		if AsError(err) == nil || env.rec.count() != 1 {
			t.Fatalf("%d: attempts = %d err = %v", status, env.rec.count(), err)
		}
	}
}

func TestNoRetryOnPOST500ButRetryOnGET500(t *testing.T) {
	env := newEnv(t, jsonReply(500, `{"error":"boom"}`))
	_, err := env.client.Lookup.Batch(context.Background(), &LookupBatchParams{Queries: []string{"x"}})
	if !IsServerError(err) || env.rec.count() != 1 {
		t.Fatalf("POST 500: attempts = %d err = %v", env.rec.count(), err)
	}
	env = newEnv(t, jsonReply(502, `{"error":"bad gateway"}`))
	_, err = env.client.Freshness.Get(context.Background(), nil)
	if !IsServerError(err) || env.rec.count() != 3 {
		t.Fatalf("GET 502: attempts = %d err = %v", env.rec.count(), err)
	}
}

func TestMaxRetriesExhaustion(t *testing.T) {
	env := newEnv(t, jsonReply(429, `{"code":"rate_limit_exceeded","detail":"slow"}`, "Retry-After", "1"), WithMaxRetries(3))
	_, err := env.client.Freshness.Get(context.Background(), nil)
	e := AsError(err)
	if e == nil || !IsRateLimited(err) || env.rec.count() != 4 || len(env.waits) != 3 {
		t.Fatalf("attempts %d waits %v err %v", env.rec.count(), env.waits, err)
	}
	if e.RetryAfter == nil || *e.RetryAfter != 1 {
		t.Fatalf("RetryAfter = %v", e.RetryAfter)
	}
	// Per-request override.
	env = newEnv(t, jsonReply(429, `{}`, "Retry-After", "1"))
	_, _ = env.client.Freshness.Get(context.Background(), nil, WithMaxRetries(0))
	if env.rec.count() != 1 {
		t.Fatalf("WithMaxRetries(0) attempts = %d", env.rec.count())
	}
}

func TestRetryWaitCappedAt60Seconds(t *testing.T) {
	env := newEnv(t, jsonReply(429, `{"code":"rate_limit_exceeded"}`, "Retry-After", "61"))
	_, err := env.client.Freshness.Get(context.Background(), nil)
	e := AsError(err)
	if e == nil || env.rec.count() != 1 || len(env.waits) != 0 {
		t.Fatalf("attempts %d waits %v err %v", env.rec.count(), env.waits, err)
	}
	if e.RetryAfter == nil || *e.RetryAfter != 61 {
		t.Fatalf("RetryAfter = %v", e.RetryAfter)
	}
	// Exactly 60 is still retried.
	env = newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		if attempt == 0 {
			jsonReply(429, `{}`, "Retry-After", "60")(w, r, attempt)
			return
		}
		jsonReply(200, `{}`)(w, r, attempt)
	})
	if _, err := env.client.Freshness.Get(context.Background(), nil); err != nil || env.waits[0] != 60*time.Second {
		t.Fatalf("err %v waits %v", err, env.waits)
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	if neg {
		b = append([]byte{'-'}, b...)
	}
	return string(b)
}
