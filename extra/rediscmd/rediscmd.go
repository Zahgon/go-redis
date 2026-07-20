package rediscmd

import (
	"github.com/redis/go-redis/v9"
)

func CmdString(cmd redis.Cmder) string { _ = "STUB: not implemented"; return "" }

func CmdsString(cmds []redis.Cmder) (string, string) { _ = "STUB: not implemented"; return "", "" }

func AppendCmd(b []byte, cmd redis.Cmder) []byte { _ = "STUB: not implemented"; return nil }

func appendArg(b []byte, v interface{}) []byte { _ = "STUB: not implemented"; return nil }

func appendUTF8String(dst []byte, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func isSimple(b []byte) bool { _ = "STUB: not implemented"; return false }

func isSimpleByte(c byte) bool { _ = "STUB: not implemented"; return false }
