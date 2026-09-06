package redis

import (
	"context"
	"time"
)

type StringCmdable interface {
	Append(ctx context.Context, key, value string) *IntCmd
	Decr(ctx context.Context, key string) *IntCmd
	DecrBy(ctx context.Context, key string, decrement int64) *IntCmd
	DelExArgs(ctx context.Context, key string, a DelExArgs) *IntCmd
	Digest(ctx context.Context, key string) *DigestCmd
	Get(ctx context.Context, key string) *StringCmd
	GetRange(ctx context.Context, key string, start, end int64) *StringCmd
	GetSet(ctx context.Context, key string, value interface{}) *StringCmd
	GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd
	GetDel(ctx context.Context, key string) *StringCmd
	GetToBuffer(ctx context.Context, key string, buf []byte) *ZeroCopyStringCmd
	Incr(ctx context.Context, key string) *IntCmd
	IncrBy(ctx context.Context, key string, value int64) *IntCmd
	IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd
	IncrEXInt(ctx context.Context, key string, args IncrEXIntArgs) *IncrEXIntCmd
	IncrEXFloat(ctx context.Context, key string, args IncrEXFloatArgs) *IncrEXFloatCmd
	LCS(ctx context.Context, q *LCSQuery) *LCSCmd
	MGet(ctx context.Context, keys ...string) *SliceCmd
	MSet(ctx context.Context, values ...interface{}) *StatusCmd
	MSetNX(ctx context.Context, values ...interface{}) *BoolCmd
	MSetEX(ctx context.Context, args MSetEXArgs, values ...interface{}) *IntCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *StatusCmd
	SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	SetFromBuffer(ctx context.Context, key string, buf []byte) *StatusCmd
	SetIFEQ(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StatusCmd
	SetIFEQGet(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StringCmd
	SetIFNE(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StatusCmd
	SetIFNEGet(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StringCmd
	SetIFDEQ(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StatusCmd
	SetIFDEQGet(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StringCmd
	SetIFDNE(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StatusCmd
	SetIFDNEGet(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StringCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
	SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd
	StrLen(ctx context.Context, key string) *IntCmd
}

func (c cmdable) Append(ctx context.Context, key, value string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Decr(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) DecrBy(ctx context.Context, key string, decrement int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

type DelExArgs struct {
	Mode string

	MatchValue interface{}

	MatchDigest uint64
}

func (c cmdable) DelExArgs(ctx context.Context, key string, a DelExArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Digest(ctx context.Context, key string) *DigestCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Get(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GetRange(ctx context.Context, key string, start, end int64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GetSet(ctx context.Context, key string, value interface{}) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GetDel(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GetToBuffer(ctx context.Context, key string, buf []byte) *ZeroCopyStringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Incr(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) IncrBy(ctx context.Context, key string, value int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

type IncrEXIntArgs struct {
	By    int64
	HasBy bool

	LBound, UBound       int64
	HasLBound, HasUBound bool

	Saturate bool

	Expiration *ExpirationOption

	ENX bool
}

type IncrEXFloatArgs struct {
	By float64

	LBound, UBound       float64
	HasLBound, HasUBound bool

	Saturate bool

	Expiration *ExpirationOption

	ENX bool
}

func (c cmdable) IncrEXInt(ctx context.Context, key string, a IncrEXIntArgs) *IncrEXIntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) IncrEXFloat(ctx context.Context, key string, a IncrEXFloatArgs) *IncrEXFloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func appendIncrEXTail(args []interface{}, exp *ExpirationOption, enx bool) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

type SetCondition string

const (
	NX SetCondition = "NX"

	XX SetCondition = "XX"
)

type ExpirationMode string

const (
	EX ExpirationMode = "EX"

	PX ExpirationMode = "PX"

	EXAT ExpirationMode = "EXAT"

	PXAT ExpirationMode = "PXAT"

	KEEPTTL ExpirationMode = "KEEPTTL"

	PERSIST ExpirationMode = "PERSIST"
)

type ExpirationOption struct {
	Mode  ExpirationMode
	Value int64
}

func (c cmdable) LCS(ctx context.Context, q *LCSQuery) *LCSCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) MGet(ctx context.Context, keys ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) MSet(ctx context.Context, values ...interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) MSetNX(ctx context.Context, values ...interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

type MSetEXArgs struct {
	Condition  SetCondition
	Expiration *ExpirationOption
}

func (c cmdable) MSetEX(ctx context.Context, args MSetEXArgs, values ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

type SetArgs struct {
	Mode string

	MatchValue interface{}

	MatchDigest uint64

	TTL      time.Duration
	ExpireAt time.Time

	Get bool

	KeepTTL bool
}

func (c cmdable) SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetFromBuffer(ctx context.Context, key string, buf []byte) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetIFEQ(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetIFEQGet(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetIFNE(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetIFNEGet(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetIFDEQ(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetIFDEQGet(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetIFDNE(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetIFDNEGet(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) StrLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}
