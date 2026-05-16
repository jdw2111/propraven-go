// Package propraven is the Go SDK for the PropRaven API
// (https://api.propraven.com).
//
// Quickstart:
//
//	client, err := propraven.NewClient(
//	    propraven.WithAPIKey(os.Getenv("PROPRAVEN_API_KEY")),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	parcel, err := client.Parcels.Get(ctx, "06037:1234-567-890")
//
// For webhook payload verification see [VerifyWebhook].
package propraven

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://api.propraven.com"
	defaultUA      = "propraven-go/" + Version
	defaultTimeout = 30 * time.Second
)

// Client is the entry point for the PropRaven SDK. Construct one with [NewClient]
// and reuse across goroutines — the underlying http.Client is concurrency-safe.
type Client struct {
	baseURL    string
	apiKey     string
	userAgent  string
	httpClient *http.Client

	// Resource handles. Each is bound to this client; do not construct directly.
	Parcels  *ParcelsService
	Owners   *OwnersService
	Deeds    *DeedsService
	Permits  *PermitsService
	Webhooks *WebhooksService
	Health   *HealthService
}

// NewClient constructs a Client. Pass [WithAPIKey] at minimum; [WithBaseURL],
// [WithHTTPClient], and [WithUserAgent] override the defaults. Returns an
// error only on contradictory options (e.g. missing API key in production).
func NewClient(opts ...Option) (*Client, error) {
	c := &Client{
		baseURL:    defaultBaseURL,
		userAgent:  defaultUA,
		httpClient: &http.Client{Timeout: defaultTimeout},
	}
	for _, opt := range opts {
		opt(c)
	}
	if strings.TrimSpace(c.apiKey) == "" {
		return nil, errors.New("propraven: API key required — pass WithAPIKey(...) or set PROPRAVEN_API_KEY")
	}

	c.Parcels = &ParcelsService{client: c}
	c.Owners = &OwnersService{client: c}
	c.Deeds = &DeedsService{client: c}
	c.Permits = &PermitsService{client: c}
	c.Webhooks = &WebhooksService{client: c}
	c.Health = &HealthService{client: c}
	return c, nil
}

// do issues an authenticated HTTP request, decodes a JSON body on success, and
// returns a typed [Error] on non-2xx responses. body may be nil for GET/DELETE.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any, out any) error {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return fmt.Errorf("propraven: bad URL %q: %w", path, err)
	}
	if query != nil {
		u.RawQuery = query.Encode()
	}

	var reqBody io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("propraven: marshal body: %w", err)
		}
		reqBody = bytes.NewReader(buf)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reqBody)
	if err != nil {
		return fmt.Errorf("propraven: build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("propraven: http error: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(io.LimitReader(resp.Body, 16<<20))
	if err != nil {
		return fmt.Errorf("propraven: read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return decodeError(resp, respBody)
	}

	if out != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, out); err != nil {
			return fmt.Errorf("propraven: decode response: %w (body=%s)", err, truncate(respBody, 256))
		}
	}
	return nil
}

func truncate(b []byte, n int) string {
	if len(b) <= n {
		return string(b)
	}
	return string(b[:n]) + "…"
}
