package redis

import (
	"context"
	"time"
)

type HashCmdable interface {
	HDel(ctx context.Context, key string, fields ...string) *IntCmd
	HExists(ctx context.Context, key, field string) *BoolCmd
	HGet(ctx context.Context, key, field string) *StringCmd
	HGetAll(ctx context.Context, key string) *MapStringStringCmd
	HGetDel(ctx context.Context, key string, fields ...string) *StringSliceCmd
	HGetEX(ctx context.Context, key string, fields ...string) *StringSliceCmd
	HGetEXWithArgs(ctx context.Context, key string, options *HGetEXOptions, fields ...string) *StringSliceCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd
	HKeys(ctx context.Context, key string) *StringSliceCmd
	HLen(ctx context.Context, key string) *IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *SliceCmd
	HSet(ctx context.Context, key string, values ...interface{}) *IntCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *BoolCmd
	HSetEX(ctx context.Context, key string, fieldsAndValues ...string) *IntCmd
	HSetEXWithArgs(ctx context.Context, key string, options *HSetEXOptions, fieldsAndValues ...string) *IntCmd
	HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	HScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	HVals(ctx context.Context, key string) *StringSliceCmd
	HRandField(ctx context.Context, key string, count int) *StringSliceCmd
	HRandFieldWithValues(ctx context.Context, key string, count int) *KeyValueSliceCmd
	HStrLen(ctx context.Context, key, field string) *IntCmd
	HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *IntSliceCmd
	HExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd
	HPExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *IntSliceCmd
	HPExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd
	HExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *IntSliceCmd
	HExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd
	HPExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *IntSliceCmd
	HPExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd
	HPersist(ctx context.Context, key string, fields ...string) *IntSliceCmd
	HExpireTime(ctx context.Context, key string, fields ...string) *IntSliceCmd
	HPExpireTime(ctx context.Context, key string, fields ...string) *IntSliceCmd
	HTTL(ctx context.Context, key string, fields ...string) *IntSliceCmd
	HPTTL(ctx context.Context, key string, fields ...string) *IntSliceCmd
}

func (c cmdable) HDel(ctx context.Context, key string, fields ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HExists(ctx context.Context, key, field string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HGet(ctx context.Context, key, field string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HGetAll(ctx context.Context, key string) *MapStringStringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HKeys(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HMGet(ctx context.Context, key string, fields ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HSet(ctx context.Context, key string, values ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HMSet(ctx context.Context, key string, values ...interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HVals(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HRandField(ctx context.Context, key string, count int) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HRandFieldWithValues(ctx context.Context, key string, count int) *KeyValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HStrLen(ctx context.Context, key, field string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

type HExpireArgs struct {
	NX bool
	XX bool
	GT bool
	LT bool
}

func (c cmdable) HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HPExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HPExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HPExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HPExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HPersist(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HExpireTime(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HPExpireTime(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HTTL(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HPTTL(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HGetDel(ctx context.Context, key string, fields ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HGetEX(ctx context.Context, key string, fields ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type HGetEXExpirationType string

const (
	HGetEXExpirationEX      HGetEXExpirationType = "EX"
	HGetEXExpirationPX      HGetEXExpirationType = "PX"
	HGetEXExpirationEXAT    HGetEXExpirationType = "EXAT"
	HGetEXExpirationPXAT    HGetEXExpirationType = "PXAT"
	HGetEXExpirationPERSIST HGetEXExpirationType = "PERSIST"
)

type HGetEXOptions struct {
	ExpirationType HGetEXExpirationType
	ExpirationVal  int64
}

func (c cmdable) HGetEXWithArgs(ctx context.Context, key string, options *HGetEXOptions, fields ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type HSetEXCondition string

const (
	HSetEXFNX HSetEXCondition = "FNX"
	HSetEXFXX HSetEXCondition = "FXX"
)

type HSetEXExpirationType string

const (
	HSetEXExpirationEX      HSetEXExpirationType = "EX"
	HSetEXExpirationPX      HSetEXExpirationType = "PX"
	HSetEXExpirationEXAT    HSetEXExpirationType = "EXAT"
	HSetEXExpirationPXAT    HSetEXExpirationType = "PXAT"
	HSetEXExpirationKEEPTTL HSetEXExpirationType = "KEEPTTL"
)

type HSetEXOptions struct {
	Condition      HSetEXCondition
	ExpirationType HSetEXExpirationType
	ExpirationVal  int64
}

func (c cmdable) HSetEX(ctx context.Context, key string, fieldsAndValues ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HSetEXWithArgs(ctx context.Context, key string, options *HSetEXOptions, fieldsAndValues ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}
