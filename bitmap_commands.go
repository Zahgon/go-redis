package redis

import (
	"context"
)

type BitMapCmdable interface {
	GetBit(ctx context.Context, key string, offset int64) *IntCmd
	SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd
	BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd
	BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpDiff(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpDiff1(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpAndOr(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpOne(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpNot(ctx context.Context, destKey string, key string) *IntCmd
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd
	BitPosSpan(ctx context.Context, key string, bit int8, start, end int64, span string) *IntCmd
	BitField(ctx context.Context, key string, values ...interface{}) *IntSliceCmd
	BitFieldRO(ctx context.Context, key string, values ...interface{}) *IntSliceCmd
}

func (c cmdable) GetBit(ctx context.Context, key string, offset int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

type BitCount struct {
	Start, End int64
	Unit       string
}

const BitCountIndexByte string = "BYTE"
const BitCountIndexBit string = "BIT"

func (c cmdable) BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) bitOp(ctx context.Context, op, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitOpNot(ctx context.Context, destKey string, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitOpDiff(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitOpDiff1(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitOpAndOr(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitOpOne(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitPosSpan(ctx context.Context, key string, bit int8, start, end int64, span string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitField(ctx context.Context, key string, values ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BitFieldRO(ctx context.Context, key string, values ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}
