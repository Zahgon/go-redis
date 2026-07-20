package redis

import (
	"context"
)

type SetCmdable interface {
	SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd
	SCard(ctx context.Context, key string) *IntCmd
	SDiff(ctx context.Context, keys ...string) *StringSliceCmd
	SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd
	SInter(ctx context.Context, keys ...string) *StringSliceCmd
	SInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd
	SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd
	SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd
	SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd
	SMembers(ctx context.Context, key string) *StringSliceCmd
	SMembersMap(ctx context.Context, key string) *StringStructMapCmd
	SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd
	SPop(ctx context.Context, key string) *StringCmd
	SPopN(ctx context.Context, key string, count int64) *StringSliceCmd
	SRandMember(ctx context.Context, key string) *StringCmd
	SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd
	SRem(ctx context.Context, key string, members ...interface{}) *IntCmd
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	SUnion(ctx context.Context, keys ...string) *StringSliceCmd
	SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd
}

func (c cmdable) SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SInter(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SMembers(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SMembersMap(ctx context.Context, key string) *StringStructMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SPop(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SPopN(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SRandMember(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SRem(ctx context.Context, key string, members ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SUnion(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}
