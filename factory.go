package gohttp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

// Response Factory

func newResponse() *response {
	return &response{accept: "*/*", buf: new(bytes.Buffer)}
}

// Request Factory

func newRequest(ctx context.Context, method, baseURL, path string, body io.ReadCloser) (*request, error) {
	url := fmt.Sprintf("%s/%s", withoutTrailslash(baseURL), withoutSlash(path))
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	return (&request{}).Wrap(req), nil
}

// Client Factory

// New will create a new http client
func New(opts ...ClientOptions) *client {
	ctx := context.Background()
	return NewWithContext(ctx, opts...)
}

// NewWithContext will create a new http client using a context
func NewWithContext(ctx context.Context, opts ...ClientOptions) *client {
	c := &client{ctx: ctx}
	clientDefaults()(c)
	c.apply(opts...)
	return c
}
