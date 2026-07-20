package redis

import (
	"time"
)

func DialRetryBackoffConstant(d time.Duration) func(attempt int) time.Duration {
	_ = "STUB: not implemented"
	return nil
}

func DialRetryBackoffExponential(minBackoff, maxBackoff time.Duration) func(attempt int) time.Duration {
	_ = "STUB: not implemented"
	return nil
}
