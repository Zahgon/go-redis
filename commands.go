package redis

import (
	"context"
	"io"
	"reflect"
	"time"
)

const KeepTTL = -1

func usePrecise(dur time.Duration) bool { _ = "STUB: not implemented"; return false }

func formatMs(ctx context.Context, dur time.Duration) int64 { _ = "STUB: not implemented"; return 0 }

func formatSec(ctx context.Context, dur time.Duration) int64 { _ = "STUB: not implemented"; return 0 }

func appendArgs(dst, src []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func appendArg(dst []interface{}, arg interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func appendStructField(dst []interface{}, v reflect.Value) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func omitEmpty(opt string) bool { _ = "STUB: not implemented"; return false }

func isEmptyValue(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

type Cmdable interface {
	Pipeline() Pipeliner
	Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error)

	TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error)
	TxPipeline() Pipeliner

	Command(ctx context.Context) *CommandsInfoCmd
	CommandList(ctx context.Context, filter *FilterBy) *StringSliceCmd
	CommandGetKeys(ctx context.Context, commands ...interface{}) *StringSliceCmd
	CommandGetKeysAndFlags(ctx context.Context, commands ...interface{}) *KeyFlagsCmd
	ClientGetName(ctx context.Context) *StringCmd
	Echo(ctx context.Context, message interface{}) *StringCmd
	Ping(ctx context.Context) *StatusCmd
	Quit(ctx context.Context) *StatusCmd
	Unlink(ctx context.Context, keys ...string) *IntCmd

	BgRewriteAOF(ctx context.Context) *StatusCmd
	BgSave(ctx context.Context) *StatusCmd
	ClientKill(ctx context.Context, ipPort string) *StatusCmd
	ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd
	ClientList(ctx context.Context) *StringCmd
	ClientInfo(ctx context.Context) *ClientInfoCmd
	ClientPause(ctx context.Context, dur time.Duration) *BoolCmd
	ClientUnpause(ctx context.Context) *BoolCmd
	ClientID(ctx context.Context) *IntCmd
	ClientUnblock(ctx context.Context, id int64) *IntCmd
	ClientUnblockWithError(ctx context.Context, id int64) *IntCmd
	ClientMaintNotifications(ctx context.Context, enabled bool, endpointType string) *StatusCmd
	ConfigGet(ctx context.Context, parameter string) *MapStringStringCmd
	ConfigResetStat(ctx context.Context) *StatusCmd
	ConfigSet(ctx context.Context, parameter, value string) *StatusCmd
	ConfigRewrite(ctx context.Context) *StatusCmd
	DBSize(ctx context.Context) *IntCmd
	FlushAll(ctx context.Context) *StatusCmd
	FlushAllAsync(ctx context.Context) *StatusCmd
	FlushDB(ctx context.Context) *StatusCmd
	FlushDBAsync(ctx context.Context) *StatusCmd
	Info(ctx context.Context, section ...string) *StringCmd
	InfoMap(ctx context.Context, section ...string) *InfoCmd
	LastSave(ctx context.Context) *IntCmd
	Save(ctx context.Context) *StatusCmd
	Shutdown(ctx context.Context) *StatusCmd
	ShutdownSave(ctx context.Context) *StatusCmd
	ShutdownNoSave(ctx context.Context) *StatusCmd
	SlaveOf(ctx context.Context, host, port string) *StatusCmd
	ReplicaOf(ctx context.Context, host, port string) *StatusCmd
	SlowLogGet(ctx context.Context, num int64) *SlowLogCmd
	SlowLogLen(ctx context.Context) *IntCmd
	SlowLogReset(ctx context.Context) *StatusCmd
	Time(ctx context.Context) *TimeCmd
	DebugObject(ctx context.Context, key string) *StringCmd
	MemoryUsage(ctx context.Context, key string, samples ...int) *IntCmd
	Latency(ctx context.Context) *LatencyCmd
	LatencyReset(ctx context.Context, events ...interface{}) *StatusCmd

	ModuleLoadex(ctx context.Context, conf *ModuleLoadexConfig) *StringCmd

	ACLCmdable
	ArrayCmdable
	BitMapCmdable
	ClusterCmdable
	GenericCmdable
	GeoCmdable
	HashCmdable
	HyperLogLogCmdable
	ListCmdable
	ProbabilisticCmdable
	PubSubCmdable
	ScriptingFunctionsCmdable
	SearchCmdable
	SetCmdable
	SortedSetCmdable
	StringCmdable
	StreamCmdable
	TimeseriesCmdable
	JSONCmdable
	VectorSetCmdable
}

type StatefulCmdable interface {
	Cmdable
	Auth(ctx context.Context, password string) *StatusCmd
	AuthACL(ctx context.Context, username, password string) *StatusCmd
	Select(ctx context.Context, index int) *StatusCmd
	SwapDB(ctx context.Context, index1, index2 int) *StatusCmd
	ClientSetName(ctx context.Context, name string) *BoolCmd
	ClientSetInfo(ctx context.Context, info LibraryInfo) *StatusCmd
	Hello(ctx context.Context, ver int, username, password, clientName string) *MapStringInterfaceCmd
}

var (
	_ Cmdable = (*Client)(nil)
	_ Cmdable = (*Tx)(nil)
	_ Cmdable = (*Ring)(nil)
	_ Cmdable = (*ClusterClient)(nil)
	_ Cmdable = (*Pipeline)(nil)
)

type cmdable func(ctx context.Context, cmd Cmder) error

type statefulCmdable func(ctx context.Context, cmd Cmder) error

func (c statefulCmdable) Auth(ctx context.Context, password string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c statefulCmdable) AuthACL(ctx context.Context, username, password string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Wait(ctx context.Context, numSlaves int, timeout time.Duration) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) WaitAOF(ctx context.Context, numLocal, numSlaves int, timeout time.Duration) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c statefulCmdable) Select(ctx context.Context, index int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c statefulCmdable) SwapDB(ctx context.Context, index1, index2 int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c statefulCmdable) ClientSetName(ctx context.Context, name string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c statefulCmdable) ClientSetInfo(ctx context.Context, info LibraryInfo) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (info LibraryInfo) Validate() error { _ = "STUB: not implemented"; return nil }

func (c statefulCmdable) Hello(ctx context.Context,
	ver int, username, password, clientName string,
) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Command(ctx context.Context) *CommandsInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

type FilterBy struct {
	Module  string
	ACLCat  string
	Pattern string
}

func (c cmdable) CommandList(ctx context.Context, filter *FilterBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CommandGetKeys(ctx context.Context, commands ...interface{}) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CommandGetKeysAndFlags(ctx context.Context, commands ...interface{}) *KeyFlagsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClientGetName(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Echo(ctx context.Context, message interface{}) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Ping(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) Do(ctx context.Context, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) DoRaw(ctx context.Context, args ...interface{}) *RawCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) DoRawWriteTo(ctx context.Context, w io.Writer, args ...interface{}) *RawWriteToCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Quit(_ context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) BgRewriteAOF(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BgSave(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ClientKill(ctx context.Context, ipPort string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClientList(ctx context.Context) *StringCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ClientPause(ctx context.Context, dur time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClientUnpause(ctx context.Context) *BoolCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ClientID(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ClientUnblock(ctx context.Context, id int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClientUnblockWithError(ctx context.Context, id int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClientInfo(ctx context.Context) *ClientInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClientMaintNotifications(ctx context.Context, enabled bool, endpointType string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ConfigGet(ctx context.Context, parameter string) *MapStringStringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ConfigResetStat(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ConfigSet(ctx context.Context, parameter, value string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ConfigRewrite(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) DBSize(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) FlushAll(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) FlushAllAsync(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FlushDB(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) FlushDBAsync(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Info(ctx context.Context, sections ...string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) InfoMap(ctx context.Context, sections ...string) *InfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LastSave(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) Save(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) shutdown(ctx context.Context, modifier string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Shutdown(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ShutdownSave(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ShutdownNoSave(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SlaveOf(ctx context.Context, host, port string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ReplicaOf(ctx context.Context, host, port string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SlowLogGet(ctx context.Context, num int64) *SlowLogCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SlowLogLen(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) SlowLogReset(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Latency(ctx context.Context) *LatencyCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) LatencyReset(ctx context.Context, events ...interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Sync(_ context.Context) { _ = "STUB: not implemented"; return }

func (c cmdable) Time(ctx context.Context) *TimeCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) DebugObject(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) MemoryUsage(ctx context.Context, key string, samples ...int) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

type ModuleLoadexConfig struct {
	Path string
	Conf map[string]interface{}
	Args []interface{}
}

func (c *ModuleLoadexConfig) toArgs() []interface{} { _ = "STUB: not implemented"; return nil }

func (c cmdable) ModuleLoadex(ctx context.Context, conf *ModuleLoadexConfig) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Monitor(ctx context.Context, ch chan string) *MonitorCmd {
	_ = "STUB: not implemented"
	return nil
}
