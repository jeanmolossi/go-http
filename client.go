package gohttp

import (
	"context"
	"io"
	"net/http"
)

// client is used to build a client wrapper
//
// it contains the requests base configs, the base client
// and wraps a context too
type client struct {
	// baseURL will be defined to localhost as default
	//
	// use the option `gohttp.WithBaseURL(baseURL)` to update
	// the baseURL
	baseURL string

	// client is the original http.Client pointer
	client *http.Client

	// baseRequestOpts memoize in the client the base
	// options to all requests
	//
	// every request outgoing from this client will
	// be applyied with the baseRequestOpts
	baseRequestOpts []RequestOption

	// ctx wraps a client context. Useful to cancel all
	// outgoing requests based on that context
	ctx context.Context
}

// mountRequest will mount base request to any method.
//
// all requests will follow same format
func (c *client) mountRequest(ctx context.Context, path, method string, body io.ReadCloser, options ...RequestOption) (*response, error) {
	// starts a response with initial values
	res := newResponse()

	// starts a newRequest
	req, err := newRequest(ctx, method, c.baseURL, path, body)
	// annotate response with acceptable header
	res.accept = req.request.Header.Get("accept")
	if err != nil {
		return res, err
	}

	req, err = req.prepareRequest(append(c.baseRequestOpts, options...)...)
	if err != nil {
		return res, err
	}

	// wrap client response with a response, error received from
	// the request result
	return res.Wrap(req.sendRequest(c)).Read()
}

// apply options to client
func (c *client) apply(opts ...ClientOptions) *client {
	if len(opts) == 0 {
		return c
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// ConnectWithContext will send a Http CONNECT request
func (c *client) ConnectWithContext(ctx context.Context, path string, options ...RequestOption) (*response, error) {
	return c.mountRequest(ctx, path, http.MethodConnect, nil, options...)
}

// Connect will send a Http CONNECT request
func (c *client) Connect(path string, options ...RequestOption) (*response, error) {
	return c.ConnectWithContext(c.ctx, path, options...)
}

// GetWithContext will send a Http GET request
func (c *client) GetWithContext(ctx context.Context, path string, options ...RequestOption) (*response, error) {
	return c.mountRequest(ctx, path, http.MethodGet, nil, options...)
}

// Get will send a Http GET request
func (c *client) Get(path string, options ...RequestOption) (*response, error) {
	return c.GetWithContext(c.ctx, path, options...)
}

// HeadWithContext will send a Http HEAD request
func (c *client) HeadWithContext(ctx context.Context, path string, options ...RequestOption) (*response, error) {
	return c.mountRequest(ctx, path, http.MethodHead, nil, options...)
}

// Head will send a Http HEAD request
func (c *client) Head(path string, options ...RequestOption) (*response, error) {
	return c.HeadWithContext(c.ctx, path, options...)
}

// OptionsWithContext will send a Http OPTIONS request
func (c *client) OptionsWithContext(ctx context.Context, path string, options ...RequestOption) (*response, error) {
	return c.mountRequest(ctx, path, http.MethodOptions, nil, options...)
}

// Options will send a Http OPTIONS request
func (c *client) Options(path string, options ...RequestOption) (*response, error) {
	return c.OptionsWithContext(c.ctx, path, options...)
}

// PostWithContext will send a Http POST request
func (c *client) PostWithContext(ctx context.Context, path string, body io.ReadCloser, options ...RequestOption) (*response, error) {
	return c.mountRequest(ctx, path, http.MethodPost, body, options...)
}

// Post will send a Http POST request
func (c *client) Post(path string, body io.ReadCloser, options ...RequestOption) (*response, error) {
	return c.PostWithContext(c.ctx, path, body, options...)
}

// PatchWithContext will send a Http PATCH request
func (c *client) PatchWithContext(ctx context.Context, path string, body io.ReadCloser, options ...RequestOption) (*response, error) {
	return c.mountRequest(ctx, path, http.MethodPatch, body, options...)
}

// Patch will send a Http PATCH request
func (c *client) Patch(path string, body io.ReadCloser, options ...RequestOption) (*response, error) {
	return c.PatchWithContext(c.ctx, path, body, options...)
}

// PutWithContext will send a Http PUT request
func (c *client) PutWithContext(ctx context.Context, path string, body io.ReadCloser, options ...RequestOption) (*response, error) {
	return c.mountRequest(ctx, path, http.MethodPut, body, options...)
}

// Put will send a Http PUT request
func (c *client) Put(path string, body io.ReadCloser, options ...RequestOption) (*response, error) {
	return c.PutWithContext(c.ctx, path, body, options...)
}

// DeleteWithContext will send a Http DELETE request
func (c *client) DeleteWithContext(ctx context.Context, path string, options ...RequestOption) (*response, error) {
	return c.mountRequest(ctx, path, http.MethodDelete, nil, options...)
}

// Delete will send a Http DELETE request
func (c *client) Delete(path string, options ...RequestOption) (*response, error) {
	return c.DeleteWithContext(c.ctx, path, options...)
}
