package propraven

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// apiRequest is one operation call, built by generated code.
type apiRequest struct {
	method string
	path   string // already escaped, starts with /api/v1
	query  url.Values
	header http.Header
	body   any // JSON-encoded when non-nil
	accept string
}

func newRequest(method, path string) *apiRequest {
	return &apiRequest{method: method, path: path, query: url.Values{}, header: http.Header{}, accept: "application/json"}
}

// decodeFunc turns a 2xx response body into the caller's result.
type decodeFunc func(header http.Header, body []byte) error

func decodeJSON(out any) decodeFunc {
	return func(_ http.Header, body []byte) error {
		if len(bytes.TrimSpace(body)) == 0 {
			return nil
		}
		if err := json.Unmarshal(body, out); err != nil {
			return fmt.Errorf("propraven: decoding response: %w", err)
		}
		return nil
	}
}

func decodeText(out *string) decodeFunc {
	return func(_ http.Header, body []byte) error {
		*out = string(body)
		return nil
	}
}

func isJSONContentType(h http.Header) bool {
	mt, _, err := mime.ParseMediaType(h.Get("Content-Type"))
	if err != nil {
		return false
	}
	return mt == "application/json" || strings.HasSuffix(mt, "+json")
}

// do sends req with retries and decodes a 2xx response with decode.
func (c *Client) do(ctx context.Context, req *apiRequest, opts []RequestOption, decode decodeFunc) error {
	if ctx == nil {
		ctx = context.Background()
	}
	core := c.core
	cfg := core.cfg.clone()
	for _, o := range opts {
		if o != nil {
			o(&cfg)
		}
	}

	u := cfg.baseURL + req.path
	if len(req.query) > 0 {
		u += "?" + req.query.Encode()
	}
	if _, err := url.Parse(u); err != nil {
		return fmt.Errorf("propraven: invalid request URL: %w", err)
	}

	var body []byte
	if req.body != nil {
		b, err := json.Marshal(req.body)
		if err != nil {
			return fmt.Errorf("propraven: encoding request body: %w", err)
		}
		body = b
	}

	for attempt := 0; ; attempt++ {
		status, header, respBody, err := core.attempt(ctx, &cfg, req, u, body)
		if err != nil {
			if ctx.Err() != nil {
				return &ConnectionError{Err: err, Timeout: errors.Is(ctx.Err(), context.DeadlineExceeded)}
			}
			cerr := &ConnectionError{Err: err, Timeout: isTimeout(err)}
			if attempt >= cfg.maxRetries {
				return cerr
			}
			if !core.sleep(ctx.Done(), core.backoff(attempt)) {
				return &ConnectionError{Err: ctx.Err(), Timeout: errors.Is(ctx.Err(), context.DeadlineExceeded)}
			}
			continue
		}
		core.recordRateLimit(header)
		if cfg.raw != nil {
			*cfg.raw = RawResponse{StatusCode: status, Header: header, Body: respBody}
		}
		if status >= 200 && status < 300 {
			if decode == nil {
				return nil
			}
			err := decode(header, respBody)
			var te *json.UnmarshalTypeError
			if errors.As(err, &te) {
				// Everything else decoded; keep the result and leave the
				// mismatched field unset rather than failing the call.
				cfg.logger.Printf("warning: %s %s: response field %q is JSON %s, not %s as the spec says; left unset",
					req.method, req.path, te.Field, te.Value, te.Type)
				return nil
			}
			return err
		}
		apiErr := newAPIError(status, header, respBody, core.now())
		if attempt >= cfg.maxRetries || !shouldRetry(req.method, status) {
			return apiErr
		}
		wait, fromServer := core.retryDelay(header, attempt)
		if fromServer && wait > maxRetryWait {
			// The server asked for a longer pause than we are willing to
			// block for (the monthly cap's Retry-After is in days).
			return apiErr
		}
		if !core.sleep(ctx.Done(), wait) {
			return apiErr
		}
	}
}

// attempt performs one HTTP round trip and reads the whole body.
func (core *clientCore) attempt(ctx context.Context, cfg *config, req *apiRequest, u string, body []byte) (int, http.Header, []byte, error) {
	actx := ctx
	if cfg.timeout > 0 {
		var cancel context.CancelFunc
		actx, cancel = context.WithTimeout(ctx, cfg.timeout)
		defer cancel()
	}
	var rd io.Reader
	if body != nil {
		rd = bytes.NewReader(body)
	}
	hr, err := http.NewRequestWithContext(actx, req.method, u, rd)
	if err != nil {
		return 0, nil, nil, err
	}
	hr.Header.Set("User-Agent", "propraven-go/"+Version)
	hr.Header.Set("Accept", req.accept)
	if body != nil {
		hr.Header.Set("Content-Type", "application/json")
	}
	for k, vs := range cfg.headers {
		for _, v := range vs {
			hr.Header.Add(k, v)
		}
	}
	for k, vs := range req.header {
		hr.Header.Del(k)
		for _, v := range vs {
			hr.Header.Add(k, v)
		}
	}
	hr.Header.Del("Authorization")
	if cfg.apiKey != "" {
		hr.Header.Set("Authorization", "Bearer "+cfg.apiKey)
	}
	resp, err := cfg.httpClient.Do(hr)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, nil, err
	}
	return resp.StatusCode, resp.Header, b, nil
}

func shouldRetry(method string, status int) bool {
	switch status {
	case http.StatusTooManyRequests, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return true
	}
	if status >= 500 {
		switch method {
		case http.MethodGet, http.MethodHead, http.MethodDelete, http.MethodOptions:
			return true
		}
	}
	return false
}

// retryDelay picks the wait before the next attempt. fromServer reports
// whether the server dictated it (Retry-After or the rate-limit reset).
func (core *clientCore) retryDelay(h http.Header, attempt int) (time.Duration, bool) {
	if secs, ok := parseRetryAfter(h.Get("Retry-After"), core.now()); ok {
		return durationSeconds(secs), true
	}
	if strings.TrimSpace(h.Get("X-RateLimit-Remaining")) == "0" {
		if reset, err := strconv.ParseInt(strings.TrimSpace(h.Get("X-RateLimit-Reset")), 10, 64); err == nil {
			secs := float64(reset) - float64(core.now().UnixMilli())/1000
			if secs < 0 {
				secs = 0
			}
			return durationSeconds(secs), true
		}
	}
	return core.backoff(attempt), false
}

// backoff is 0.5 * 2^attempt seconds with +/-25% jitter, capped at 60s.
func (core *clientCore) backoff(attempt int) time.Duration {
	secs := 0.5 * math.Pow(2, float64(attempt))
	secs *= 1 + 0.25*core.jitter()
	d := durationSeconds(secs)
	if d > maxRetryWait {
		d = maxRetryWait
	}
	return d
}

func durationSeconds(s float64) time.Duration {
	if s <= 0 {
		return 0
	}
	if s > 1e6 {
		s = 1e6
	}
	return time.Duration(s * float64(time.Second))
}

// parseRetryAfter reads a Retry-After value in seconds or as an HTTP date.
func parseRetryAfter(v string, now time.Time) (float64, bool) {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0, false
	}
	if f, err := strconv.ParseFloat(v, 64); err == nil {
		if f < 0 || math.IsNaN(f) || math.IsInf(f, 0) {
			return 0, false
		}
		return f, true
	}
	if t, err := http.ParseTime(v); err == nil {
		s := t.Sub(now).Seconds()
		if s < 0 {
			s = 0
		}
		return s, true
	}
	return 0, false
}

func (core *clientCore) recordRateLimit(h http.Header) {
	parse := func(name string) (int64, bool) {
		v := strings.TrimSpace(h.Get(name))
		if v == "" {
			return -1, false
		}
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			if f, ferr := strconv.ParseFloat(v, 64); ferr == nil {
				return int64(f), true
			}
			return -1, false
		}
		return n, true
	}
	l, okL := parse("X-RateLimit-Limit")
	r, okR := parse("X-RateLimit-Remaining")
	s, okS := parse("X-RateLimit-Reset")
	if !okL && !okR && !okS {
		return
	}
	core.mu.Lock()
	core.lastRL = &RateLimit{Limit: l, Remaining: r, Reset: s}
	core.mu.Unlock()
}

func isTimeout(err error) bool {
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	var te interface{ Timeout() bool }
	return errors.As(err, &te) && te.Timeout()
}

func sleepCtx(done <-chan struct{}, d time.Duration) bool {
	if d <= 0 {
		select {
		case <-done:
			return false
		default:
			return true
		}
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-done:
		return false
	case <-t.C:
		return true
	}
}

func defaultJitter() float64 { return rand.Float64()*2 - 1 }

// ---- helpers used by generated code ----

type scalar interface {
	string | int64 | float64 | bool
}

func formatScalar[T scalar](v T) string {
	switch x := any(v).(type) {
	case string:
		return x
	case int64:
		return strconv.FormatInt(x, 10)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(x)
	}
	return fmt.Sprint(v)
}

// addQuery sets key when v is non-nil. Booleans encode as true/false.
func addQuery[T scalar](q url.Values, key string, v *T) {
	if v != nil {
		q.Set(key, formatScalar(*v))
	}
}

// addQueryValue always sets key (used for required, non-pointer fields).
func addQueryValue[T scalar](q url.Values, key string, v T) {
	q.Set(key, formatScalar(v))
}

// addQueryList comma-joins v, or repeats key per item when explode is true.
func addQueryList[T scalar](q url.Values, key string, v []T, explode bool) {
	if len(v) == 0 {
		return
	}
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = formatScalar(x)
	}
	if explode {
		q[key] = parts
		return
	}
	q.Set(key, strings.Join(parts, ","))
}

func setHeader[T scalar](h http.Header, key string, v *T) {
	if v != nil {
		h.Set(key, formatScalar(*v))
	}
}

func setHeaderValue[T scalar](h http.Header, key string, v T) {
	h.Set(key, formatScalar(v))
}

// pathParam escapes one path segment. "37:119:1" stays readable; "/" and
// spaces are escaped.
func pathParam[T scalar](v T) string {
	return url.PathEscape(formatScalar(v))
}
