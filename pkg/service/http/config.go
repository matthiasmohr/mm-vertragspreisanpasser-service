package http

import (
	"time"
)

// RetryerConfig defines the set of configuration options for retrying requests.
type RetryerConfig struct {
	maxRetries  uint
	waitTime    time.Duration
	maxWaitTime time.Duration
	timeout     time.Duration
}
