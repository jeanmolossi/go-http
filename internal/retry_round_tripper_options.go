package internal

import (
	"net/http"
	"time"
)

type RetryRoundTripperOptions func(*RetryRoundTripper)

func retryRoundTripperDefaults() RetryRoundTripperOptions {
	return func(rrt *RetryRoundTripper) {
		rrt.attempts = 3
		rrt.gteCodeError = http.StatusInternalServerError
		rrt.delay = time.Millisecond * 200
	}
}

func RetryWithAttempts(attempts int) RetryRoundTripperOptions {
	return func(rrt *RetryRoundTripper) {
		rrt.attempts = attempts
	}
}

func MinStatusCodeToRetry(statusCode int) RetryRoundTripperOptions {
	return func(rrt *RetryRoundTripper) {
		rrt.gteCodeError = statusCode
	}
}

func IntervalBetweenRetry(i time.Duration) RetryRoundTripperOptions {
	return func(rrt *RetryRoundTripper) {
		rrt.delay = i
	}
}
