package propraven

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Error is the typed error returned from any SDK call that received a non-2xx
// HTTP response. Use [errors.As] to extract:
//
//	var pe *propraven.Error
//	if errors.As(err, &pe) && pe.Status == 404 { ... }
type Error struct {
	Status     int    // HTTP status code
	Code       string // PropRaven application error code if surfaced (e.g. "rate_limited")
	Message    string // Human-readable error message
	RequestID  string // X-Request-Id header value; quote in support tickets
	RawBody    []byte // Raw response body, for diagnostics
}

func (e *Error) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("propraven: %d %s: %s", e.Status, e.Code, e.Message)
	}
	return fmt.Sprintf("propraven: %d %s", e.Status, e.Message)
}

// IsRateLimited reports whether the error is a 429 rate-limit response. Customers
// should back off and retry; rate-limit ceilings are documented per pricing tier.
func IsRateLimited(err error) bool {
	var pe *Error
	if errors.As(err, &pe) {
		return pe.Status == http.StatusTooManyRequests
	}
	return false
}

// IsNotFound reports whether the error is a 404. Common when a parcel id was
// typo'd or the caller's API key doesn't have access to the parcel's state.
func IsNotFound(err error) bool {
	var pe *Error
	if errors.As(err, &pe) {
		return pe.Status == http.StatusNotFound
	}
	return false
}

// decodeError builds an Error from an http.Response and its raw body. Used
// internally by Client.do.
func decodeError(resp *http.Response, body []byte) error {
	e := &Error{
		Status:    resp.StatusCode,
		RawBody:   body,
		RequestID: resp.Header.Get("X-Request-Id"),
		Message:   http.StatusText(resp.StatusCode),
	}
	var wire struct {
		Error  string `json:"error"`
		Detail string `json:"detail"`
		Code   string `json:"code"`
	}
	if err := json.Unmarshal(body, &wire); err == nil {
		if wire.Error != "" {
			e.Message = wire.Error
		}
		if wire.Detail != "" && e.Message != wire.Detail {
			e.Message = e.Message + ": " + wire.Detail
		}
		if wire.Code != "" {
			e.Code = wire.Code
		}
	}
	return e
}
