package gohttp

import (
	"net/http"
)

// Rounder is the interface used to decorate the original
// Client RoundTripper
type Rounder interface {
	// http.RoundTripper composes the roundtripper interface
	http.RoundTripper
	// Wrap is required method to decorate the original roundTripper.
	//
	// The default implementation of this method should be like:
	//
	// Example:
	//
	//	// Wrap will set original prop roundtripper on the *RounderImpl struct
	//	func (r *RounderImpl) Wrap(h http.RoundTripper) { r.original = h }
	Wrap(rt http.RoundTripper)
}
