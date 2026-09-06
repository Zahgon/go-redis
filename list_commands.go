package redis

import (
	"context"
	"time"
)

type ListCmdable interface {
	BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
	BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *KeyValuesCmd
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd
	LIndex(ctx context.Context, key string, index int64) *StringCmd
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *IntCmd
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *IntCmd
	LLen(ctx context.Context, key string) *IntCmd
	LMPop(ctx context.Context, direction string, count int64, keys ...string) *KeyValuesCmd
	LPop(ctx context.Context, key string) *StringCmd
	LPopCount(ctx context.Context, key string, count int) *StringSliceCmd
	LPos(ctx context.Context, key string, value string, args LPosArgs) *IntCmd
	LPosCount(ctx context.Context, key string, value string, count int64, args LPosArgs) *IntSliceCmd
	LPush(ctx context.Context, key string, values ...interface{}) *IntCmd
	LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd
	LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd
	LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd
	RPop(ctx context.Context, key string) *StringCmd
	RPopCount(ctx context.Context, key string, count int) *StringSliceCmd
	RPopLPush(ctx context.Context, source, destination string) *StringCmd
	RPush(ctx context.Context, key string, values ...interface{}) *IntCmd
	RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
	LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd
	BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *StringCmd
	LMoveM(ctx context.Context, source, destination, srcpos, destpos string, args LMoveMArgs) *StringSliceCmd
	BLMoveM(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration, args LMoveMArgs) *StringSliceCmd
}

func (c cmdable) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *KeyValuesCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LIndex(ctx context.Context, key string, index int64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LMPop(ctx context.Context, direction string, count int64, keys ...string) *KeyValuesCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LPop(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LPopCount(ctx context.Context, key string, count int) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type LPosArgs struct {
	Rank, MaxLen int64
}

type LMoveMMode string

const (
	LMoveMCount   LMoveMMode = "COUNT"
	LMoveMExactly LMoveMMode = "EXACTLY"
)

type LMoveMOrder string

const (
	LMoveMOBO  LMoveMOrder = "OBO"
	LMoveMBulk LMoveMOrder = "BULK"
)

type LMoveMArgs struct {
	Mode  LMoveMMode
	Count int64
	Order LMoveMOrder
}

func (a LMoveMArgs) appendArgs(args []interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LPos(ctx context.Context, key string, value string, a LPosArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LPosCount(ctx context.Context, key string, value string, count int64, a LPosArgs) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LPush(ctx context.Context, key string, values ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) RPop(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) RPopCount(ctx context.Context, key string, count int) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) RPopLPush(ctx context.Context, source, destination string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) RPush(ctx context.Context, key string, values ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BLMove(
	ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration,
) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) LMoveM(ctx context.Context, source, destination, srcpos, destpos string, a LMoveMArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BLMoveM(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration, a LMoveMArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}
