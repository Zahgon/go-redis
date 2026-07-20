package redis

import (
	"context"
	"time"
)

const XTrimLimitDisabled = -1

func appendXTrimLimit(args []interface{}, limit int64) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

type StreamCmdable interface {
	XAdd(ctx context.Context, a *XAddArgs) *StringCmd
	XAckDel(ctx context.Context, stream string, group string, mode string, ids ...string) *SliceCmd
	XDel(ctx context.Context, stream string, ids ...string) *IntCmd
	XDelEx(ctx context.Context, stream string, mode string, ids ...string) *SliceCmd
	XLen(ctx context.Context, stream string) *IntCmd
	XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd
	XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd
	XRevRange(ctx context.Context, stream string, start, stop string) *XMessageSliceCmd
	XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *XMessageSliceCmd
	XRead(ctx context.Context, a *XReadArgs) *XStreamSliceCmd
	XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd
	XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd
	XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd
	XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd
	XGroupDestroy(ctx context.Context, stream, group string) *IntCmd
	XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
	XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
	XReadGroup(ctx context.Context, a *XReadGroupArgs) *XStreamSliceCmd
	XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd
	XNack(ctx context.Context, a *XNackArgs) *IntCmd
	XPending(ctx context.Context, stream, group string) *XPendingCmd
	XPendingExt(ctx context.Context, a *XPendingExtArgs) *XPendingExtCmd
	XClaim(ctx context.Context, a *XClaimArgs) *XMessageSliceCmd
	XClaimJustID(ctx context.Context, a *XClaimArgs) *StringSliceCmd
	XAutoClaim(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimCmd
	XAutoClaimWithDeleted(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimWithDeletedCmd
	XAutoClaimJustID(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimJustIDCmd
	XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd
	XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd
	XTrimMaxLenMode(ctx context.Context, key string, maxLen int64, mode string) *IntCmd
	XTrimMaxLenApproxMode(ctx context.Context, key string, maxLen, limit int64, mode string) *IntCmd
	XTrimMinID(ctx context.Context, key string, minID string) *IntCmd
	XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd
	XTrimMinIDMode(ctx context.Context, key string, minID string, mode string) *IntCmd
	XTrimMinIDApproxMode(ctx context.Context, key string, minID string, limit int64, mode string) *IntCmd
	XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd
	XInfoStream(ctx context.Context, key string) *XInfoStreamCmd
	XInfoStreamFull(ctx context.Context, key string, count int) *XInfoStreamFullCmd
	XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd
	XCfgSet(ctx context.Context, a *XCfgSetArgs) *StatusCmd
}

type XAddArgs struct {
	Stream     string
	NoMkStream bool
	MaxLen     int64
	MinID      string

	Approx bool

	Limit          int64
	Mode           string
	ID             string
	Values         interface{}
	ProducerID     string
	IdempotentID   string
	IdempotentAuto bool
}

func (c cmdable) XAdd(ctx context.Context, a *XAddArgs) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XAckDel(ctx context.Context, stream string, group string, mode string, ids ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XDel(ctx context.Context, stream string, ids ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XDelEx(ctx context.Context, stream string, mode string, ids ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XLen(ctx context.Context, stream string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XRevRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XRevRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type XReadArgs struct {
	Streams []string
	Count   int64
	Block   time.Duration
	ID      string
}

func (c cmdable) XRead(ctx context.Context, a *XReadArgs) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XGroupDestroy(ctx context.Context, stream, group string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

type XReadGroupArgs struct {
	Group    string
	Consumer string
	Streams  []string
	Count    int64
	Block    time.Duration
	NoAck    bool
	Claim    time.Duration
}

func (c cmdable) XReadGroup(ctx context.Context, a *XReadGroupArgs) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

const (
	XNackModeSilent = "SILENT"
	XNackModeFail   = "FAIL"
	XNackModeFatal  = "FATAL"
)

type XNackArgs struct {
	Stream string
	Group  string

	Mode string

	IDs []string

	RetryCount *uint64

	Force bool
}

func (c cmdable) XNack(ctx context.Context, a *XNackArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XPending(ctx context.Context, stream, group string) *XPendingCmd {
	_ = "STUB: not implemented"
	return nil
}

type XPendingExtArgs struct {
	Stream   string
	Group    string
	Idle     time.Duration
	Start    string
	End      string
	Count    int64
	Consumer string
}

func (c cmdable) XPendingExt(ctx context.Context, a *XPendingExtArgs) *XPendingExtCmd {
	_ = "STUB: not implemented"
	return nil
}

type XAutoClaimArgs struct {
	Stream   string
	Group    string
	MinIdle  time.Duration
	Start    string
	Count    int64
	Consumer string
}

func (c cmdable) XAutoClaim(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XAutoClaimWithDeleted(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimWithDeletedCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XAutoClaimJustID(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimJustIDCmd {
	_ = "STUB: not implemented"
	return nil
}

func xAutoClaimArgs(ctx context.Context, a *XAutoClaimArgs) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

type XClaimArgs struct {
	Stream   string
	Group    string
	Consumer string
	MinIdle  time.Duration
	Messages []string
}

func (c cmdable) XClaim(ctx context.Context, a *XClaimArgs) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XClaimJustID(ctx context.Context, a *XClaimArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func xClaimArgs(a *XClaimArgs) []interface{} { _ = "STUB: not implemented"; return nil }

func (c cmdable) xTrim(
	ctx context.Context, key, strategy string,
	approx bool, threshold interface{}, limit int64,
) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XTrimMinID(ctx context.Context, key string, minID string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) xTrimMode(
	ctx context.Context, key, strategy string,
	approx bool, threshold interface{}, limit int64,
	mode string,
) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XTrimMaxLenMode(ctx context.Context, key string, maxLen int64, mode string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XTrimMaxLenApproxMode(ctx context.Context, key string, maxLen, limit int64, mode string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XTrimMinIDMode(ctx context.Context, key string, minID string, mode string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XTrimMinIDApproxMode(ctx context.Context, key string, minID string, limit int64, mode string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XInfoStream(ctx context.Context, key string) *XInfoStreamCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) XInfoStreamFull(ctx context.Context, key string, count int) *XInfoStreamFullCmd {
	_ = "STUB: not implemented"
	return nil
}

type XCfgSetArgs struct {
	Stream   string
	Duration int64
	MaxSize  int64
}

func (c cmdable) XCfgSet(ctx context.Context, a *XCfgSetArgs) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}
