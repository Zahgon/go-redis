//go:build !appengine
// +build !appengine

package rediscmd

func String(b []byte) string { _ = "STUB: not implemented"; return "" }

func Bytes(s string) []byte { _ = "STUB: not implemented"; return nil }
