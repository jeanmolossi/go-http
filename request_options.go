package gohttp

import "net/url"

// RequestOption is a func to define options to request
type RequestOption func(r *request)

// Headers is a object of headers who can be added
type Headers map[string]string

func requestDefaults() RequestOption {
	return func(r *request) {
		r.request.Header.Set("content-type", "application/json")
		r.request.Header.Set("accept", "application/json")
	}
}

// WithQuery will append a new query param to existing query params
func WithQuery(k, v string) RequestOption {
	return func(r *request) {
		if r.request == nil {
			return
		}

		q := r.request.URL.Query()
		q.Add(k, v)

		r.request.URL.RawQuery = q.Encode()
	}
}

// WithUrlValues will override the query params
func WithUrlValues(uv url.Values) RequestOption {
	return func(r *request) {
		r.request.URL.RawQuery = uv.Encode()
	}
}

// WithHeader will set received key and value to headers
// of the request
func WithHeader(k, v string) RequestOption {
	return func(r *request) {
		r.request.Header.Set(k, v)
	}
}

// WithHeaders will set the Headers map received to
// the headers request
//
// If a header is set, and incomes from the Headers map,
// it will be replaced
func WithHeaders(headers Headers) RequestOption {
	return func(r *request) {
		if len(headers) > 0 {
			for k, v := range headers {
				r.request.Header.Set(k, v)
			}
		}
	}
}
