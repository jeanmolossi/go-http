package gohttp

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/jeanmolossi/go-http/internal"
)

// ClientOptions is a option definer
type ClientOptions func(c *client)

// clientDefaults apply de default configs to the client
func clientDefaults() ClientOptions {
	return func(c *client) {
		if c.baseURL == "" {
			c.baseURL = "http://localhost"
		}

		if c.ctx == nil {
			c.ctx = context.Background()
		}

		c.client = http.DefaultClient
		c.client.Timeout = time.Second * 30

		c.baseRequestOpts = []RequestOption{
			WithHeaders(Headers{
				"content-type": "application/json",
				"accept":       "application/json",
			}),
		}
	}
}

// WithRequestOptions defines on the client the all
// requests baseOptions. All requests outgoing from the
// parent client will call with this options as base options
//
// Ex:
//
//	gohttp.WithRequestOptions(
//		gohttp.WithQuery("key", "value")
//	)
func WithRequestOptions(opts ...RequestOption) ClientOptions {
	return func(c *client) {
		c.baseRequestOpts = opts
	}
}

// WithBaseURL is used to define baseURL in client
func WithBaseURL(baseURL string) ClientOptions {
	return func(c *client) { c.baseURL = baseURL }
}

// WithRoundTripper appends a roundtripper to http client.
//
// *IMPORTANT*
//
// Becareful using this. Make sure you append a roundtripper
// and are calling the original roundtripper inside your custom.
//
// Look more into internal roundtrippers.
//
// *ADDITIONAL INFO*
//
// It will NOT replace the original roundtripper, it will only
// decorate the existing roundtrippers.
func WithRoundTripper(r Rounder) ClientOptions {
	return func(c *client) {
		if c.client == nil {
			c.client = http.DefaultClient
		}

		if r == nil {
			r = &internal.ClientRounder{}
		}

		r.Wrap(c.client.Transport)
		c.client.Transport = r
	}
}

// WithRetry will config a retry strategy to requests
//
//   - attempts is how many retries will occur
//   - considerErrFromStatus is the status who will be consider an error status
//     -- Default considerErrFromStatus is 500.
//   - interval is duration who requests wait between retries
func WithRetry(attempts, considerErrFromStatus int, interval time.Duration) ClientOptions {
	return func(c *client) {
		WithRoundTripper(internal.NewRetryRoundTripper(
			internal.RetryWithAttempts(attempts),
			internal.MinStatusCodeToRetry(considerErrFromStatus),
			internal.IntervalBetweenRetry(interval),
		))(c)
	}
}

// WithProxy will configure a proxy to do the requests
func WithProxy(proxy string) ClientOptions {
	url, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}

	return func(c *client) {
		internal.NewProxyRoundTripper(url)
	}
}

// WithTimeout will configure the timeout of the requests.
//
// If it is not set the Default timeout is 30 seconds
func WithTimeout(timeout time.Duration) ClientOptions {
	return func(c *client) {
		c.client.Timeout = timeout
	}
}
