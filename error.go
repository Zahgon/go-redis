package redis

import (
	"github.com/redis/go-redis/v9/internal/pool"
	"github.com/redis/go-redis/v9/internal/proto"
)

var ErrClosed = pool.ErrClosed

var ErrPoolExhausted = pool.ErrPoolExhausted

var ErrPoolTimeout = pool.ErrPoolTimeout

var ErrCrossSlot = proto.RedisError("CROSSSLOT Keys in request don't hash to the same slot")

var ErrNoScript = proto.RedisError("NOSCRIPT No matching script. Please use EVAL.")

func HasErrorPrefix(err error, prefix string) bool { _ = "STUB: not implemented"; return false }

type Error interface {
	error

	RedisError()
}

var _ Error = proto.RedisError("")

func isContextError(err error) bool { _ = "STUB: not implemented"; return false }

func isTimeoutError(err error) (isTimeout bool, hasTimeoutFlag bool) {
	_ = "STUB: not implemented"
	return false, false
}

func shouldRetry(err error, retryTimeout bool) bool { _ = "STUB: not implemented"; return false }

func isRedisError(err error) bool { _ = "STUB: not implemented"; return false }

func isBadConn(err error, allowTimeout bool, addr string) bool {
	_ = "STUB: not implemented"
	return false
}

func isMovedError(err error) (moved bool, ask bool, addr string) {
	_ = "STUB: not implemented"
	return false, false, ""
}

func isLoadingError(err error) bool { _ = "STUB: not implemented"; return false }

func isReadOnlyError(err error) bool { _ = "STUB: not implemented"; return false }

func isMovedSameConnAddr(err error, addr string) bool { _ = "STUB: not implemented"; return false }

func IsLoadingError(err error) bool { _ = "STUB: not implemented"; return false }

func IsReadOnlyError(err error) bool { _ = "STUB: not implemented"; return false }

func IsClusterDownError(err error) bool { _ = "STUB: not implemented"; return false }

func IsTryAgainError(err error) bool { _ = "STUB: not implemented"; return false }

func IsMasterDownError(err error) bool { _ = "STUB: not implemented"; return false }

func IsMaxClientsError(err error) bool { _ = "STUB: not implemented"; return false }

func IsMovedError(err error) (addr string, ok bool) { _ = "STUB: not implemented"; return "", false }

func IsAskError(err error) (addr string, ok bool) { _ = "STUB: not implemented"; return "", false }

func IsAuthError(err error) bool { _ = "STUB: not implemented"; return false }

func IsPermissionError(err error) bool { _ = "STUB: not implemented"; return false }

func IsExecAbortError(err error) bool { _ = "STUB: not implemented"; return false }

func IsOOMError(err error) bool { _ = "STUB: not implemented"; return false }

func IsNoReplicasError(err error) bool { _ = "STUB: not implemented"; return false }

type timeoutError interface {
	Timeout() bool
}
