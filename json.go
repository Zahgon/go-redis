package redis

import (
	"context"

	"github.com/redis/go-redis/v9/internal/proto"
)

type JSONCmdable interface {
	JSONArrAppend(ctx context.Context, key, path string, values ...interface{}) *IntSliceCmd
	JSONArrIndex(ctx context.Context, key, path string, value ...interface{}) *IntSliceCmd
	JSONArrIndexWithArgs(ctx context.Context, key, path string, options *JSONArrIndexArgs, value ...interface{}) *IntSliceCmd
	JSONArrInsert(ctx context.Context, key, path string, index int64, values ...interface{}) *IntSliceCmd
	JSONArrLen(ctx context.Context, key, path string) *IntSliceCmd
	JSONArrPop(ctx context.Context, key, path string, index int) *StringSliceCmd
	JSONArrTrim(ctx context.Context, key, path string) *IntSliceCmd
	JSONArrTrimWithArgs(ctx context.Context, key, path string, options *JSONArrTrimArgs) *IntSliceCmd
	JSONClear(ctx context.Context, key, path string) *IntCmd
	JSONDebugMemory(ctx context.Context, key, path string) *IntCmd
	JSONDel(ctx context.Context, key, path string) *IntCmd
	JSONForget(ctx context.Context, key, path string) *IntCmd
	JSONGet(ctx context.Context, key string, paths ...string) *JSONCmd
	JSONGetWithArgs(ctx context.Context, key string, options *JSONGetArgs, paths ...string) *JSONCmd
	JSONMerge(ctx context.Context, key, path string, value string) *StatusCmd
	JSONMSetArgs(ctx context.Context, docs []JSONSetArgs) *StatusCmd
	JSONMSet(ctx context.Context, params ...interface{}) *StatusCmd
	JSONMGet(ctx context.Context, path string, keys ...string) *JSONSliceCmd
	JSONNumIncrBy(ctx context.Context, key, path string, value float64) *JSONCmd
	JSONObjKeys(ctx context.Context, key, path string) *SliceCmd
	JSONObjLen(ctx context.Context, key, path string) *IntPointerSliceCmd
	JSONSet(ctx context.Context, key, path string, value interface{}) *StatusCmd
	JSONSetMode(ctx context.Context, key, path string, value interface{}, mode string) *StatusCmd
	JSONSetWithArgs(ctx context.Context, key, path string, value interface{}, options *JSONSetArgsOptions) *StatusCmd
	JSONStrAppend(ctx context.Context, key, path, value string) *IntPointerSliceCmd
	JSONStrLen(ctx context.Context, key, path string) *IntPointerSliceCmd
	JSONToggle(ctx context.Context, key, path string) *IntPointerSliceCmd
	JSONType(ctx context.Context, key, path string) *JSONSliceCmd
}

type JSONSetArgs struct {
	Key   string
	Path  string
	Value interface{}
}

type JSONArrIndexArgs struct {
	Start int
	Stop  *int
}

type JSONArrTrimArgs struct {
	Start int
	Stop  *int
}

type FPHAType string

const (
	FPHATypeBF16 FPHAType = "BF16"
	FPHATypeFP16 FPHAType = "FP16"
	FPHATypeFP32 FPHAType = "FP32"
	FPHATypeFP64 FPHAType = "FP64"
)

type JSONSetArgsOptions struct {
	Mode string
	FPHA FPHAType
}

type JSONCmd struct {
	baseCmd
	val      string
	expanded interface{}
}

var _ Cmder = (*JSONCmd)(nil)

func newJSONCmd(ctx context.Context, args ...interface{}) *JSONCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *JSONCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *JSONCmd) SetVal(val string) { _ = "STUB: not implemented"; return }

func (cmd *JSONCmd) Val() string { _ = "STUB: not implemented"; return "" }

func (cmd *JSONCmd) Result() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cmd *JSONCmd) Expanded() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *JSONCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *JSONCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type JSONSliceCmd struct {
	baseCmd
	val []interface{}
}

func NewJSONSliceCmd(ctx context.Context, args ...interface{}) *JSONSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *JSONSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *JSONSliceCmd) SetVal(val []interface{}) { _ = "STUB: not implemented"; return }

func (cmd *JSONSliceCmd) Val() []interface{} { _ = "STUB: not implemented"; return nil }

func (cmd *JSONSliceCmd) Result() ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *JSONSliceCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *JSONSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type IntPointerSliceCmd struct {
	baseCmd
	val []*int64
}

func NewIntPointerSliceCmd(ctx context.Context, args ...interface{}) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *IntPointerSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *IntPointerSliceCmd) SetVal(val []*int64) { _ = "STUB: not implemented"; return }

func (cmd *IntPointerSliceCmd) Val() []*int64 { _ = "STUB: not implemented"; return nil }

func (cmd *IntPointerSliceCmd) Result() ([]*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *IntPointerSliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *IntPointerSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) JSONArrAppend(ctx context.Context, key, path string, values ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONArrIndex(ctx context.Context, key, path string, value ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONArrIndexWithArgs(ctx context.Context, key, path string, options *JSONArrIndexArgs, value ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONArrInsert(ctx context.Context, key, path string, index int64, values ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONArrLen(ctx context.Context, key, path string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONArrPop(ctx context.Context, key, path string, index int) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONArrTrim(ctx context.Context, key, path string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONArrTrimWithArgs(ctx context.Context, key, path string, options *JSONArrTrimArgs) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONClear(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONDebugMemory(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONDel(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONForget(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONGet(ctx context.Context, key string, paths ...string) *JSONCmd {
	_ = "STUB: not implemented"
	return nil
}

type JSONGetArgs struct {
	Indent  string
	Newline string
	Space   string
}

func (c cmdable) JSONGetWithArgs(ctx context.Context, key string, options *JSONGetArgs, paths ...string) *JSONCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONMerge(ctx context.Context, key, path string, value string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONMGet(ctx context.Context, path string, keys ...string) *JSONSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONMSetArgs(ctx context.Context, docs []JSONSetArgs) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONMSet(ctx context.Context, params ...interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONNumIncrBy(ctx context.Context, key, path string, value float64) *JSONCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONObjKeys(ctx context.Context, key, path string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONObjLen(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONSet(ctx context.Context, key, path string, value interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONSetMode(ctx context.Context, key, path string, value interface{}, mode string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONSetWithArgs(ctx context.Context, key, path string, value interface{}, options *JSONSetArgsOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONStrAppend(ctx context.Context, key, path, value string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONStrLen(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONToggle(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) JSONType(ctx context.Context, key, path string) *JSONSliceCmd {
	_ = "STUB: not implemented"
	return nil
}
