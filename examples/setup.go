package examples

import (
	"fmt"
	"time"
)

var (
	BaseURL    = "https://httpbin.org"
	DeletePath = "/delete"
	GetPath    = "/get"
	PatchPath  = "/patch"
	PostPath   = "/post"
	PutPath    = "/put"

	StatusPath = "/status"
	DelayPath  = "/delay"
	StreamPath = "/stream"
)

type (
	TraceHeaders struct {
		AmznTraceID string `json:"X-Amzn-Trace-Id"`
	}

	Tracer struct {
		Headers TraceHeaders
	}
)

func intervaler(identifier string) func() {
	start := time.Now()

	fmt.Printf("[INFO] %s starts...\n", identifier)

	return func() {
		fmt.Printf("[INFO] %s taken %s\n", identifier, time.Since(start))
	}
}

func status(sc int) string {
	return fmt.Sprintf("%s/%d", StatusPath, sc)
}

func delay(sec int) string {
	if sec > 10 {
		sec = 10
	}

	if sec <= 0 {
		sec = 1
	}

	return fmt.Sprintf("%s/%d", DelayPath, sec)
}

// func stream(n int) string {
// 	return fmt.Sprintf("%s/%d", StreamPath, n)
// }
