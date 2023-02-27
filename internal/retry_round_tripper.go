package internal

import (
	"fmt"
	"net/http"
	"time"
)

// RetryRoundTripper is a struct who implements Rounder interface
//
// Its handle retry failed requests. It's useful to retry the requests
// who can fail, but can be recovered in a following call
type RetryRoundTripper struct {
	// original is the original roundtripper to call
	original http.RoundTripper
	// attemps to retry.
	//
	// by default its 3
	attempts int
	// gteCodeError is greather than or equal defined status code
	// which be consider an error
	//
	// by default its 500.
	gteCodeError int
	// retry delay
	//
	// by default its 200 milliseconds
	delay time.Duration
}

func (r *RetryRoundTripper) Wrap(h http.RoundTripper) { r.original = h }

func (r RetryRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	attempts := r.attempts
	maxloops := r.attempts + 3

	original := r.original
	if original == nil {
		original = http.DefaultTransport
	}

	for {
		resp, err := original.RoundTrip(req)

		maxloops--
		attempts--

		// max attempts reached
		if attempts == 0 {
			return resp, err
		}

		// good outcome from request
		if err == nil && resp.StatusCode < r.gteCodeError {
			return resp, err
		}

		// maximum loop reached - like a deadlock
		if maxloops <= 0 {
			return nil, fmt.Errorf("request roundtrip deadlock")
		}

		select {
		// context done
		case <-req.Context().Done():
			return resp, req.Context().Err()

		// sleep between retry
		case <-time.After(r.delay):
		}
	}
}

// RetryRoundTripper Factory

func NewRetryRoundTripper(opts ...RetryRoundTripperOptions) *RetryRoundTripper {
	rr := &RetryRoundTripper{}
	retryRoundTripperDefaults()(rr)
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(rr)
		}
	}

	return rr
}
