package redis

import (
	"context"
	"time"
)

type GenericCmdable interface {
	Del(ctx context.Context, keys ...string) *IntCmd
	Dump(ctx context.Context, key string) *StringCmd
	Exists(ctx context.Context, keys ...string) *IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
	ExpireTime(ctx context.Context, key string) *DurationCmd
	ExpireNX(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireXX(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireGT(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireLT(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	Keys(ctx context.Context, pattern string) *StringSliceCmd
	Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *StatusCmd
	Move(ctx context.Context, key string, db int) *BoolCmd
	ObjectFreq(ctx context.Context, key string) *IntCmd
	ObjectRefCount(ctx context.Context, key string) *IntCmd
	ObjectEncoding(ctx context.Context, key string) *StringCmd
	ObjectIdleTime(ctx context.Context, key string) *DurationCmd
	Persist(ctx context.Context, key string) *BoolCmd
	PExpire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	PExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
	PExpireTime(ctx context.Context, key string) *DurationCmd
	PTTL(ctx context.Context, key string) *DurationCmd
	RandomKey(ctx context.Context) *StringCmd
	Rename(ctx context.Context, key, newkey string) *StatusCmd
	RenameNX(ctx context.Context, key, newkey string) *BoolCmd
	Restore(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd
	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd
	Sort(ctx context.Context, key string, sort *Sort) *StringSliceCmd
	SortRO(ctx context.Context, key string, sort *Sort) *StringSliceCmd
	SortStore(ctx context.Context, key, store string, sort *Sort) *IntCmd
	SortInterfaces(ctx context.Context, key string, sort *Sort) *SliceCmd
	Touch(ctx context.Context, keys ...string) *IntCmd
	TTL(ctx context.Context, key string) *DurationCmd
	Type(ctx context.Context, key string) *StatusCmd
	Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) *IntCmd

	Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd
	ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd
}

func (c cmdable) Del(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Unlink(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Dump(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Exists(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ExpireNX(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ExpireXX(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ExpireGT(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ExpireLT(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) expire(
	ctx context.Context, key string, expiration time.Duration, mode string,
) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ExpireTime(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Keys(ctx context.Context, pattern string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Move(ctx context.Context, key string, db int) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ObjectFreq(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ObjectRefCount(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ObjectEncoding(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ObjectIdleTime(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Persist(ctx context.Context, key string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PExpire(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PExpireTime(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PTTL(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) RandomKey(ctx context.Context) *StringCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) Rename(ctx context.Context, key, newkey string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) RenameNX(ctx context.Context, key, newkey string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Restore(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

type Sort struct {
	By            string
	Offset, Count int64
	Get           []string
	Order         string
	Alpha         bool
}

func (sort *Sort) args(command, key string) []interface{} { _ = "STUB: not implemented"; return nil }

func (c cmdable) SortRO(ctx context.Context, key string, sort *Sort) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Sort(ctx context.Context, key string, sort *Sort) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SortStore(ctx context.Context, key, store string, sort *Sort) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SortInterfaces(ctx context.Context, key string, sort *Sort) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Touch(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) TTL(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Type(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Copy(ctx context.Context, sourceKey, destKey string, db int, replace bool) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}
