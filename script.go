package redis

import (
	"context"
	"sync"
)

type Scripter interface {
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd
	EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
	EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd
	ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd
	ScriptLoad(ctx context.Context, script string) *StringCmd
}

var (
	_ Scripter = (*Client)(nil)
	_ Scripter = (*Ring)(nil)
	_ Scripter = (*ClusterClient)(nil)
)

type Script struct {
	src       string
	mu        sync.RWMutex
	hash      string
	serverSHA bool
}

func NewScript(src string) *Script { _ = "STUB: not implemented"; return nil }

func NewScriptServerSHA(src string) *Script { _ = "STUB: not implemented"; return nil }

func (s *Script) Hash() string { _ = "STUB: not implemented"; return "" }

func (s *Script) Load(ctx context.Context, c Scripter) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) Exists(ctx context.Context, c Scripter) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) Eval(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) EvalRO(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) ensureHash(ctx context.Context, c Scripter) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) EvalSha(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) EvalShaRO(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) Run(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) RunRO(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}
