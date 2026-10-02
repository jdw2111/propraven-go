package propraven

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Error is returned for every non-2xx API response. Inspect it with
// errors.As:
//
//	var apiErr *propraven.Error
//	if errors.As(err, &apiErr) {
//		log.Println(apiErr.StatusCode, apiErr.Code, apiErr.RequestID)
//	}
//
// or with the predicates [IsNotFound], [IsRateLimited], [IsPaymentRequired],
// [IsAuthentication] and friends. The body is parsed from RFC 7807
// problem+json, an x402 payment-required envelope, or a legacy
// {"error": "..."} body.
type Error struct {
	// StatusCode is the HTTP status.
	StatusCode int
	// Type is the problem type URI.
	Type string
	// Title is the short problem title.
	Title string
	// Detail is the human-readable explanation.
	Detail string
	// Code is the stable machine-readable code (invalid_parameter,
	// not_found, monthly_cap_reached, ...).
	Code string
	// Errors lists per-parameter validation failures (400s).
	Errors []FieldError
	// RequestID identifies the request for support.
	RequestID string
	// RetryAfter is the server's requested wait in seconds (Retry-After
	// header or the body's retry_after), or nil.
	RetryAfter *float64
	// Accepts holds the x402 payment requirements of a 402 (empty otherwise).
	Accepts []map[string]any
	// Header is the response header.
	Header http.Header
	// Body is the raw response body.
	Body []byte
}

// FieldError is one per-parameter validation failure.
type FieldError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

// Error formats as "<status> <code>: <detail>".
func (e *Error) Error() string {
	detail := e.Detail
	if detail == "" {
		detail = e.Title
	}
	if detail == "" {
		detail = http.StatusText(e.StatusCode)
	}
	if e.Code == "" {
		return fmt.Sprintf("%d: %s", e.StatusCode, detail)
	}
	return fmt.Sprintf("%d %s: %s", e.StatusCode, e.Code, detail)
}

// Class names the error category shared by every PropRaven SDK
// (BadRequestError, AuthenticationError, PaymentRequiredError, ...).
func (e *Error) Class() string {
	switch e.StatusCode {
	case 400:
		return "BadRequestError"
	case 401:
		return "AuthenticationError"
	case 402:
		return "PaymentRequiredError"
	case 403:
		return "PermissionDeniedError"
	case 404:
		return "NotFoundError"
	case 405:
		return "MethodNotAllowedError"
	case 409:
		return "ConflictError"
	case 413:
		return "PayloadTooLargeError"
	case 422:
		return "UnprocessableEntityError"
	case 429:
		return "RateLimitError"
	case 503:
		return "ServiceUnavailableError"
	case 504:
		return "GatewayTimeoutError"
	}
	if e.StatusCode >= 500 {
		return "InternalServerError"
	}
	return "APIError"
}

func newAPIError(status int, h http.Header, body []byte, now time.Time) *Error {
	e := &Error{StatusCode: status, Header: h, Body: body, Accepts: []map[string]any{}}
	var raw map[string]json.RawMessage
	if json.Unmarshal(body, &raw) == nil && raw != nil {
		str := func(k string) string {
			var s string
			if v, ok := raw[k]; ok && json.Unmarshal(v, &s) == nil {
				return s
			}
			return ""
		}
		e.Type = str("type")
		e.Title = str("title")
		e.Detail = str("detail")
		e.Code = str("code")
		e.RequestID = str("request_id")
		if e.RequestID == "" {
			e.RequestID = str("requestId")
		}
		if v, ok := raw["errors"]; ok {
			_ = json.Unmarshal(v, &e.Errors)
		}
		if v, ok := raw["accepts"]; ok {
			var acc []map[string]any
			if json.Unmarshal(v, &acc) == nil && acc != nil {
				e.Accepts = acc
			}
		}
		// Legacy {"error": "..."} and x402 {"x402Version", "error", "accepts"}.
		if v, ok := raw["error"]; ok {
			var s string
			if json.Unmarshal(v, &s) == nil {
				if e.Detail == "" {
					e.Detail = s
				}
			} else {
				// {"error": {"code": ..., "message": ...}}
				var obj struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				}
				if json.Unmarshal(v, &obj) == nil {
					if e.Code == "" {
						e.Code = obj.Code
					}
					if e.Detail == "" {
						e.Detail = obj.Message
					}
				}
			}
		}
		if e.Code == "" {
			if _, ok := raw["x402Version"]; ok {
				e.Code = "payment_required"
			}
		}
		if v, ok := raw["retry_after"]; ok {
			var f float64
			if json.Unmarshal(v, &f) == nil {
				e.RetryAfter = &f
			}
		}
	} else if s := strings.TrimSpace(string(body)); s != "" && len(s) < 500 {
		e.Detail = s
	}
	if e.RequestID == "" {
		e.RequestID = firstHeader(h, "X-Request-Id", "X-Vercel-Id")
	}
	if secs, ok := parseRetryAfter(h.Get("Retry-After"), now); ok {
		e.RetryAfter = &secs
	}
	return e
}

func firstHeader(h http.Header, names ...string) string {
	for _, n := range names {
		if v := h.Get(n); v != "" {
			return v
		}
	}
	return ""
}

// ConnectionError is returned when no HTTP response was received (DNS,
// connection refused, TLS, timeout, cancelled context). It wraps the cause,
// so errors.Is(err, context.DeadlineExceeded) works.
type ConnectionError struct {
	Err error
	// Timeout is true when the request timed out (APITimeoutError in the
	// other SDKs).
	Timeout bool
}

func (e *ConnectionError) Error() string {
	if e.Timeout {
		return "propraven: request timed out: " + errString(e.Err)
	}
	return "propraven: connection error: " + errString(e.Err)
}

func (e *ConnectionError) Unwrap() error { return e.Err }

func errString(err error) string {
	if err == nil {
		return "unknown"
	}
	return err.Error()
}

func statusOf(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.StatusCode
	}
	return 0
}

// AsError returns the *Error inside err, or nil.
func AsError(err error) *Error {
	var e *Error
	if errors.As(err, &e) {
		return e
	}
	return nil
}

// IsBadRequest reports a 400.
func IsBadRequest(err error) bool { return statusOf(err) == 400 }

// IsAuthentication reports a 401 (missing or invalid API key).
func IsAuthentication(err error) bool { return statusOf(err) == 401 }

// IsPaymentRequired reports a 402 (x402 payment, plan limit or monthly cap).
func IsPaymentRequired(err error) bool { return statusOf(err) == 402 }

// IsPermissionDenied reports a 403.
func IsPermissionDenied(err error) bool { return statusOf(err) == 403 }

// IsNotFound reports a 404.
func IsNotFound(err error) bool { return statusOf(err) == 404 }

// IsConflict reports a 409.
func IsConflict(err error) bool { return statusOf(err) == 409 }

// IsRateLimited reports a 429.
func IsRateLimited(err error) bool { return statusOf(err) == 429 }

// IsServerError reports any 5xx.
func IsServerError(err error) bool { return statusOf(err) >= 500 }

// IsConnectionError reports that no response was received.
func IsConnectionError(err error) bool {
	var ce *ConnectionError
	return errors.As(err, &ce)
}

// IsTimeout reports that the request timed out.
func IsTimeout(err error) bool {
	var ce *ConnectionError
	return errors.As(err, &ce) && ce.Timeout
}
