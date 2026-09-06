package logging

import (
	"context"

	"github.com/redis/go-redis/v9/internal"
)

type LogLevelT = internal.LogLevelT

const (
	LogLevelError = internal.LogLevelError
	LogLevelWarn  = internal.LogLevelWarn
	LogLevelInfo  = internal.LogLevelInfo
	LogLevelDebug = internal.LogLevelDebug
)

type VoidLogger struct{}

func (v *VoidLogger) Printf(_ context.Context, _ string, _ ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Disable() { _ = "STUB: not implemented"; return }

func Enable() { _ = "STUB: not implemented"; return }

func SetLogLevel(logLevel LogLevelT) { _ = "STUB: not implemented"; return }

func NewBlacklistLogger(substr []string) internal.Logging {
	_ = "STUB: not implemented"
	return *new(internal.Logging)
}

func NewWhitelistLogger(substr []string) internal.Logging {
	_ = "STUB: not implemented"
	return *new(internal.Logging)
}

type filterLogger struct {
	logger    internal.Logging
	blacklist bool
	substr    []string
}

func (l *filterLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}
