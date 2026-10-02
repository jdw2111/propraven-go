package propraven

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jdw2111/propraven-go/internal"
)

// Version is the SDK version, sent in the User-Agent header.
const Version = internal.PackageVersion

// DefaultBaseURL is the production API origin. Every operation path already
// starts with /api/v1.
const DefaultBaseURL = "https://propraven.com"

const (
	defaultTimeout    = 60 * time.Second
	defaultMaxRetries = 2
	maxRetryWait      = 60 * time.Second
)

// Option configures a [Client] (passed to [NewClient]) or a single request
// (passed as a trailing argument to any operation method). Options given to a
// method apply to that call only and override the client's settings.
type Option func(*config)

// RequestOption is an [Option] passed to a single operation call. It is the
// same type as Option; the separate name documents intent in signatures.
type RequestOption = Option

type config struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	maxRetries int
	timeout    time.Duration
	headers    http.Header
	logger     *log.Logger
	raw        *RawResponse
}

func (c config) clone() config {
	c.headers = c.headers.Clone()
	return c
}

// WithAPIKey sets the API key sent as "Authorization: Bearer <key>". It
// overrides the PROPRAVEN_API_KEY environment variable. Keys start with
// "pz_"; NewClient logs a warning (and still proceeds) when they do not.
func WithAPIKey(key string) Option {
	return func(c *config) { c.apiKey = strings.TrimSpace(key) }
}

// WithBaseURL overrides the API origin (default https://propraven.com, or
// the PROPRAVEN_BASE_URL environment variable). Do not include /api/v1.
func WithBaseURL(baseURL string) Option {
	return func(c *config) { c.baseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/") }
}

// WithHTTPClient sets the *http.Client used for requests. Its Timeout, if
// any, applies in addition to [WithTimeout].
func WithHTTPClient(hc *http.Client) Option {
	return func(c *config) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}

// WithMaxRetries sets how many times a failed request is retried (default 2).
// Zero disables retries.
func WithMaxRetries(n int) Option {
	return func(c *config) {
		if n < 0 {
			n = 0
		}
		c.maxRetries = n
	}
}

// WithTimeout sets the timeout for each HTTP attempt (default 60s). Zero
// disables the SDK timeout; the context deadline still applies.
func WithTimeout(d time.Duration) Option {
	return func(c *config) {
		if d < 0 {
			d = 0
		}
		c.timeout = d
	}
}

// WithHeader adds a header to every request (or to one request when passed
// to a method). It cannot override Authorization; use [WithAPIKey].
func WithHeader(key, value string) Option {
	return func(c *config) {
		if c.headers == nil {
			c.headers = http.Header{}
		}
		c.headers.Set(key, value)
	}
}

// WithLogger sets where the SDK writes warnings (for example an API key that
// does not start with "pz_"). The default writes to stderr; nil silences it.
func WithLogger(l *log.Logger) Option {
	return func(c *config) {
		if l == nil {
			l = log.New(io.Discard, "", 0)
		}
		c.logger = l
	}
}

// RawResponse receives the final HTTP response of a call when passed through
// [WithRawResponse]. It is filled for both successful and failed calls.
type RawResponse struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}

// WithRawResponse captures the raw status, headers and body of one call into
// dst. Useful for fields the typed response does not model yet.
func WithRawResponse(dst *RawResponse) RequestOption {
	return func(c *config) { c.raw = dst }
}

// RateLimit is the rate-limit state reported by the most recent response.
// A field is -1 when its header was absent.
type RateLimit struct {
	// Limit is X-RateLimit-Limit: requests allowed in the current window.
	Limit int64
	// Remaining is X-RateLimit-Remaining: requests left in the window.
	Remaining int64
	// Reset is X-RateLimit-Reset: when the window resets, Unix seconds.
	Reset int64
}

// ResetTime returns Reset as a time.Time (the zero Time when unknown).
func (r RateLimit) ResetTime() time.Time {
	if r.Reset < 0 {
		return time.Time{}
	}
	return time.Unix(r.Reset, 0)
}

// clientCore is the hand-written transport state behind a generated Client.
type clientCore struct {
	cfg config

	mu     sync.Mutex
	lastRL *RateLimit

	// Test hooks.
	sleep  func(ctxDone <-chan struct{}, d time.Duration) bool
	now    func() time.Time
	jitter func() float64 // returns a value in [-1, 1)
}

// NewClient builds a client. With no options it reads PROPRAVEN_API_KEY and
// PROPRAVEN_BASE_URL from the environment. A missing API key is allowed:
// several endpoints are key-optional and the server answers 401 where a key
// is required.
//
// The PropRaven REST API is meant to be called from servers: API keys are
// secrets and the API sends no CORS headers.
func NewClient(opts ...Option) *Client {
	cfg := config{
		apiKey:     strings.TrimSpace(os.Getenv("PROPRAVEN_API_KEY")),
		baseURL:    DefaultBaseURL,
		httpClient: &http.Client{},
		maxRetries: defaultMaxRetries,
		timeout:    defaultTimeout,
		headers:    http.Header{},
		logger:     log.New(os.Stderr, "propraven: ", log.LstdFlags),
	}
	if env := strings.TrimSpace(os.Getenv("PROPRAVEN_BASE_URL")); env != "" {
		cfg.baseURL = strings.TrimRight(env, "/")
	}
	for _, o := range opts {
		if o != nil {
			o(&cfg)
		}
	}
	cfg.raw = nil
	if cfg.apiKey != "" && !strings.HasPrefix(cfg.apiKey, "pz_") {
		cfg.logger.Printf("warning: API key does not start with \"pz_\"; PropRaven keys normally do (continuing anyway)")
	}
	c := &Client{core: &clientCore{
		cfg:    cfg,
		sleep:  sleepCtx,
		now:    time.Now,
		jitter: defaultJitter,
	}}
	c.initServices()
	return c
}

// LastRateLimit returns the rate-limit headers of the most recent response
// that carried any, or nil if none has yet.
func (c *Client) LastRateLimit() *RateLimit {
	c.core.mu.Lock()
	defer c.core.mu.Unlock()
	if c.core.lastRL == nil {
		return nil
	}
	rl := *c.core.lastRL
	return &rl
}

// BaseURL returns the API origin the client sends requests to.
func (c *Client) BaseURL() string { return c.core.cfg.baseURL }
