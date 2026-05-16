package propraven

import (
	"net/http"
	"time"
)

// Option configures a [Client]. Apply via [NewClient].
type Option func(*Client)

// WithAPIKey sets the bearer token. Required.
func WithAPIKey(key string) Option {
	return func(c *Client) { c.apiKey = key }
}

// WithBaseURL overrides the API endpoint. Defaults to https://api.propraven.com.
// Use for self-hosted deployments or for hitting a staging environment.
func WithBaseURL(u string) Option {
	return func(c *Client) { c.baseURL = u }
}

// WithHTTPClient injects a custom http.Client — typically to attach OpenTelemetry
// transport, custom transport-level retry, or a proxy. The client's Timeout is
// honored verbatim (the SDK does not override it).
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.httpClient = h }
}

// WithTimeout sets the request timeout when [WithHTTPClient] is not used.
// Defaults to 30s. Ignored if WithHTTPClient is passed.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		if c.httpClient == nil {
			c.httpClient = &http.Client{}
		}
		c.httpClient.Timeout = d
	}
}

// WithUserAgent appends to the default User-Agent. The PropRaven server uses
// User-Agent for usage analytics; honest identification helps support.
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = defaultUA + " " + ua
	}
}
