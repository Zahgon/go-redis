package redis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9/internal/proto"
)

type TimeseriesCmdable interface {
	TSAdd(ctx context.Context, key string, timestamp interface{}, value float64) *IntCmd
	TSAddWithArgs(ctx context.Context, key string, timestamp interface{}, value float64, options *TSOptions) *IntCmd
	TSCreate(ctx context.Context, key string) *StatusCmd
	TSCreateWithArgs(ctx context.Context, key string, options *TSOptions) *StatusCmd
	TSAlter(ctx context.Context, key string, options *TSAlterOptions) *StatusCmd
	TSCreateRule(ctx context.Context, sourceKey string, destKey string, aggregator Aggregator, bucketDuration int) *StatusCmd
	TSCreateRuleWithArgs(ctx context.Context, sourceKey string, destKey string, aggregator Aggregator, bucketDuration int, options *TSCreateRuleOptions) *StatusCmd
	TSIncrBy(ctx context.Context, Key string, timestamp float64) *IntCmd
	TSIncrByWithArgs(ctx context.Context, key string, timestamp float64, options *TSIncrDecrOptions) *IntCmd
	TSDecrBy(ctx context.Context, Key string, timestamp float64) *IntCmd
	TSDecrByWithArgs(ctx context.Context, key string, timestamp float64, options *TSIncrDecrOptions) *IntCmd
	TSDel(ctx context.Context, Key string, fromTimestamp int, toTimestamp int) *IntCmd
	TSDeleteRule(ctx context.Context, sourceKey string, destKey string) *StatusCmd
	TSGet(ctx context.Context, key string) *TSTimestampValueCmd
	TSGetWithArgs(ctx context.Context, key string, options *TSGetOptions) *TSTimestampValueCmd
	TSInfo(ctx context.Context, key string) *MapStringInterfaceCmd
	TSInfoWithArgs(ctx context.Context, key string, options *TSInfoOptions) *MapStringInterfaceCmd
	TSMAdd(ctx context.Context, ktvSlices [][]interface{}) *IntSliceCmd
	TSQueryIndex(ctx context.Context, filterExpr []string) *StringSliceCmd
	TSRevRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *TSTimestampValueSliceCmd
	TSRevRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *TSRevRangeOptions) *TSTimestampValueSliceCmd
	TSRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *TSTimestampValueSliceCmd
	TSRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *TSRangeOptions) *TSTimestampValueSliceCmd
	TSMRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *MapStringSliceInterfaceCmd
	TSMRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *TSMRangeOptions) *MapStringSliceInterfaceCmd
	TSMRevRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *MapStringSliceInterfaceCmd
	TSMRevRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *TSMRevRangeOptions) *MapStringSliceInterfaceCmd
	TSMGet(ctx context.Context, filters []string) *MapStringSliceInterfaceCmd
	TSMGetWithArgs(ctx context.Context, filters []string, options *TSMGetOptions) *MapStringSliceInterfaceCmd
}

type TSOptions struct {
	Retention         int
	ChunkSize         int
	Encoding          string
	DuplicatePolicy   string
	Labels            map[string]string
	IgnoreMaxTimeDiff int64
	IgnoreMaxValDiff  float64
}
type TSIncrDecrOptions struct {
	Timestamp         int64
	Retention         int
	ChunkSize         int
	Uncompressed      bool
	DuplicatePolicy   string
	Labels            map[string]string
	IgnoreMaxTimeDiff int64
	IgnoreMaxValDiff  float64
}

type TSAlterOptions struct {
	Retention         int
	ChunkSize         int
	DuplicatePolicy   string
	Labels            map[string]string
	IgnoreMaxTimeDiff int64
	IgnoreMaxValDiff  float64
}

type TSCreateRuleOptions struct {
	alignTimestamp int64
}

type TSGetOptions struct {
	Latest bool
}

type TSInfoOptions struct {
	Debug bool
}
type Aggregator int

const (
	Invalid = Aggregator(iota)
	Avg
	Sum
	Min
	Max
	Range
	Count
	First
	Last
	StdP
	StdS
	VarP
	VarS
	Twa
	CountNaN
	CountAll
)

func (a Aggregator) String() string { _ = "STUB: not implemented"; return "" }

var (
	errTSMultiAggregationGroupBy = errors.New("redis: GROUPBY is not allowed when multiple aggregators are specified")
	errTSAggregationConflict     = errors.New("redis: setting both Aggregator and Aggregators is not allowed; use Aggregators instead because Aggregator is deprecated")
)

func formatAggregationArgs(aggregator Aggregator, aggregators []Aggregator) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func formatAggregatorArg(aggregator Aggregator) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type TSRangeOptions struct {
	Latest        bool
	FilterByTS    []int
	FilterByValue []int
	Count         int
	Align         interface{}

	Aggregator      Aggregator
	Aggregators     []Aggregator
	BucketDuration  int
	BucketTimestamp interface{}
	Empty           bool
}

type TSRevRangeOptions struct {
	Latest        bool
	FilterByTS    []int
	FilterByValue []int
	Count         int
	Align         interface{}

	Aggregator      Aggregator
	Aggregators     []Aggregator
	BucketDuration  int
	BucketTimestamp interface{}
	Empty           bool
}

type TSMRangeOptions struct {
	Latest         bool
	FilterByTS     []int
	FilterByValue  []int
	WithLabels     bool
	SelectedLabels []interface{}
	Count          int
	Align          interface{}

	Aggregator      Aggregator
	Aggregators     []Aggregator
	BucketDuration  int
	BucketTimestamp interface{}
	Empty           bool
	GroupByLabel    interface{}
	Reducer         interface{}
}

type TSMRevRangeOptions struct {
	Latest         bool
	FilterByTS     []int
	FilterByValue  []int
	WithLabels     bool
	SelectedLabels []interface{}
	Count          int
	Align          interface{}

	Aggregator      Aggregator
	Aggregators     []Aggregator
	BucketDuration  int
	BucketTimestamp interface{}
	Empty           bool
	GroupByLabel    interface{}
	Reducer         interface{}
}

type TSMGetOptions struct {
	Latest         bool
	WithLabels     bool
	SelectedLabels []interface{}
}

func (c cmdable) TSAdd(ctx context.Context, key string, timestamp interface{}, value float64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSAddWithArgs(ctx context.Context, key string, timestamp interface{}, value float64, options *TSOptions) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSCreate(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSCreateWithArgs(ctx context.Context, key string, options *TSOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSAlter(ctx context.Context, key string, options *TSAlterOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSCreateRule(ctx context.Context, sourceKey string, destKey string, aggregator Aggregator, bucketDuration int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSCreateRuleWithArgs(ctx context.Context, sourceKey string, destKey string, aggregator Aggregator, bucketDuration int, options *TSCreateRuleOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSIncrBy(ctx context.Context, Key string, timestamp float64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSIncrByWithArgs(ctx context.Context, key string, timestamp float64, options *TSIncrDecrOptions) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSDecrBy(ctx context.Context, Key string, timestamp float64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSDecrByWithArgs(ctx context.Context, key string, timestamp float64, options *TSIncrDecrOptions) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSDel(ctx context.Context, Key string, fromTimestamp int, toTimestamp int) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSDeleteRule(ctx context.Context, sourceKey string, destKey string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSGetWithArgs(ctx context.Context, key string, options *TSGetOptions) *TSTimestampValueCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSGet(ctx context.Context, key string) *TSTimestampValueCmd {
	_ = "STUB: not implemented"
	return nil
}

type TSTimestampValue struct {
	Timestamp int64
	Value     float64
	Values    []float64
}

func (tv TSTimestampValue) String() string { _ = "STUB: not implemented"; return "" }

type TSTimestampValueCmd struct {
	baseCmd
	val TSTimestampValue
}

func newTSTimestampValueCmd(ctx context.Context, args ...interface{}) *TSTimestampValueCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *TSTimestampValueCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *TSTimestampValueCmd) SetVal(val TSTimestampValue) { _ = "STUB: not implemented"; return }

func (cmd *TSTimestampValueCmd) Result() (TSTimestampValue, error) {
	_ = "STUB: not implemented"
	return *new(TSTimestampValue), nil
}

func (cmd *TSTimestampValueCmd) Val() TSTimestampValue {
	_ = "STUB: not implemented"
	return *new(TSTimestampValue)
}

func (cmd *TSTimestampValueCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *TSTimestampValueCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) TSInfo(ctx context.Context, key string) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSInfoWithArgs(ctx context.Context, key string, options *TSInfoOptions) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSMAdd(ctx context.Context, ktvSlices [][]interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSQueryIndex(ctx context.Context, filterExpr []string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSRevRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSRevRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *TSRevRangeOptions) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *TSRangeOptions) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type TSTimestampValueSliceCmd struct {
	baseCmd
	val []TSTimestampValue
}

func newTSTimestampValueSliceCmd(ctx context.Context, args ...interface{}) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *TSTimestampValueSliceCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *TSTimestampValueSliceCmd) SetVal(val []TSTimestampValue) {
	_ = "STUB: not implemented"
	return
}

func (cmd *TSTimestampValueSliceCmd) Result() ([]TSTimestampValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *TSTimestampValueSliceCmd) Val() []TSTimestampValue {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *TSTimestampValueSliceCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *TSTimestampValueSliceCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) TSMRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSMRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *TSMRangeOptions) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSMRevRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSMRevRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *TSMRevRangeOptions) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSMGet(ctx context.Context, filters []string) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TSMGetWithArgs(ctx context.Context, filters []string, options *TSMGetOptions) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}
