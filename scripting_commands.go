package redis

import "context"

type ScriptingFunctionsCmdable interface {
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd
	EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
	EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd
	ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd
	ScriptFlush(ctx context.Context) *StatusCmd
	ScriptKill(ctx context.Context) *StatusCmd
	ScriptLoad(ctx context.Context, script string) *StringCmd

	FunctionLoad(ctx context.Context, code string) *StringCmd
	FunctionLoadReplace(ctx context.Context, code string) *StringCmd
	FunctionDelete(ctx context.Context, libName string) *StringCmd
	FunctionFlush(ctx context.Context) *StringCmd
	FunctionKill(ctx context.Context) *StringCmd
	FunctionFlushAsync(ctx context.Context) *StringCmd
	FunctionList(ctx context.Context, q FunctionListQuery) *FunctionListCmd
	FunctionDump(ctx context.Context) *StringCmd
	FunctionRestore(ctx context.Context, libDump string) *StringCmd
	FunctionStats(ctx context.Context) *FunctionStatsCmd
	FCall(ctx context.Context, function string, keys []string, args ...interface{}) *Cmd
	FCallRo(ctx context.Context, function string, keys []string, args ...interface{}) *Cmd
	FCallRO(ctx context.Context, function string, keys []string, args ...interface{}) *Cmd
}

func (c cmdable) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) eval(ctx context.Context, name, payload string, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ScriptFlush(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ScriptKill(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ScriptLoad(ctx context.Context, script string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

type FunctionListQuery struct {
	LibraryNamePattern string
	WithCode           bool
}

func (c cmdable) FunctionLoad(ctx context.Context, code string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FunctionLoadReplace(ctx context.Context, code string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FunctionDelete(ctx context.Context, libName string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FunctionFlush(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FunctionKill(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FunctionFlushAsync(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FunctionList(ctx context.Context, q FunctionListQuery) *FunctionListCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FunctionDump(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FunctionRestore(ctx context.Context, libDump string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FunctionStats(ctx context.Context) *FunctionStatsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FCall(ctx context.Context, function string, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FCallRo(ctx context.Context, function string, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FCallRO(ctx context.Context, function string, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func fcallArgs(command string, function string, keys []string, args ...interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}
