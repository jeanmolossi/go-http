package gohttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	acceptAny       = `\*\/\*`
	applicationJSON = `application\/(hal[+\-])?json`
	textHTML        = `text\/html`
)

type response struct {
	// accept is the request accept header
	// the default is: "*/*"
	accept string
	// if sendRequest result with an error it will
	// be defined
	err error
	// response is the original http.Response
	response *http.Response
	// buf is used to read the request result body
	buf *bytes.Buffer
}

// Wrap will wrap original http.Response, err to the client response
func (r *response) Wrap(res *http.Response, err error) *response {
	if r == nil {
		return &response{response: res, err: err}
	}

	r.response = res
	r.err = err
	if err != nil {
		return r
	}

	_, err = r.buf.ReadFrom(res.Body)
	if err != nil {
		r.err = err
	}
	defer res.Body.Close()

	return r
}

// Read will get this response pointer and err if is set
func (r *response) Read() (*response, error) {
	if r.err != nil {
		return r, r.err
	}

	return r, nil
}

// Success returns true to a response status code lower than 300
func (r *response) Success() bool {
	return r.response.StatusCode < 300
}

// Status will returns the statusCode int from the result
func (r *response) Status() int {
	return r.response.StatusCode
}

// Bytes get the response body in bytes format
func (r *response) Bytes() []byte {
	return r.buf.Bytes()
}

// Interface will unmarshal the body to a interface.
// Useful if you want to unmarshal response who is not
// a defined struct
func (r *response) Interface(target any) error {
	return json.Unmarshal(r.buf.Bytes(), target)
}

// JSON is a method used to Unmarshal the result body
// into target received.
//
// It will result an error if acceptable header does
// not match with "application/json" or "*/*"
func (r *response) JSON(target any) error {
	cc := r.response.Header.Get("content-type")

	if !isAcceptable(applicationJSON, cc) {
		return fmt.Errorf("can not parse target to a non json response")
	}

	if !isAcceptable(applicationJSON, r.accept) {
		return fmt.Errorf("can not parse JSON to request who accept %v", r.accept)
	}

	return json.Unmarshal(r.Bytes(), target)
}
