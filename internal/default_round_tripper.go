package internal

import "net/http"

type ClientRounder struct{}

func (c *ClientRounder) Wrap(http.RoundTripper) {}
func (c ClientRounder) RoundTrip(req *http.Request) (*http.Response, error) {
	return http.DefaultTransport.RoundTrip(req)
}
