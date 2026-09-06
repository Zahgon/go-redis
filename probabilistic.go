package redis

import (
	"context"

	"github.com/redis/go-redis/v9/internal/proto"
)

type ProbabilisticCmdable interface {
	BFAdd(ctx context.Context, key string, element interface{}) *BoolCmd
	BFCard(ctx context.Context, key string) *IntCmd
	BFExists(ctx context.Context, key string, element interface{}) *BoolCmd
	BFInfo(ctx context.Context, key string) *BFInfoCmd
	BFInfoArg(ctx context.Context, key, option string) *BFInfoCmd
	BFInfoCapacity(ctx context.Context, key string) *BFInfoCmd
	BFInfoSize(ctx context.Context, key string) *BFInfoCmd
	BFInfoFilters(ctx context.Context, key string) *BFInfoCmd
	BFInfoItems(ctx context.Context, key string) *BFInfoCmd
	BFInfoExpansion(ctx context.Context, key string) *BFInfoCmd
	BFInsert(ctx context.Context, key string, options *BFInsertOptions, elements ...interface{}) *BoolSliceCmd
	BFMAdd(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd
	BFMExists(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd
	BFReserve(ctx context.Context, key string, errorRate float64, capacity int64) *StatusCmd
	BFReserveExpansion(ctx context.Context, key string, errorRate float64, capacity, expansion int64) *StatusCmd
	BFReserveNonScaling(ctx context.Context, key string, errorRate float64, capacity int64) *StatusCmd
	BFReserveWithArgs(ctx context.Context, key string, options *BFReserveOptions) *StatusCmd
	BFScanDump(ctx context.Context, key string, iterator int64) *ScanDumpCmd
	BFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *StatusCmd

	CFAdd(ctx context.Context, key string, element interface{}) *BoolCmd
	CFAddNX(ctx context.Context, key string, element interface{}) *BoolCmd
	CFCount(ctx context.Context, key string, element interface{}) *IntCmd
	CFDel(ctx context.Context, key string, element interface{}) *BoolCmd
	CFExists(ctx context.Context, key string, element interface{}) *BoolCmd
	CFInfo(ctx context.Context, key string) *CFInfoCmd
	CFInsert(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *BoolSliceCmd
	CFInsertNX(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *IntSliceCmd
	CFMExists(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd
	CFReserve(ctx context.Context, key string, capacity int64) *StatusCmd
	CFReserveWithArgs(ctx context.Context, key string, options *CFReserveOptions) *StatusCmd
	CFReserveExpansion(ctx context.Context, key string, capacity int64, expansion int64) *StatusCmd
	CFReserveBucketSize(ctx context.Context, key string, capacity int64, bucketsize int64) *StatusCmd
	CFReserveMaxIterations(ctx context.Context, key string, capacity int64, maxiterations int64) *StatusCmd
	CFScanDump(ctx context.Context, key string, iterator int64) *ScanDumpCmd
	CFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *StatusCmd

	CMSIncrBy(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd
	CMSInfo(ctx context.Context, key string) *CMSInfoCmd
	CMSInitByDim(ctx context.Context, key string, width, height int64) *StatusCmd
	CMSInitByProb(ctx context.Context, key string, errorRate, probability float64) *StatusCmd
	CMSMerge(ctx context.Context, destKey string, sourceKeys ...string) *StatusCmd
	CMSMergeWithWeight(ctx context.Context, destKey string, sourceKeys map[string]int64) *StatusCmd
	CMSQuery(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd

	TopKAdd(ctx context.Context, key string, elements ...interface{}) *StringSliceCmd
	TopKCount(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd
	TopKIncrBy(ctx context.Context, key string, elements ...interface{}) *StringSliceCmd
	TopKInfo(ctx context.Context, key string) *TopKInfoCmd
	TopKList(ctx context.Context, key string) *StringSliceCmd
	TopKListWithCount(ctx context.Context, key string) *MapStringIntCmd
	TopKQuery(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd
	TopKReserve(ctx context.Context, key string, k int64) *StatusCmd
	TopKReserveWithOptions(ctx context.Context, key string, k int64, width, depth int64, decay float64) *StatusCmd

	TDigestAdd(ctx context.Context, key string, elements ...float64) *StatusCmd
	TDigestByRank(ctx context.Context, key string, rank ...uint64) *FloatSliceCmd
	TDigestByRevRank(ctx context.Context, key string, rank ...uint64) *FloatSliceCmd
	TDigestCDF(ctx context.Context, key string, elements ...float64) *FloatSliceCmd
	TDigestCreate(ctx context.Context, key string) *StatusCmd
	TDigestCreateWithCompression(ctx context.Context, key string, compression int64) *StatusCmd
	TDigestInfo(ctx context.Context, key string) *TDigestInfoCmd
	TDigestMax(ctx context.Context, key string) *FloatCmd
	TDigestMin(ctx context.Context, key string) *FloatCmd
	TDigestMerge(ctx context.Context, destKey string, options *TDigestMergeOptions, sourceKeys ...string) *StatusCmd
	TDigestQuantile(ctx context.Context, key string, elements ...float64) *FloatSliceCmd
	TDigestRank(ctx context.Context, key string, values ...float64) *IntSliceCmd
	TDigestReset(ctx context.Context, key string) *StatusCmd
	TDigestRevRank(ctx context.Context, key string, values ...float64) *IntSliceCmd
	TDigestTrimmedMean(ctx context.Context, key string, lowCutQuantile, highCutQuantile float64) *FloatCmd
}

type BFInsertOptions struct {
	Capacity   int64
	Error      float64
	Expansion  int64
	NonScaling bool
	NoCreate   bool
}

type BFReserveOptions struct {
	Capacity   int64
	Error      float64
	Expansion  int64
	NonScaling bool
}

type CFReserveOptions struct {
	Capacity      int64
	BucketSize    int64
	MaxIterations int64
	Expansion     int64
}

type CFInsertOptions struct {
	Capacity int64
	NoCreate bool
}

func (c cmdable) BFReserve(ctx context.Context, key string, errorRate float64, capacity int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFReserveExpansion(ctx context.Context, key string, errorRate float64, capacity, expansion int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFReserveNonScaling(ctx context.Context, key string, errorRate float64, capacity int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFReserveWithArgs(ctx context.Context, key string, options *BFReserveOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFAdd(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFExists(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFScanDump(ctx context.Context, key string, iterator int64) *ScanDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

type ScanDump struct {
	Iter int64
	Data string
}

type ScanDumpCmd struct {
	baseCmd

	val ScanDump
}

func newScanDumpCmd(ctx context.Context, args ...interface{}) *ScanDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ScanDumpCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *ScanDumpCmd) SetVal(val ScanDump) { _ = "STUB: not implemented"; return }

func (cmd *ScanDumpCmd) Result() (ScanDump, error) {
	_ = "STUB: not implemented"
	return *new(ScanDump), nil
}

func (cmd *ScanDumpCmd) Val() ScanDump { _ = "STUB: not implemented"; return *new(ScanDump) }

func (cmd *ScanDumpCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ScanDumpCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) BFInfo(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

type BFInfo struct {
	Capacity      int64
	Size          int64
	Filters       int64
	ItemsInserted int64
	ExpansionRate int64
}

type BFInfoCmd struct {
	baseCmd

	val BFInfo
}

func NewBFInfoCmd(ctx context.Context, args ...interface{}) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *BFInfoCmd) SetVal(val BFInfo) { _ = "STUB: not implemented"; return }

func (cmd *BFInfoCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *BFInfoCmd) Val() BFInfo { _ = "STUB: not implemented"; return *new(BFInfo) }

func (cmd *BFInfoCmd) Result() (BFInfo, error) { _ = "STUB: not implemented"; return *new(BFInfo), nil }

func (cmd *BFInfoCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *BFInfoCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) BFInfoCapacity(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFInfoSize(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFInfoFilters(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFInfoItems(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFInfoExpansion(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFInfoArg(ctx context.Context, key, option string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFInsert(ctx context.Context, key string, options *BFInsertOptions, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFMAdd(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BFMExists(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFReserve(ctx context.Context, key string, capacity int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFReserveExpansion(ctx context.Context, key string, capacity int64, expansion int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFReserveBucketSize(ctx context.Context, key string, capacity int64, bucketsize int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFReserveMaxIterations(ctx context.Context, key string, capacity int64, maxiterations int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFReserveWithArgs(ctx context.Context, key string, options *CFReserveOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFAdd(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFAddNX(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFCount(ctx context.Context, key string, element interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFDel(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFExists(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFScanDump(ctx context.Context, key string, iterator int64) *ScanDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

type CFInfo struct {
	Size             int64
	NumBuckets       int64
	NumFilters       int64
	NumItemsInserted int64
	NumItemsDeleted  int64
	BucketSize       int64
	ExpansionRate    int64
	MaxIteration     int64
}

type CFInfoCmd struct {
	baseCmd

	val CFInfo
}

func NewCFInfoCmd(ctx context.Context, args ...interface{}) *CFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *CFInfoCmd) SetVal(val CFInfo) { _ = "STUB: not implemented"; return }

func (cmd *CFInfoCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *CFInfoCmd) Val() CFInfo { _ = "STUB: not implemented"; return *new(CFInfo) }

func (cmd *CFInfoCmd) Result() (CFInfo, error) { _ = "STUB: not implemented"; return *new(CFInfo), nil }

func (cmd *CFInfoCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *CFInfoCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) CFInfo(ctx context.Context, key string) *CFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFInsert(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFInsertNX(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) getCfInsertWithArgs(args []interface{}, options *CFInsertOptions, elements ...interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CFMExists(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CMSIncrBy(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type CMSInfo struct {
	Width int64
	Depth int64
	Count int64
}

type CMSInfoCmd struct {
	baseCmd

	val CMSInfo
}

func NewCMSInfoCmd(ctx context.Context, args ...interface{}) *CMSInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *CMSInfoCmd) SetVal(val CMSInfo) { _ = "STUB: not implemented"; return }

func (cmd *CMSInfoCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *CMSInfoCmd) Val() CMSInfo { _ = "STUB: not implemented"; return *new(CMSInfo) }

func (cmd *CMSInfoCmd) Result() (CMSInfo, error) {
	_ = "STUB: not implemented"
	return *new(CMSInfo), nil
}

func (cmd *CMSInfoCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *CMSInfoCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) CMSInfo(ctx context.Context, key string) *CMSInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CMSInitByDim(ctx context.Context, key string, width, depth int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CMSInitByProb(ctx context.Context, key string, errorRate, probability float64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CMSMerge(ctx context.Context, destKey string, sourceKeys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CMSMergeWithWeight(ctx context.Context, destKey string, sourceKeys map[string]int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) CMSQuery(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TopKAdd(ctx context.Context, key string, elements ...interface{}) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TopKReserve(ctx context.Context, key string, k int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TopKReserveWithOptions(ctx context.Context, key string, k int64, width, depth int64, decay float64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

type TopKInfo struct {
	K     int64
	Width int64
	Depth int64
	Decay float64
}

type TopKInfoCmd struct {
	baseCmd

	val TopKInfo
}

func NewTopKInfoCmd(ctx context.Context, args ...interface{}) *TopKInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *TopKInfoCmd) SetVal(val TopKInfo) { _ = "STUB: not implemented"; return }

func (cmd *TopKInfoCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *TopKInfoCmd) Val() TopKInfo { _ = "STUB: not implemented"; return *new(TopKInfo) }

func (cmd *TopKInfoCmd) Result() (TopKInfo, error) {
	_ = "STUB: not implemented"
	return *new(TopKInfo), nil
}

func (cmd *TopKInfoCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *TopKInfoCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) TopKInfo(ctx context.Context, key string) *TopKInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TopKQuery(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TopKCount(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TopKIncrBy(ctx context.Context, key string, elements ...interface{}) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TopKList(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TopKListWithCount(ctx context.Context, key string) *MapStringIntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestAdd(ctx context.Context, key string, elements ...float64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestByRank(ctx context.Context, key string, rank ...uint64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestByRevRank(ctx context.Context, key string, rank ...uint64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestCDF(ctx context.Context, key string, elements ...float64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestCreate(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestCreateWithCompression(ctx context.Context, key string, compression int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

type TDigestInfo struct {
	Compression       int64
	Capacity          int64
	MergedNodes       int64
	UnmergedNodes     int64
	MergedWeight      int64
	UnmergedWeight    int64
	Observations      int64
	TotalCompressions int64
	MemoryUsage       int64
}

type TDigestInfoCmd struct {
	baseCmd

	val TDigestInfo
}

func NewTDigestInfoCmd(ctx context.Context, args ...interface{}) *TDigestInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *TDigestInfoCmd) SetVal(val TDigestInfo) { _ = "STUB: not implemented"; return }

func (cmd *TDigestInfoCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *TDigestInfoCmd) Val() TDigestInfo { _ = "STUB: not implemented"; return *new(TDigestInfo) }

func (cmd *TDigestInfoCmd) Result() (TDigestInfo, error) {
	_ = "STUB: not implemented"
	return *new(TDigestInfo), nil
}

func (cmd *TDigestInfoCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *TDigestInfoCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) TDigestInfo(ctx context.Context, key string) *TDigestInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestMax(ctx context.Context, key string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

type TDigestMergeOptions struct {
	Compression int64
	Override    bool
}

func (c cmdable) TDigestMerge(ctx context.Context, destKey string, options *TDigestMergeOptions, sourceKeys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestMin(ctx context.Context, key string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestQuantile(ctx context.Context, key string, elements ...float64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestRank(ctx context.Context, key string, values ...float64) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestReset(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestRevRank(ctx context.Context, key string, values ...float64) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TDigestTrimmedMean(ctx context.Context, key string, lowCutQuantile, highCutQuantile float64) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}
