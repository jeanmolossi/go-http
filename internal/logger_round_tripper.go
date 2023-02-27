package internal

import (
	"fmt"
	"log"
	"net/http"
)

type Logger interface {
	Info(format string, v ...any)
	Error(format string, v ...any)
}

type LogRounder struct {
	original http.RoundTripper
	log      Logger
}

func (l *LogRounder) Wrap(h http.RoundTripper) { l.original = h }
func (l LogRounder) RoundTrip(req *http.Request) (*http.Response, error) {
	logg := l.log
	if logg == nil {
		logg = &logger{}
	}

	logg.Info("started %s request in %v", req.Method, req.URL)

	original := l.original
	if original == nil {
		original = http.DefaultTransport
	}

	resp, err := original.RoundTrip(req)
	if err != nil {
		logg.Error("fail doing request: %v", err)
		return nil, err
	}

	logg.Info("end %s request in %v", req.Method, req.URL)
	return resp, err
}

func NewLogRounder() *LogRounder {
	r := &LogRounder{}
	r.log = &logger{}
	return r
}

type logger struct{}

func (l *logger) Info(format string, v ...any) {
	f := fmt.Sprintf("[INFO] %s", format)
	log.Printf(f, v...)
}

func (l *logger) Error(format string, v ...any) {
	f := fmt.Sprintf("[ERROR] %s", format)
	log.Printf(f, v...)
}
