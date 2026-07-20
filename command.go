package redis

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/redis/go-redis/v9/internal"
	"github.com/redis/go-redis/v9/internal/proto"
	"github.com/redis/go-redis/v9/internal/routing"
)

var keylessCommands = map[string]struct{}{
	"acl":          {},
	"asking":       {},
	"auth":         {},
	"bgrewriteaof": {},
	"bgsave":       {},
	"client":       {},
	"cluster":      {},
	"config":       {},
	"debug":        {},
	"discard":      {},
	"echo":         {},
	"exec":         {},
	"failover":     {},
	"function":     {},
	"hello":        {},
	"hotkeys":      {},
	"latency":      {},
	"lolwut":       {},
	"module":       {},
	"monitor":      {},
	"multi":        {},
	"pfselftest":   {},
	"ping":         {},
	"psubscribe":   {},
	"psync":        {},
	"publish":      {},
	"pubsub":       {},
	"punsubscribe": {},
	"quit":         {},
	"readonly":     {},
	"readwrite":    {},
	"replconf":     {},
	"replicaof":    {},
	"role":         {},
	"save":         {},
	"script":       {},
	"select":       {},
	"shutdown":     {},
	"slaveof":      {},
	"slowlog":      {},
	"subscribe":    {},
	"swapdb":       {},
	"sync":         {},
	"time":         {},
	"unsubscribe":  {},
	"unwatch":      {},
	"wait":         {},
}

type CmdTyper interface {
	GetCmdType() CmdType
}

type CmdTypeGetter interface {
	GetCmdType() CmdType
}

type CmdType uint8

const (
	CmdTypeGeneric CmdType = iota
	CmdTypeString
	CmdTypeInt
	CmdTypeBool
	CmdTypeFloat
	CmdTypeStringSlice
	CmdTypeIntSlice
	CmdTypeFloatSlice
	CmdTypeBoolSlice
	CmdTypeMapStringString
	CmdTypeMapStringInt
	CmdTypeMapStringInterface
	CmdTypeMapStringInterfaceSlice
	CmdTypeSlice
	CmdTypeStatus
	CmdTypeDuration
	CmdTypeTime
	CmdTypeKeyValueSlice
	CmdTypeStringStructMap
	CmdTypeXMessageSlice
	CmdTypeXStreamSlice
	CmdTypeXPending
	CmdTypeXPendingExt
	CmdTypeXAutoClaim
	CmdTypeXAutoClaimWithDeleted
	CmdTypeXAutoClaimJustID
	CmdTypeXInfoConsumers
	CmdTypeXInfoGroups
	CmdTypeXInfoStream
	CmdTypeXInfoStreamFull
	CmdTypeZSlice
	CmdTypeZWithKey
	CmdTypeScan
	CmdTypeClusterSlots
	CmdTypeGeoLocation
	CmdTypeGeoSearchLocation
	CmdTypeGeoPos
	CmdTypeCommandsInfo
	CmdTypeSlowLog
	CmdTypeMapStringStringSlice
	CmdTypeMapMapStringInterface
	CmdTypeKeyValues
	CmdTypeZSliceWithKey
	CmdTypeFunctionList
	CmdTypeFunctionStats
	CmdTypeLCS
	CmdTypeKeyFlags
	CmdTypeClusterLinks
	CmdTypeClusterShards
	CmdTypeRankWithScore
	CmdTypeClientInfo
	CmdTypeACLLog
	CmdTypeInfo
	CmdTypeMonitor
	CmdTypeJSON
	CmdTypeJSONSlice
	CmdTypeIntPointerSlice
	CmdTypeScanDump
	CmdTypeBFInfo
	CmdTypeCFInfo
	CmdTypeCMSInfo
	CmdTypeTopKInfo
	CmdTypeTDigestInfo
	CmdTypeFTSynDump
	CmdTypeAggregate
	CmdTypeFTInfo
	CmdTypeFTSpellCheck
	CmdTypeFTSearch
	CmdTypeTSTimestampValue
	CmdTypeTSTimestampValueSlice
	CmdTypeHotKeys
	CmdTypeIncrEXInt
	CmdTypeIncrEXFloat
	CmdTypeUint
	CmdTypeUintSlice
	CmdTypeAREntrySlice
)

type (
	CmdTypeXAutoClaimValue struct {
		messages []XMessage
		start    string
	}

	CmdTypeXAutoClaimWithDeletedValue struct {
		messages   []XMessage
		start      string
		deletedIDs []string
	}

	CmdTypeXAutoClaimJustIDValue struct {
		ids   []string
		start string
	}

	CmdTypeScanValue struct {
		keys   []string
		cursor uint64
	}

	CmdTypeKeyValuesValue struct {
		key    string
		values []string
	}

	CmdTypeZSliceWithKeyValue struct {
		key    string
		zSlice []Z
	}
)

type Cmder interface {
	Name() string

	FullName() string

	Args() []interface{}

	String() string

	Clone() Cmder

	stringArg(int) string
	firstKeyPos() int8
	SetFirstKeyPos(int8)
	stepCount() int8
	SetStepCount(int8)

	readTimeout() *time.Duration
	readReply(rd *proto.Reader) error
	readRawReply(rd *proto.Reader) error
	SetErr(error)
	Err() error

	NoRetry() bool

	GetCmdType() CmdType
}

func setCmdsErr(cmds []Cmder, e error) { _ = "STUB: not implemented"; return }

func cmdsFirstErr(cmds []Cmder) error { _ = "STUB: not implemented"; return nil }

func cmdsContainNoRetry(cmds []Cmder) bool { _ = "STUB: not implemented"; return false }

func writeCmds(wr *proto.Writer, cmds []Cmder) error { _ = "STUB: not implemented"; return nil }

func writeCmd(wr *proto.Writer, cmd Cmder) error { _ = "STUB: not implemented"; return nil }

func cmdFirstKeyPosWithInfo(cmd Cmder, info *CommandInfo) int { _ = "STUB: not implemented"; return 0 }

func cmdString(cmd Cmder, val interface{}) string { _ = "STUB: not implemented"; return "" }

type baseCmd struct {
	ctx          context.Context
	args         []interface{}
	err          error
	keyPos       int8
	_stepCount   int8
	rawVal       interface{}
	_readTimeout *time.Duration
	cmdType      CmdType
}

var _ Cmder = (*Cmd)(nil)

func (cmd *baseCmd) Name() string { _ = "STUB: not implemented"; return "" }

func (cmd *baseCmd) FullName() string { _ = "STUB: not implemented"; return "" }

func (cmd *baseCmd) Args() []interface{} { _ = "STUB: not implemented"; return nil }

func (cmd *baseCmd) stringArg(pos int) string { _ = "STUB: not implemented"; return "" }

func (cmd *baseCmd) firstKeyPos() int8 { _ = "STUB: not implemented"; return 0 }

func (cmd *baseCmd) SetFirstKeyPos(keyPos int8) { _ = "STUB: not implemented"; return }

func (cmd *baseCmd) stepCount() int8 { _ = "STUB: not implemented"; return 0 }

func (cmd *baseCmd) SetStepCount(stepCount int8) { _ = "STUB: not implemented"; return }

func (cmd *baseCmd) SetErr(e error) { _ = "STUB: not implemented"; return }

func (cmd *baseCmd) Err() error { _ = "STUB: not implemented"; return nil }

func (cmd *baseCmd) readTimeout() *time.Duration { _ = "STUB: not implemented"; return nil }

func (cmd *baseCmd) setReadTimeout(d time.Duration) { _ = "STUB: not implemented"; return }

func (cmd *baseCmd) readRawReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *baseCmd) NoRetry() bool { _ = "STUB: not implemented"; return false }

func (cmd *baseCmd) GetCmdType() CmdType { _ = "STUB: not implemented"; return *new(CmdType) }

func (cmd *baseCmd) cloneBaseCmd() baseCmd { _ = "STUB: not implemented"; return *new(baseCmd) }

type Cmd struct {
	baseCmd

	val interface{}
}

func NewCmd(ctx context.Context, args ...interface{}) *Cmd { _ = "STUB: not implemented"; return nil }

func (cmd *Cmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *Cmd) SetVal(val interface{}) { _ = "STUB: not implemented"; return }

func (cmd *Cmd) Val() interface{} { _ = "STUB: not implemented"; return nil }

func (cmd *Cmd) Result() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) Text() (string, error) { _ = "STUB: not implemented"; return "", nil }

func toString(val interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cmd *Cmd) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *Cmd) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func toInt64(val interface{}) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *Cmd) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func toUint64(val interface{}) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *Cmd) Float32() (float32, error) { _ = "STUB: not implemented"; return 0, nil }

func toFloat32(val interface{}) (float32, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *Cmd) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func toFloat64(val interface{}) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *Cmd) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func toBool(val interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (cmd *Cmd) Slice() ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) StringSlice() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) Int64Slice() ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) Uint64Slice() ([]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) Float32Slice() ([]float32, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) Float64Slice() ([]float64, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) BoolSlice() ([]bool, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) readReply(rd *proto.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (cmd *Cmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type RawCmd struct {
	baseCmd
	val []byte
}

var _ Cmder = (*RawCmd)(nil)

func NewRawCmd(ctx context.Context, args ...interface{}) *RawCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *RawCmd) SetVal(val []byte) { _ = "STUB: not implemented"; return }

func (cmd *RawCmd) Val() []byte { _ = "STUB: not implemented"; return nil }

func (cmd *RawCmd) Result() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *RawCmd) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *RawCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *RawCmd) readReply(rd *proto.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (cmd *RawCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type RawWriteToCmd struct {
	baseCmd
	w       io.Writer
	written int64
}

var _ Cmder = (*RawWriteToCmd)(nil)

func NewRawWriteToCmd(ctx context.Context, w io.Writer, args ...interface{}) *RawWriteToCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *RawWriteToCmd) SetVal(written int64) { _ = "STUB: not implemented"; return }

func (cmd *RawWriteToCmd) Val() int64 { _ = "STUB: not implemented"; return 0 }

func (cmd *RawWriteToCmd) Result() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *RawWriteToCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *RawWriteToCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *RawWriteToCmd) NoRetry() bool { _ = "STUB: not implemented"; return false }

func (cmd *RawWriteToCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type ZeroCopyStringCmd struct {
	baseCmd
	buf    []byte
	n      int
	cloned bool
}

var _ Cmder = (*ZeroCopyStringCmd)(nil)

func NewZeroCopyStringCmd(ctx context.Context, buf []byte, args ...interface{}) *ZeroCopyStringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ZeroCopyStringCmd) SetVal(n int) { _ = "STUB: not implemented"; return }

func (cmd *ZeroCopyStringCmd) Val() int { _ = "STUB: not implemented"; return 0 }

func (cmd *ZeroCopyStringCmd) Result() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *ZeroCopyStringCmd) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (cmd *ZeroCopyStringCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ZeroCopyStringCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ZeroCopyStringCmd) NoRetry() bool { _ = "STUB: not implemented"; return false }

func (cmd *ZeroCopyStringCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type SliceCmd struct {
	baseCmd

	val []interface{}
}

var _ Cmder = (*SliceCmd)(nil)

func NewSliceCmd(ctx context.Context, args ...interface{}) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *SliceCmd) SetVal(val []interface{}) { _ = "STUB: not implemented"; return }

func (cmd *SliceCmd) Val() []interface{} { _ = "STUB: not implemented"; return nil }

func (cmd *SliceCmd) Result() ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *SliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *SliceCmd) Scan(dst interface{}) error { _ = "STUB: not implemented"; return nil }

func (cmd *SliceCmd) readReply(rd *proto.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (cmd *SliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type StatusCmd struct {
	baseCmd

	val string
}

var _ Cmder = (*StatusCmd)(nil)

func NewStatusCmd(ctx context.Context, args ...interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *StatusCmd) SetVal(val string) { _ = "STUB: not implemented"; return }

func (cmd *StatusCmd) Val() string { _ = "STUB: not implemented"; return "" }

func (cmd *StatusCmd) Result() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cmd *StatusCmd) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *StatusCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *StatusCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *StatusCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type IntCmd struct {
	baseCmd

	val int64
}

var _ Cmder = (*IntCmd)(nil)

func NewIntCmd(ctx context.Context, args ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *IntCmd) SetVal(val int64) { _ = "STUB: not implemented"; return }

func (cmd *IntCmd) Val() int64 { _ = "STUB: not implemented"; return 0 }

func (cmd *IntCmd) Result() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *IntCmd) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *IntCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *IntCmd) readReply(rd *proto.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (cmd *IntCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type UintCmd struct {
	baseCmd

	val uint64
}

var _ Cmder = (*UintCmd)(nil)

func NewUintCmd(ctx context.Context, args ...any) *UintCmd { _ = "STUB: not implemented"; return nil }

func (cmd *UintCmd) SetVal(val uint64) { _ = "STUB: not implemented"; return }

func (cmd *UintCmd) Val() uint64 { _ = "STUB: not implemented"; return 0 }

func (cmd *UintCmd) Result() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *UintCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *UintCmd) readReply(rd *proto.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (cmd *UintCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type DigestCmd struct {
	baseCmd

	val uint64
}

var _ Cmder = (*DigestCmd)(nil)

func NewDigestCmd(ctx context.Context, args ...interface{}) *DigestCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *DigestCmd) SetVal(val uint64) { _ = "STUB: not implemented"; return }

func (cmd *DigestCmd) Val() uint64 { _ = "STUB: not implemented"; return 0 }

func (cmd *DigestCmd) Result() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *DigestCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *DigestCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (cmd *DigestCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type IntSliceCmd struct {
	baseCmd

	val []int64
}

var _ Cmder = (*IntSliceCmd)(nil)

func NewIntSliceCmd(ctx context.Context, args ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *IntSliceCmd) SetVal(val []int64) { _ = "STUB: not implemented"; return }

func (cmd *IntSliceCmd) Val() []int64 { _ = "STUB: not implemented"; return nil }

func (cmd *IntSliceCmd) Result() ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *IntSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *IntSliceCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *IntSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type UintSliceCmd struct {
	baseCmd

	val []uint64
}

var _ Cmder = (*UintSliceCmd)(nil)

func NewUintSliceCmd(ctx context.Context, args ...any) *UintSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *UintSliceCmd) SetVal(val []uint64) { _ = "STUB: not implemented"; return }

func (cmd *UintSliceCmd) Val() []uint64 { _ = "STUB: not implemented"; return nil }

func (cmd *UintSliceCmd) Result() ([]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *UintSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *UintSliceCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *UintSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type DurationCmd struct {
	baseCmd

	val       time.Duration
	precision time.Duration
}

var _ Cmder = (*DurationCmd)(nil)

func NewDurationCmd(ctx context.Context, precision time.Duration, args ...interface{}) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *DurationCmd) SetVal(val time.Duration) { _ = "STUB: not implemented"; return }

func (cmd *DurationCmd) Val() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (cmd *DurationCmd) Result() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (cmd *DurationCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *DurationCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *DurationCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type TimeCmd struct {
	baseCmd

	val time.Time
}

var _ Cmder = (*TimeCmd)(nil)

func NewTimeCmd(ctx context.Context, args ...interface{}) *TimeCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *TimeCmd) SetVal(val time.Time) { _ = "STUB: not implemented"; return }

func (cmd *TimeCmd) Val() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (cmd *TimeCmd) Result() (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (cmd *TimeCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *TimeCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *TimeCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type BoolCmd struct {
	baseCmd

	val bool
}

var _ Cmder = (*BoolCmd)(nil)

func NewBoolCmd(ctx context.Context, args ...interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *BoolCmd) SetVal(val bool) { _ = "STUB: not implemented"; return }

func (cmd *BoolCmd) Val() bool { _ = "STUB: not implemented"; return false }

func (cmd *BoolCmd) Result() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (cmd *BoolCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *BoolCmd) readReply(rd *proto.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (cmd *BoolCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type StringCmd struct {
	baseCmd

	val string
}

var _ Cmder = (*StringCmd)(nil)

func NewStringCmd(ctx context.Context, args ...interface{}) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *StringCmd) SetVal(val string) { _ = "STUB: not implemented"; return }

func (cmd *StringCmd) Val() string { _ = "STUB: not implemented"; return "" }

func (cmd *StringCmd) Result() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cmd *StringCmd) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *StringCmd) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (cmd *StringCmd) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *StringCmd) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *StringCmd) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *StringCmd) Float32() (float32, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *StringCmd) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *StringCmd) Time() (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (cmd *StringCmd) Scan(val interface{}) error { _ = "STUB: not implemented"; return nil }

func (cmd *StringCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *StringCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *StringCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type FloatCmd struct {
	baseCmd

	val float64
}

var _ Cmder = (*FloatCmd)(nil)

func NewFloatCmd(ctx context.Context, args ...interface{}) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FloatCmd) SetVal(val float64) { _ = "STUB: not implemented"; return }

func (cmd *FloatCmd) Val() float64 { _ = "STUB: not implemented"; return 0 }

func (cmd *FloatCmd) Result() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *FloatCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *FloatCmd) readReply(rd *proto.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (cmd *FloatCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type FloatSliceCmd struct {
	baseCmd

	val []float64
}

var _ Cmder = (*FloatSliceCmd)(nil)

func NewFloatSliceCmd(ctx context.Context, args ...interface{}) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FloatSliceCmd) SetVal(val []float64) { _ = "STUB: not implemented"; return }

func (cmd *FloatSliceCmd) Val() []float64 { _ = "STUB: not implemented"; return nil }

func (cmd *FloatSliceCmd) Result() ([]float64, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *FloatSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *FloatSliceCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *FloatSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type StringSliceCmd struct {
	baseCmd

	val []string
}

var _ Cmder = (*StringSliceCmd)(nil)

func NewStringSliceCmd(ctx context.Context, args ...interface{}) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *StringSliceCmd) SetVal(val []string) { _ = "STUB: not implemented"; return }

func (cmd *StringSliceCmd) Val() []string { _ = "STUB: not implemented"; return nil }

func (cmd *StringSliceCmd) Result() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *StringSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *StringSliceCmd) ScanSlice(container interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *StringSliceCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *StringSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type StringSliceSliceCmd struct {
	baseCmd

	val [][]string
}

var _ Cmder = (*StringSliceSliceCmd)(nil)

func NewStringSliceSliceCmd(ctx context.Context, args ...any) *StringSliceSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *StringSliceSliceCmd) SetVal(val [][]string) { _ = "STUB: not implemented"; return }

func (cmd *StringSliceSliceCmd) Val() [][]string { _ = "STUB: not implemented"; return nil }

func (cmd *StringSliceSliceCmd) Result() ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *StringSliceSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *StringSliceSliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *StringSliceSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type KeyValue struct {
	Key   string
	Value string
}

type KeyValueSliceCmd struct {
	baseCmd

	val []KeyValue
}

var _ Cmder = (*KeyValueSliceCmd)(nil)

func NewKeyValueSliceCmd(ctx context.Context, args ...interface{}) *KeyValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *KeyValueSliceCmd) SetVal(val []KeyValue) { _ = "STUB: not implemented"; return }

func (cmd *KeyValueSliceCmd) Val() []KeyValue { _ = "STUB: not implemented"; return nil }

func (cmd *KeyValueSliceCmd) Result() ([]KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *KeyValueSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *KeyValueSliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *KeyValueSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type BoolSliceCmd struct {
	baseCmd

	val []bool
}

var _ Cmder = (*BoolSliceCmd)(nil)

func NewBoolSliceCmd(ctx context.Context, args ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *BoolSliceCmd) SetVal(val []bool) { _ = "STUB: not implemented"; return }

func (cmd *BoolSliceCmd) Val() []bool { _ = "STUB: not implemented"; return nil }

func (cmd *BoolSliceCmd) Result() ([]bool, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *BoolSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *BoolSliceCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *BoolSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type MapStringStringCmd struct {
	baseCmd

	val map[string]string
}

var _ Cmder = (*MapStringStringCmd)(nil)

func NewMapStringStringCmd(ctx context.Context, args ...interface{}) *MapStringStringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringStringCmd) Val() map[string]string { _ = "STUB: not implemented"; return nil }

func (cmd *MapStringStringCmd) SetVal(val map[string]string) { _ = "STUB: not implemented"; return }

func (cmd *MapStringStringCmd) Result() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *MapStringStringCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *MapStringStringCmd) Scan(dest interface{}) error { _ = "STUB: not implemented"; return nil }

func (cmd *MapStringStringCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringStringCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type MapStringIntCmd struct {
	baseCmd

	val map[string]int64
}

var _ Cmder = (*MapStringIntCmd)(nil)

func NewMapStringIntCmd(ctx context.Context, args ...interface{}) *MapStringIntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringIntCmd) SetVal(val map[string]int64) { _ = "STUB: not implemented"; return }

func (cmd *MapStringIntCmd) Val() map[string]int64 { _ = "STUB: not implemented"; return nil }

func (cmd *MapStringIntCmd) Result() (map[string]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *MapStringIntCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *MapStringIntCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringIntCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type MapStringSliceInterfaceCmd struct {
	baseCmd
	val map[string][]interface{}
}

func NewMapStringSliceInterfaceCmd(ctx context.Context, args ...interface{}) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringSliceInterfaceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *MapStringSliceInterfaceCmd) SetVal(val map[string][]interface{}) {
	_ = "STUB: not implemented"
	return
}

func (cmd *MapStringSliceInterfaceCmd) Result() (map[string][]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *MapStringSliceInterfaceCmd) Val() map[string][]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringSliceInterfaceCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringSliceInterfaceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type StringStructMapCmd struct {
	baseCmd

	val map[string]struct{}
}

var _ Cmder = (*StringStructMapCmd)(nil)

func NewStringStructMapCmd(ctx context.Context, args ...interface{}) *StringStructMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *StringStructMapCmd) SetVal(val map[string]struct{}) { _ = "STUB: not implemented"; return }

func (cmd *StringStructMapCmd) Val() map[string]struct{} { _ = "STUB: not implemented"; return nil }

func (cmd *StringStructMapCmd) Result() (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *StringStructMapCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *StringStructMapCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *StringStructMapCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type XMessage struct {
	ID     string
	Values map[string]interface{}

	MillisElapsedFromDelivery int64

	DeliveredCount int64
}

type XMessageSliceCmd struct {
	baseCmd

	val []XMessage
}

var _ Cmder = (*XMessageSliceCmd)(nil)

func NewXMessageSliceCmd(ctx context.Context, args ...interface{}) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XMessageSliceCmd) SetVal(val []XMessage) { _ = "STUB: not implemented"; return }

func (cmd *XMessageSliceCmd) Val() []XMessage { _ = "STUB: not implemented"; return nil }

func (cmd *XMessageSliceCmd) Result() ([]XMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *XMessageSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XMessageSliceCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XMessageSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func readXMessageSlice(rd *proto.Reader) ([]XMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readXMessage(rd *proto.Reader) (XMessage, error) {
	_ = "STUB: not implemented"
	return *new(XMessage), nil
}

func stringInterfaceMapParser(rd *proto.Reader) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type XStream struct {
	Stream   string
	Messages []XMessage
}

type XStreamSliceCmd struct {
	baseCmd

	val []XStream
}

var _ Cmder = (*XStreamSliceCmd)(nil)

func NewXStreamSliceCmd(ctx context.Context, args ...interface{}) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XStreamSliceCmd) SetVal(val []XStream) { _ = "STUB: not implemented"; return }

func (cmd *XStreamSliceCmd) Val() []XStream { _ = "STUB: not implemented"; return nil }

func (cmd *XStreamSliceCmd) Result() ([]XStream, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *XStreamSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XStreamSliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XStreamSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type XPending struct {
	Count     int64
	Lower     string
	Higher    string
	Consumers map[string]int64
}

type XPendingCmd struct {
	baseCmd
	val *XPending
}

var _ Cmder = (*XPendingCmd)(nil)

func NewXPendingCmd(ctx context.Context, args ...interface{}) *XPendingCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XPendingCmd) SetVal(val *XPending) { _ = "STUB: not implemented"; return }

func (cmd *XPendingCmd) Val() *XPending { _ = "STUB: not implemented"; return nil }

func (cmd *XPendingCmd) Result() (*XPending, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *XPendingCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XPendingCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *XPendingCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type XPendingExt struct {
	ID         string
	Consumer   string
	Idle       time.Duration
	RetryCount int64
}

type XPendingExtCmd struct {
	baseCmd
	val []XPendingExt
}

var _ Cmder = (*XPendingExtCmd)(nil)

func NewXPendingExtCmd(ctx context.Context, args ...interface{}) *XPendingExtCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XPendingExtCmd) SetVal(val []XPendingExt) { _ = "STUB: not implemented"; return }

func (cmd *XPendingExtCmd) Val() []XPendingExt { _ = "STUB: not implemented"; return nil }

func (cmd *XPendingExtCmd) Result() ([]XPendingExt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *XPendingExtCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XPendingExtCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *XPendingExtCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type XAutoClaimCmd struct {
	baseCmd

	start string
	val   []XMessage
}

var _ Cmder = (*XAutoClaimCmd)(nil)

func NewXAutoClaimCmd(ctx context.Context, args ...interface{}) *XAutoClaimCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XAutoClaimCmd) SetVal(val []XMessage, start string) { _ = "STUB: not implemented"; return }

func (cmd *XAutoClaimCmd) Val() (messages []XMessage, start string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func (cmd *XAutoClaimCmd) Result() (messages []XMessage, start string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (cmd *XAutoClaimCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XAutoClaimCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *XAutoClaimCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type XAutoClaimWithDeletedCmd struct {
	baseCmd

	start      string
	val        []XMessage
	deletedIDs []string
}

var _ Cmder = (*XAutoClaimWithDeletedCmd)(nil)

func NewXAutoClaimWithDeletedCmd(ctx context.Context, args ...interface{}) *XAutoClaimWithDeletedCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XAutoClaimWithDeletedCmd) SetVal(val []XMessage, start string, deletedIDs []string) {
	_ = "STUB: not implemented"
	return
}

func (cmd *XAutoClaimWithDeletedCmd) Val() (messages []XMessage, start string, deletedIDs []string) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (cmd *XAutoClaimWithDeletedCmd) Result() (messages []XMessage, start string, deletedIDs []string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil, nil
}

func (cmd *XAutoClaimWithDeletedCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XAutoClaimWithDeletedCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XAutoClaimWithDeletedCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type XAutoClaimJustIDCmd struct {
	baseCmd

	start string
	val   []string
}

var _ Cmder = (*XAutoClaimJustIDCmd)(nil)

func NewXAutoClaimJustIDCmd(ctx context.Context, args ...interface{}) *XAutoClaimJustIDCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XAutoClaimJustIDCmd) SetVal(val []string, start string) {
	_ = "STUB: not implemented"
	return
}

func (cmd *XAutoClaimJustIDCmd) Val() (ids []string, start string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func (cmd *XAutoClaimJustIDCmd) Result() (ids []string, start string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (cmd *XAutoClaimJustIDCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XAutoClaimJustIDCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XAutoClaimJustIDCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type XInfoConsumersCmd struct {
	baseCmd
	val []XInfoConsumer
}

type XInfoConsumer struct {
	Name     string
	Pending  int64
	Idle     time.Duration
	Inactive time.Duration
}

var _ Cmder = (*XInfoConsumersCmd)(nil)

func NewXInfoConsumersCmd(ctx context.Context, stream string, group string) *XInfoConsumersCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XInfoConsumersCmd) SetVal(val []XInfoConsumer) { _ = "STUB: not implemented"; return }

func (cmd *XInfoConsumersCmd) Val() []XInfoConsumer { _ = "STUB: not implemented"; return nil }

func (cmd *XInfoConsumersCmd) Result() ([]XInfoConsumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *XInfoConsumersCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XInfoConsumersCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XInfoConsumersCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type XInfoGroupsCmd struct {
	baseCmd
	val []XInfoGroup
}

type XInfoGroup struct {
	Name            string
	Consumers       int64
	Pending         int64
	LastDeliveredID string
	EntriesRead     int64

	Lag int64
}

var _ Cmder = (*XInfoGroupsCmd)(nil)

func NewXInfoGroupsCmd(ctx context.Context, stream string) *XInfoGroupsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XInfoGroupsCmd) SetVal(val []XInfoGroup) { _ = "STUB: not implemented"; return }

func (cmd *XInfoGroupsCmd) Val() []XInfoGroup { _ = "STUB: not implemented"; return nil }

func (cmd *XInfoGroupsCmd) Result() ([]XInfoGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *XInfoGroupsCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XInfoGroupsCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *XInfoGroupsCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type XInfoStreamCmd struct {
	baseCmd
	val *XInfoStream
}

type XInfoStream struct {
	Length               int64
	RadixTreeKeys        int64
	RadixTreeNodes       int64
	Groups               int64
	LastGeneratedID      string
	MaxDeletedEntryID    string
	EntriesAdded         int64
	FirstEntry           XMessage
	LastEntry            XMessage
	RecordedFirstEntryID string

	IDMPDuration   int64
	IDMPMaxSize    int64
	PIDsTracked    int64
	IIDsTracked    int64
	IIDsAdded      int64
	IIDsDuplicates int64
}

var _ Cmder = (*XInfoStreamCmd)(nil)

func NewXInfoStreamCmd(ctx context.Context, stream string) *XInfoStreamCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XInfoStreamCmd) SetVal(val *XInfoStream) { _ = "STUB: not implemented"; return }

func (cmd *XInfoStreamCmd) Val() *XInfoStream { _ = "STUB: not implemented"; return nil }

func (cmd *XInfoStreamCmd) Result() (*XInfoStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *XInfoStreamCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XInfoStreamCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *XInfoStreamCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type XInfoStreamFullCmd struct {
	baseCmd
	val *XInfoStreamFull
}

type XInfoStreamFull struct {
	Length               int64
	RadixTreeKeys        int64
	RadixTreeNodes       int64
	LastGeneratedID      string
	MaxDeletedEntryID    string
	EntriesAdded         int64
	Entries              []XMessage
	Groups               []XInfoStreamGroup
	RecordedFirstEntryID string
	IDMPDuration         int64
	IDMPMaxSize          int64
	PIDsTracked          int64
	IIDsTracked          int64
	IIDsAdded            int64
	IIDsDuplicates       int64
}

type XInfoStreamGroup struct {
	Name            string
	LastDeliveredID string
	EntriesRead     int64
	Lag             int64
	PelCount        int64
	NackedCount     uint64
	Pending         []XInfoStreamGroupPending
	Consumers       []XInfoStreamConsumer
}

type XInfoStreamGroupPending struct {
	ID            string
	Consumer      string
	DeliveryTime  time.Time
	DeliveryCount int64
}

type XInfoStreamConsumer struct {
	Name       string
	SeenTime   time.Time
	ActiveTime time.Time
	PelCount   int64
	Pending    []XInfoStreamConsumerPending
}

type XInfoStreamConsumerPending struct {
	ID            string
	DeliveryTime  time.Time
	DeliveryCount int64
}

var _ Cmder = (*XInfoStreamFullCmd)(nil)

func NewXInfoStreamFullCmd(ctx context.Context, args ...interface{}) *XInfoStreamFullCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XInfoStreamFullCmd) SetVal(val *XInfoStreamFull) { _ = "STUB: not implemented"; return }

func (cmd *XInfoStreamFullCmd) Val() *XInfoStreamFull { _ = "STUB: not implemented"; return nil }

func (cmd *XInfoStreamFullCmd) Result() (*XInfoStreamFull, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *XInfoStreamFullCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *XInfoStreamFullCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func readStreamGroups(rd *proto.Reader) ([]XInfoStreamGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readXInfoStreamGroupPending(rd *proto.Reader) ([]XInfoStreamGroupPending, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readXInfoStreamConsumers(rd *proto.Reader) ([]XInfoStreamConsumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *XInfoStreamFullCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type ZSliceCmd struct {
	baseCmd

	val []Z
}

var _ Cmder = (*ZSliceCmd)(nil)

func NewZSliceCmd(ctx context.Context, args ...interface{}) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ZSliceCmd) SetVal(val []Z) { _ = "STUB: not implemented"; return }

func (cmd *ZSliceCmd) Val() []Z { _ = "STUB: not implemented"; return nil }

func (cmd *ZSliceCmd) Result() ([]Z, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *ZSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ZSliceCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *ZSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type ZWithKeyCmd struct {
	baseCmd

	val *ZWithKey
}

var _ Cmder = (*ZWithKeyCmd)(nil)

func NewZWithKeyCmd(ctx context.Context, args ...interface{}) *ZWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ZWithKeyCmd) SetVal(val *ZWithKey) { _ = "STUB: not implemented"; return }

func (cmd *ZWithKeyCmd) Val() *ZWithKey { _ = "STUB: not implemented"; return nil }

func (cmd *ZWithKeyCmd) Result() (*ZWithKey, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *ZWithKeyCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ZWithKeyCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ZWithKeyCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type ScanCmd struct {
	baseCmd

	page   []string
	cursor uint64

	process cmdable
}

var _ Cmder = (*ScanCmd)(nil)

func NewScanCmd(ctx context.Context, process cmdable, args ...interface{}) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ScanCmd) SetVal(page []string, cursor uint64) { _ = "STUB: not implemented"; return }

func (cmd *ScanCmd) Val() (keys []string, cursor uint64) { _ = "STUB: not implemented"; return nil, 0 }

func (cmd *ScanCmd) Result() (keys []string, cursor uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (cmd *ScanCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ScanCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *ScanCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (cmd *ScanCmd) Iterator() *ScanIterator { _ = "STUB: not implemented"; return nil }

type ClusterNode struct {
	ID                 string
	Addr               string
	NetworkingMetadata map[string]string
}

type ClusterSlot struct {
	Start int
	End   int
	Nodes []ClusterNode
}

type ClusterSlotsCmd struct {
	baseCmd

	val []ClusterSlot
}

var _ Cmder = (*ClusterSlotsCmd)(nil)

func NewClusterSlotsCmd(ctx context.Context, args ...interface{}) *ClusterSlotsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ClusterSlotsCmd) SetVal(val []ClusterSlot) { _ = "STUB: not implemented"; return }

func (cmd *ClusterSlotsCmd) Val() []ClusterSlot { _ = "STUB: not implemented"; return nil }

func (cmd *ClusterSlotsCmd) Result() ([]ClusterSlot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *ClusterSlotsCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ClusterSlotsCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ClusterSlotsCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type GeoLocation struct {
	Name                      string
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

type GeoRadiusQuery struct {
	Radius float64

	Unit        string
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
	Count       int

	Sort      string
	Store     string
	StoreDist string

	withLen int
}

type GeoLocationCmd struct {
	baseCmd

	q         *GeoRadiusQuery
	locations []GeoLocation
}

var _ Cmder = (*GeoLocationCmd)(nil)

func NewGeoLocationCmd(ctx context.Context, q *GeoRadiusQuery, args ...interface{}) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func geoLocationArgs(q *GeoRadiusQuery, args ...interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *GeoLocationCmd) SetVal(locations []GeoLocation) { _ = "STUB: not implemented"; return }

func (cmd *GeoLocationCmd) Val() []GeoLocation { _ = "STUB: not implemented"; return nil }

func (cmd *GeoLocationCmd) Result() ([]GeoLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *GeoLocationCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *GeoLocationCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *GeoLocationCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type GeoSearchQuery struct {
	Member string

	Longitude float64
	Latitude  float64

	Radius     float64
	RadiusUnit string

	BoxWidth  float64
	BoxHeight float64
	BoxUnit   string

	Sort     string
	Count    int
	CountAny bool
}

type GeoSearchLocationQuery struct {
	GeoSearchQuery

	WithCoord bool
	WithDist  bool
	WithHash  bool
}

type GeoSearchStoreQuery struct {
	GeoSearchQuery

	StoreDist bool
}

func geoSearchLocationArgs(q *GeoSearchLocationQuery, args []interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func geoSearchArgs(q *GeoSearchQuery, args []interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

type GeoSearchLocationCmd struct {
	baseCmd

	opt *GeoSearchLocationQuery
	val []GeoLocation
}

var _ Cmder = (*GeoSearchLocationCmd)(nil)

func NewGeoSearchLocationCmd(
	ctx context.Context, opt *GeoSearchLocationQuery, args ...interface{},
) *GeoSearchLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *GeoSearchLocationCmd) SetVal(val []GeoLocation) { _ = "STUB: not implemented"; return }

func (cmd *GeoSearchLocationCmd) Val() []GeoLocation { _ = "STUB: not implemented"; return nil }

func (cmd *GeoSearchLocationCmd) Result() ([]GeoLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *GeoSearchLocationCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *GeoSearchLocationCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *GeoSearchLocationCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type GeoPos struct {
	Longitude, Latitude float64
}

type GeoPosCmd struct {
	baseCmd

	val []*GeoPos
}

var _ Cmder = (*GeoPosCmd)(nil)

func NewGeoPosCmd(ctx context.Context, args ...interface{}) *GeoPosCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *GeoPosCmd) SetVal(val []*GeoPos) { _ = "STUB: not implemented"; return }

func (cmd *GeoPosCmd) Val() []*GeoPos { _ = "STUB: not implemented"; return nil }

func (cmd *GeoPosCmd) Result() ([]*GeoPos, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *GeoPosCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *GeoPosCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *GeoPosCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type CommandInfo struct {
	Name          string
	Arity         int8
	Flags         []string
	ACLFlags      []string
	FirstKeyPos   int8
	LastKeyPos    int8
	StepCount     int8
	ReadOnly      bool
	CommandPolicy *routing.CommandPolicy
}

type CommandsInfoCmd struct {
	baseCmd

	val map[string]*CommandInfo
}

var _ Cmder = (*CommandsInfoCmd)(nil)

func NewCommandsInfoCmd(ctx context.Context, args ...interface{}) *CommandsInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *CommandsInfoCmd) SetVal(val map[string]*CommandInfo) { _ = "STUB: not implemented"; return }

func (cmd *CommandsInfoCmd) Val() map[string]*CommandInfo { _ = "STUB: not implemented"; return nil }

func (cmd *CommandsInfoCmd) Result() (map[string]*CommandInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *CommandsInfoCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *CommandsInfoCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *CommandsInfoCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type cmdsInfoCache struct {
	fn func(ctx context.Context) (map[string]*CommandInfo, error)

	once        internal.Once
	refreshLock sync.RWMutex
	cmds        map[string]*CommandInfo
}

func newCmdsInfoCache(fn func(ctx context.Context) (map[string]*CommandInfo, error)) *cmdsInfoCache {
	_ = "STUB: not implemented"
	return nil
}

func (c *cmdsInfoCache) Get(ctx context.Context) (map[string]*CommandInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cmdsInfoCache) Refresh() { _ = "STUB: not implemented"; return }

func (c *cmdsInfoCache) Peek() map[string]*CommandInfo { _ = "STUB: not implemented"; return nil }

const (
	requestPolicy  = "request_policy"
	responsePolicy = "response_policy"
)

func parseCommandPolicies(commandInfoTips map[string]string, firstKeyPos int8) *routing.CommandPolicy {
	_ = "STUB: not implemented"
	return nil
}

type SlowLog struct {
	ID       int64
	Time     time.Time
	Duration time.Duration
	Args     []string

	ClientAddr string
	ClientName string

	CommandArgc int64
}

type SlowLogCmd struct {
	baseCmd

	val []SlowLog
}

var _ Cmder = (*SlowLogCmd)(nil)

func NewSlowLogCmd(ctx context.Context, args ...interface{}) *SlowLogCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *SlowLogCmd) SetVal(val []SlowLog) { _ = "STUB: not implemented"; return }

func (cmd *SlowLogCmd) Val() []SlowLog { _ = "STUB: not implemented"; return nil }

func (cmd *SlowLogCmd) Result() ([]SlowLog, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *SlowLogCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *SlowLogCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *SlowLogCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type Latency struct {
	Name   string
	Time   time.Time
	Latest time.Duration
	Max    time.Duration
}

type LatencyCmd struct {
	baseCmd
	val []Latency
}

var _ Cmder = (*LatencyCmd)(nil)

func NewLatencyCmd(ctx context.Context, args ...interface{}) *LatencyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *LatencyCmd) SetVal(val []Latency) { _ = "STUB: not implemented"; return }

func (cmd *LatencyCmd) Val() []Latency { _ = "STUB: not implemented"; return nil }

func (cmd *LatencyCmd) Result() ([]Latency, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *LatencyCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *LatencyCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *LatencyCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type HotKeysSlotRange []int64

type HotKeysKeyEntry struct {
	Key   string
	Value interface{}
}

type HotKeysResult struct {
	TrackingActive                       bool
	SampleRatio                          uint8
	SelectedSlots                        []HotKeysSlotRange
	SampledCommandsSelectedSlots         time.Duration
	AllCommandsSelectedSlots             time.Duration
	AllCommandsAllSlots                  time.Duration
	NetBytesSampledCommandsSelectedSlots int64
	NetBytesAllCommandsSelectedSlots     int64
	NetBytesAllCommandsAllSlots          int64
	CollectionStartTime                  time.Time
	CollectionDuration                   time.Duration
	UsedCPUSys                           time.Duration
	UsedCPUUser                          time.Duration
	TotalNetBytes                        int64
	ByCPUTime                            []HotKeysKeyEntry
	ByNetBytes                           []HotKeysKeyEntry
}

type HotKeysCmd struct {
	baseCmd

	val *HotKeysResult
}

var _ Cmder = (*HotKeysCmd)(nil)

func NewHotKeysCmd(ctx context.Context, args ...interface{}) *HotKeysCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *HotKeysCmd) SetVal(val *HotKeysResult) { _ = "STUB: not implemented"; return }

func (cmd *HotKeysCmd) Val() *HotKeysResult { _ = "STUB: not implemented"; return nil }

func (cmd *HotKeysCmd) Result() (*HotKeysResult, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *HotKeysCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *HotKeysCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func parseHotKeysKeyEntries(v []interface{}) []HotKeysKeyEntry {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *HotKeysCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type MapStringInterfaceCmd struct {
	baseCmd

	val map[string]interface{}
}

var _ Cmder = (*MapStringInterfaceCmd)(nil)

func NewMapStringInterfaceCmd(ctx context.Context, args ...interface{}) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringInterfaceCmd) SetVal(val map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func (cmd *MapStringInterfaceCmd) Val() map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringInterfaceCmd) Result() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *MapStringInterfaceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *MapStringInterfaceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringInterfaceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type MapStringStringSliceCmd struct {
	baseCmd

	val []map[string]string
}

var _ Cmder = (*MapStringStringSliceCmd)(nil)

func NewMapStringStringSliceCmd(ctx context.Context, args ...interface{}) *MapStringStringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringStringSliceCmd) SetVal(val []map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (cmd *MapStringStringSliceCmd) Val() []map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringStringSliceCmd) Result() ([]map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *MapStringStringSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *MapStringStringSliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringStringSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type MapMapStringInterfaceCmd struct {
	baseCmd
	val map[string]interface{}
}

func NewMapMapStringInterfaceCmd(ctx context.Context, args ...interface{}) *MapMapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapMapStringInterfaceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *MapMapStringInterfaceCmd) SetVal(val map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func (cmd *MapMapStringInterfaceCmd) Result() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *MapMapStringInterfaceCmd) Val() map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapMapStringInterfaceCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapMapStringInterfaceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type MapStringInterfaceSliceCmd struct {
	baseCmd

	val []map[string]interface{}
}

var _ Cmder = (*MapStringInterfaceSliceCmd)(nil)

func NewMapStringInterfaceSliceCmd(ctx context.Context, args ...interface{}) *MapStringInterfaceSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringInterfaceSliceCmd) SetVal(val []map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func (cmd *MapStringInterfaceSliceCmd) Val() []map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringInterfaceSliceCmd) Result() ([]map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *MapStringInterfaceSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *MapStringInterfaceSliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MapStringInterfaceSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type KeyValuesCmd struct {
	baseCmd

	key string
	val []string
}

var _ Cmder = (*KeyValuesCmd)(nil)

func NewKeyValuesCmd(ctx context.Context, args ...interface{}) *KeyValuesCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *KeyValuesCmd) SetVal(key string, val []string) { _ = "STUB: not implemented"; return }

func (cmd *KeyValuesCmd) Val() (string, []string) { _ = "STUB: not implemented"; return "", nil }

func (cmd *KeyValuesCmd) Result() (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (cmd *KeyValuesCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *KeyValuesCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *KeyValuesCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type ZSliceWithKeyCmd struct {
	baseCmd

	key string
	val []Z
}

var _ Cmder = (*ZSliceWithKeyCmd)(nil)

func NewZSliceWithKeyCmd(ctx context.Context, args ...interface{}) *ZSliceWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ZSliceWithKeyCmd) SetVal(key string, val []Z) { _ = "STUB: not implemented"; return }

func (cmd *ZSliceWithKeyCmd) Val() (string, []Z) { _ = "STUB: not implemented"; return "", nil }

func (cmd *ZSliceWithKeyCmd) Result() (string, []Z, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (cmd *ZSliceWithKeyCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ZSliceWithKeyCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ZSliceWithKeyCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type Function struct {
	Name        string
	Description string
	Flags       []string
}

type Library struct {
	Name      string
	Engine    string
	Functions []Function
	Code      string
}

type FunctionListCmd struct {
	baseCmd

	val []Library
}

var _ Cmder = (*FunctionListCmd)(nil)

func NewFunctionListCmd(ctx context.Context, args ...interface{}) *FunctionListCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FunctionListCmd) SetVal(val []Library) { _ = "STUB: not implemented"; return }

func (cmd *FunctionListCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *FunctionListCmd) Val() []Library { _ = "STUB: not implemented"; return nil }

func (cmd *FunctionListCmd) Result() ([]Library, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *FunctionListCmd) First() (*Library, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *FunctionListCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FunctionListCmd) readFunctions(rd *proto.Reader) ([]Function, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FunctionListCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type FunctionStats struct {
	Engines   []Engine
	isRunning bool
	rs        RunningScript
	allrs     []RunningScript
}

func (fs *FunctionStats) Running() bool { _ = "STUB: not implemented"; return false }

func (fs *FunctionStats) RunningScript() (RunningScript, bool) {
	_ = "STUB: not implemented"
	return *new(RunningScript), false
}

func (fs *FunctionStats) AllRunningScripts() []RunningScript { _ = "STUB: not implemented"; return nil }

type RunningScript struct {
	Name     string
	Command  []string
	Duration time.Duration
}

type Engine struct {
	Language       string
	LibrariesCount int64
	FunctionsCount int64
}

type FunctionStatsCmd struct {
	baseCmd
	val FunctionStats
}

var _ Cmder = (*FunctionStatsCmd)(nil)

func NewFunctionStatsCmd(ctx context.Context, args ...interface{}) *FunctionStatsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FunctionStatsCmd) SetVal(val FunctionStats) { _ = "STUB: not implemented"; return }

func (cmd *FunctionStatsCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *FunctionStatsCmd) Val() FunctionStats {
	_ = "STUB: not implemented"
	return *new(FunctionStats)
}

func (cmd *FunctionStatsCmd) Result() (FunctionStats, error) {
	_ = "STUB: not implemented"
	return *new(FunctionStats), nil
}

func (cmd *FunctionStatsCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FunctionStatsCmd) readRunningScript(rd *proto.Reader) (RunningScript, bool, error) {
	_ = "STUB: not implemented"
	return *new(RunningScript), false, nil
}

func (cmd *FunctionStatsCmd) readEngines(rd *proto.Reader) ([]Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FunctionStatsCmd) readDuration(rd *proto.Reader) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (cmd *FunctionStatsCmd) readCommand(rd *proto.Reader) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FunctionStatsCmd) readRunningScripts(rd *proto.Reader) ([]RunningScript, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (cmd *FunctionStatsCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type LCSQuery struct {
	Key1         string
	Key2         string
	Len          bool
	Idx          bool
	MinMatchLen  int
	WithMatchLen bool
}

type LCSMatch struct {
	MatchString string
	Matches     []LCSMatchedPosition
	Len         int64
}

type LCSMatchedPosition struct {
	Key1 LCSPosition
	Key2 LCSPosition

	MatchLen int64
}

type LCSPosition struct {
	Start int64
	End   int64
}

type LCSCmd struct {
	baseCmd

	readType uint8
	val      *LCSMatch
}

func NewLCSCmd(ctx context.Context, q *LCSQuery) *LCSCmd { _ = "STUB: not implemented"; return nil }

func (cmd *LCSCmd) SetVal(val *LCSMatch) { _ = "STUB: not implemented"; return }

func (cmd *LCSCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *LCSCmd) Val() *LCSMatch { _ = "STUB: not implemented"; return nil }

func (cmd *LCSCmd) Result() (*LCSMatch, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *LCSCmd) readReply(rd *proto.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (cmd *LCSCmd) readMatchedPositions(rd *proto.Reader) ([]LCSMatchedPosition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *LCSCmd) readPosition(rd *proto.Reader) (pos LCSPosition, err error) {
	_ = "STUB: not implemented"
	return *new(LCSPosition), nil
}

func (cmd *LCSCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type KeyFlags struct {
	Key   string
	Flags []string
}

type KeyFlagsCmd struct {
	baseCmd

	val []KeyFlags
}

var _ Cmder = (*KeyFlagsCmd)(nil)

func NewKeyFlagsCmd(ctx context.Context, args ...interface{}) *KeyFlagsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *KeyFlagsCmd) SetVal(val []KeyFlags) { _ = "STUB: not implemented"; return }

func (cmd *KeyFlagsCmd) Val() []KeyFlags { _ = "STUB: not implemented"; return nil }

func (cmd *KeyFlagsCmd) Result() ([]KeyFlags, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *KeyFlagsCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *KeyFlagsCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *KeyFlagsCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type ClusterLink struct {
	Direction           string
	Node                string
	CreateTime          int64
	Events              string
	SendBufferAllocated int64
	SendBufferUsed      int64
}

type ClusterLinksCmd struct {
	baseCmd

	val []ClusterLink
}

var _ Cmder = (*ClusterLinksCmd)(nil)

func NewClusterLinksCmd(ctx context.Context, args ...interface{}) *ClusterLinksCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ClusterLinksCmd) SetVal(val []ClusterLink) { _ = "STUB: not implemented"; return }

func (cmd *ClusterLinksCmd) Val() []ClusterLink { _ = "STUB: not implemented"; return nil }

func (cmd *ClusterLinksCmd) Result() ([]ClusterLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *ClusterLinksCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ClusterLinksCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ClusterLinksCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type SlotRange struct {
	Start int64
	End   int64
}

type Node struct {
	ID                string
	Endpoint          string
	IP                string
	Hostname          string
	Port              int64
	TLSPort           int64
	Role              string
	ReplicationOffset int64
	Health            string
}

type ClusterShard struct {
	Slots []SlotRange
	Nodes []Node
}

type ClusterShardsCmd struct {
	baseCmd

	val []ClusterShard
}

var _ Cmder = (*ClusterShardsCmd)(nil)

func NewClusterShardsCmd(ctx context.Context, args ...interface{}) *ClusterShardsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ClusterShardsCmd) SetVal(val []ClusterShard) { _ = "STUB: not implemented"; return }

func (cmd *ClusterShardsCmd) Val() []ClusterShard { _ = "STUB: not implemented"; return nil }

func (cmd *ClusterShardsCmd) Result() ([]ClusterShard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *ClusterShardsCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ClusterShardsCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ClusterShardsCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type RankScore struct {
	Rank  int64
	Score float64
}

type RankWithScoreCmd struct {
	baseCmd

	val RankScore
}

var _ Cmder = (*RankWithScoreCmd)(nil)

func NewRankWithScoreCmd(ctx context.Context, args ...interface{}) *RankWithScoreCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *RankWithScoreCmd) SetVal(val RankScore) { _ = "STUB: not implemented"; return }

func (cmd *RankWithScoreCmd) Val() RankScore { _ = "STUB: not implemented"; return *new(RankScore) }

func (cmd *RankWithScoreCmd) Result() (RankScore, error) {
	_ = "STUB: not implemented"
	return *new(RankScore), nil
}

func (cmd *RankWithScoreCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *RankWithScoreCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *RankWithScoreCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type ClientFlags uint64

const (
	ClientSlave            ClientFlags = 1 << 0
	ClientMaster           ClientFlags = 1 << 1
	ClientMonitor          ClientFlags = 1 << 2
	ClientMulti            ClientFlags = 1 << 3
	ClientBlocked          ClientFlags = 1 << 4
	ClientDirtyCAS         ClientFlags = 1 << 5
	ClientCloseAfterReply  ClientFlags = 1 << 6
	ClientUnBlocked        ClientFlags = 1 << 7
	ClientScript           ClientFlags = 1 << 8
	ClientAsking           ClientFlags = 1 << 9
	ClientCloseASAP        ClientFlags = 1 << 10
	ClientUnixSocket       ClientFlags = 1 << 11
	ClientDirtyExec        ClientFlags = 1 << 12
	ClientMasterForceReply ClientFlags = 1 << 13
	ClientForceAOF         ClientFlags = 1 << 14
	ClientForceRepl        ClientFlags = 1 << 15
	ClientPrePSync         ClientFlags = 1 << 16
	ClientReadOnly         ClientFlags = 1 << 17
	ClientPubSub           ClientFlags = 1 << 18
	ClientPreventAOFProp   ClientFlags = 1 << 19
	ClientPreventReplProp  ClientFlags = 1 << 20
	ClientPreventProp      ClientFlags = ClientPreventAOFProp | ClientPreventReplProp
	ClientPendingWrite     ClientFlags = 1 << 21
	ClientReplyOff         ClientFlags = 1 << 22
	ClientReplySkipNext    ClientFlags = 1 << 23
	ClientReplySkip        ClientFlags = 1 << 24
	ClientLuaDebug         ClientFlags = 1 << 25
	ClientLuaDebugSync     ClientFlags = 1 << 26
	ClientModule           ClientFlags = 1 << 27
	ClientProtected        ClientFlags = 1 << 28
	ClientExecutingCommand ClientFlags = 1 << 29

	ClientPendingCommand      ClientFlags = 1 << 30
	ClientTracking            ClientFlags = 1 << 31
	ClientTrackingBrokenRedir ClientFlags = 1 << 32
	ClientTrackingBCAST       ClientFlags = 1 << 33
	ClientTrackingOptIn       ClientFlags = 1 << 34
	ClientTrackingOptOut      ClientFlags = 1 << 35
	ClientTrackingCaching     ClientFlags = 1 << 36
	ClientTrackingNoLoop      ClientFlags = 1 << 37
	ClientInTimeoutTable      ClientFlags = 1 << 38
	ClientProtocolError       ClientFlags = 1 << 39
	ClientCloseAfterCommand   ClientFlags = 1 << 40
	ClientDenyBlocking        ClientFlags = 1 << 41
	ClientReplRDBOnly         ClientFlags = 1 << 42
	ClientNoEvict             ClientFlags = 1 << 43
	ClientAllowOOM            ClientFlags = 1 << 44
	ClientNoTouch             ClientFlags = 1 << 45
	ClientPushing             ClientFlags = 1 << 46
)

type ClientInfo struct {
	ID                 int64
	Addr               string
	LAddr              string
	FD                 int64
	Name               string
	Age                time.Duration
	Idle               time.Duration
	Flags              ClientFlags
	DB                 int
	Sub                int
	PSub               int
	SSub               int
	Multi              int
	Watch              int
	QueryBuf           int
	QueryBufFree       int
	ArgvMem            int
	MultiMem           int
	BufferSize         int
	BufferPeak         int
	OutputBufferLength int
	OutputListLength   int
	OutputMemory       int
	TotalMemory        int
	TotalNetIn         int
	TotalNetOut        int
	TotalCmds          int
	IoThread           int
	Events             string
	LastCmd            string
	User               string
	Redir              int64
	Resp               int
	LibName            string
	LibVer             string
	ReadEvents         uint64
	AvgPipelineLenSum  uint64
	AvgPipelineLenCnt  uint64
}

type ClientInfoCmd struct {
	baseCmd

	val *ClientInfo
}

var _ Cmder = (*ClientInfoCmd)(nil)

func NewClientInfoCmd(ctx context.Context, args ...interface{}) *ClientInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ClientInfoCmd) SetVal(val *ClientInfo) { _ = "STUB: not implemented"; return }

func (cmd *ClientInfoCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ClientInfoCmd) Val() *ClientInfo { _ = "STUB: not implemented"; return nil }

func (cmd *ClientInfoCmd) Result() (*ClientInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *ClientInfoCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseClientInfo(txt string) (info *ClientInfo, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *ClientInfoCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type ACLLogEntry struct {
	Count                int64
	Reason               string
	Context              string
	Object               string
	Username             string
	AgeSeconds           float64
	ClientInfo           *ClientInfo
	EntryID              int64
	TimestampCreated     int64
	TimestampLastUpdated int64
}

type ACLLogCmd struct {
	baseCmd

	val []*ACLLogEntry
}

var _ Cmder = (*ACLLogCmd)(nil)

func NewACLLogCmd(ctx context.Context, args ...interface{}) *ACLLogCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ACLLogCmd) SetVal(val []*ACLLogEntry) { _ = "STUB: not implemented"; return }

func (cmd *ACLLogCmd) Val() []*ACLLogEntry { _ = "STUB: not implemented"; return nil }

func (cmd *ACLLogCmd) Result() ([]*ACLLogEntry, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *ACLLogCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ACLLogCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *ACLLogCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type LibraryInfo struct {
	LibName *string
	LibVer  *string
}

func WithLibraryName(libName string) LibraryInfo {
	_ = "STUB: not implemented"
	return *new(LibraryInfo)
}

func WithLibraryVersion(libVer string) LibraryInfo {
	_ = "STUB: not implemented"
	return *new(LibraryInfo)
}

type InfoCmd struct {
	baseCmd
	val map[string]map[string]string
}

var _ Cmder = (*InfoCmd)(nil)

func NewInfoCmd(ctx context.Context, args ...interface{}) *InfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *InfoCmd) SetVal(val map[string]map[string]string) { _ = "STUB: not implemented"; return }

func (cmd *InfoCmd) Val() map[string]map[string]string { _ = "STUB: not implemented"; return nil }

func (cmd *InfoCmd) Result() (map[string]map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *InfoCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *InfoCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *InfoCmd) Item(section, key string) string { _ = "STUB: not implemented"; return "" }

func (cmd *InfoCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type MonitorStatus int

const (
	monitorStatusIdle MonitorStatus = iota
	monitorStatusStart
	monitorStatusStop
)

type MonitorCmd struct {
	baseCmd
	ch     chan string
	status MonitorStatus
	mu     sync.Mutex
}

func newMonitorCmd(ctx context.Context, ch chan string) *MonitorCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MonitorCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *MonitorCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *MonitorCmd) readMonitor(rd *proto.Reader, cancel context.CancelFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *MonitorCmd) Start() { _ = "STUB: not implemented"; return }

func (cmd *MonitorCmd) Stop() { _ = "STUB: not implemented"; return }

type VectorScoreSliceCmd struct {
	baseCmd

	val []VectorScore
}

var _ Cmder = (*VectorScoreSliceCmd)(nil)

func NewVectorScoreSliceCmd(ctx context.Context, args ...any) *VectorScoreSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewVectorInfoSliceCmd(ctx context.Context, args ...any) *VectorScoreSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *VectorScoreSliceCmd) SetVal(val []VectorScore) { _ = "STUB: not implemented"; return }

func (cmd *VectorScoreSliceCmd) Val() []VectorScore { _ = "STUB: not implemented"; return nil }

func (cmd *VectorScoreSliceCmd) Result() ([]VectorScore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *VectorScoreSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *VectorScoreSliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *VectorScoreSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type VectorScoreSliceSliceCmd struct {
	baseCmd

	val [][]VectorScore
}

var _ Cmder = (*VectorScoreSliceSliceCmd)(nil)

func NewVectorScoreSliceSliceCmd(ctx context.Context, args ...any) *VectorScoreSliceSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *VectorScoreSliceSliceCmd) SetVal(val [][]VectorScore) { _ = "STUB: not implemented"; return }

func (cmd *VectorScoreSliceSliceCmd) Val() [][]VectorScore { _ = "STUB: not implemented"; return nil }

func (cmd *VectorScoreSliceSliceCmd) Result() ([][]VectorScore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *VectorScoreSliceSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *VectorScoreSliceSliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *VectorScoreSliceSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func readVectorAttribStringOrNil(rd *proto.Reader) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type VectorAttribSliceCmd struct {
	baseCmd

	val []VectorAttrib
}

var _ Cmder = (*VectorAttribSliceCmd)(nil)

func NewVectorAttribSliceCmd(ctx context.Context, args ...any) *VectorAttribSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *VectorAttribSliceCmd) SetVal(val []VectorAttrib) { _ = "STUB: not implemented"; return }

func (cmd *VectorAttribSliceCmd) Val() []VectorAttrib { _ = "STUB: not implemented"; return nil }

func (cmd *VectorAttribSliceCmd) Result() ([]VectorAttrib, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *VectorAttribSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *VectorAttribSliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *VectorAttribSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type VectorScoreAttribSliceCmd struct {
	baseCmd

	val []VectorScoreAttrib
}

var _ Cmder = (*VectorScoreAttribSliceCmd)(nil)

func NewVectorScoreAttribSliceCmd(ctx context.Context, args ...any) *VectorScoreAttribSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *VectorScoreAttribSliceCmd) SetVal(val []VectorScoreAttrib) {
	_ = "STUB: not implemented"
	return
}

func (cmd *VectorScoreAttribSliceCmd) Val() []VectorScoreAttrib {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *VectorScoreAttribSliceCmd) Result() ([]VectorScoreAttrib, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *VectorScoreAttribSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *VectorScoreAttribSliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *VectorScoreAttribSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (cmd *MonitorCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func ExtractCommandValue(cmd interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type IncrEXIntResult struct {
	Value            int64
	AppliedIncrement int64
}

type IncrEXIntCmd struct {
	baseCmd

	val IncrEXIntResult
}

var _ Cmder = (*IncrEXIntCmd)(nil)

func NewIncrEXIntCmd(ctx context.Context, args ...interface{}) *IncrEXIntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *IncrEXIntCmd) SetVal(val IncrEXIntResult) { _ = "STUB: not implemented"; return }
func (cmd *IncrEXIntCmd) Val() IncrEXIntResult {
	_ = "STUB: not implemented"
	return *new(IncrEXIntResult)
}
func (cmd *IncrEXIntCmd) Result() (IncrEXIntResult, error) {
	_ = "STUB: not implemented"
	return *new(IncrEXIntResult), nil
}

func (cmd *IncrEXIntCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *IncrEXIntCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *IncrEXIntCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type IncrEXFloatResult struct {
	Value            float64
	AppliedIncrement float64
}

type IncrEXFloatCmd struct {
	baseCmd

	val IncrEXFloatResult
}

var _ Cmder = (*IncrEXFloatCmd)(nil)

func NewIncrEXFloatCmd(ctx context.Context, args ...interface{}) *IncrEXFloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *IncrEXFloatCmd) SetVal(val IncrEXFloatResult) { _ = "STUB: not implemented"; return }
func (cmd *IncrEXFloatCmd) Val() IncrEXFloatResult {
	_ = "STUB: not implemented"
	return *new(IncrEXFloatResult)
}
func (cmd *IncrEXFloatCmd) Result() (IncrEXFloatResult, error) {
	_ = "STUB: not implemented"
	return *new(IncrEXFloatResult), nil
}

func (cmd *IncrEXFloatCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *IncrEXFloatCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (cmd *IncrEXFloatCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type AREntrySliceCmd struct {
	baseCmd
	val []AREntry
}

var _ Cmder = (*AREntrySliceCmd)(nil)

func NewAREntrySliceCmd(ctx context.Context, args ...any) *AREntrySliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *AREntrySliceCmd) SetVal(val []AREntry) { _ = "STUB: not implemented"; return }

func (cmd *AREntrySliceCmd) Val() []AREntry { _ = "STUB: not implemented"; return nil }

func (cmd *AREntrySliceCmd) Result() ([]AREntry, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *AREntrySliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *AREntrySliceCmd) readReply(rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *AREntrySliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }
