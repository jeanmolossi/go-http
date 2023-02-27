package gohttp

import (
	"fmt"
	"net/http"
)

type request struct {
	request *http.Request
}

// Wrap will wrap the http.Request in client request
func (r *request) Wrap(req *http.Request) *request { r.request = req; return r }

// prepareRequest will apply RequestOptions to the request
//
// it can result an error if http.Request is not Wrap
func (r *request) prepareRequest(opts ...RequestOption) (*request, error) {
	if r.request == nil {
		return r, fmt.Errorf("wrap a request before prepare it")
	}

	r.apply(opts...)
	return r, nil
}

// apply will apply the request options to the request
func (r *request) apply(opts ...RequestOption) {
	requestDefaults()(r)

	if len(opts) > 0 {
		for _, opt := range opts {
			opt(r)
		}
	}
}

// sendRequest will execute the client.Do method who really sends
// the request
//
// it can result an error if http.Request was not Wrap
//
// When client has no internal http.Client it will call using
// http.DefaultClient
func (r *request) sendRequest(c *client) (*http.Response, error) {
	if r.request == nil {
		return nil, fmt.Errorf("wrap a request before send it")
	}

	if c.client == nil {
		return http.DefaultClient.Do(r.request)
	}

	return c.client.Do(r.request)
}
