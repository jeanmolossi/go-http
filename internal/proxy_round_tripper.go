package internal

import (
	"log"
	"net/http"
	"net/url"
)

type ProxyRoundTripper struct {
	original http.RoundTripper
	url      *url.URL
}

func (p *ProxyRoundTripper) Wrap(h http.RoundTripper) { p.original = h }

func (p ProxyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	original := p.original
	if original == nil {
		original = http.DefaultTransport
	}

	if o, ok := original.(*http.Transport); ok {
		o.Proxy = http.ProxyURL(p.url)
		log.Println("getting with proxy", p.url)
		return o.RoundTrip(req)
	}

	return original.RoundTrip(req)
}

func NewProxyRoundTripper(url *url.URL) *ProxyRoundTripper {
	return &ProxyRoundTripper{url: url}
}
