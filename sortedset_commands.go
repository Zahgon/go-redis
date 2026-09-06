package redis

import (
	"context"
	"time"
)

type SortedSetCmdable interface {
	BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
	BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
	BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *ZSliceWithKeyCmd
	ZAdd(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddLT(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddGT(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddNX(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddXX(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd
	ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd
	ZCard(ctx context.Context, key string) *IntCmd
	ZCount(ctx context.Context, key, min, max string) *IntCmd
	ZLexCount(ctx context.Context, key, min, max string) *IntCmd
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd
	ZInter(ctx context.Context, store *ZStore) *StringSliceCmd
	ZInterWithScores(ctx context.Context, store *ZStore) *ZSliceCmd
	ZInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd
	ZInterStore(ctx context.Context, destination string, store *ZStore) *IntCmd
	ZMPop(ctx context.Context, order string, count int64, keys ...string) *ZSliceWithKeyCmd
	ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd
	ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd
	ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd
	ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
	ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
	ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd
	ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd
	ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd
	ZRank(ctx context.Context, key, member string) *IntCmd
	ZRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd
	ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd
	ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd
	ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd
	ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
	ZRevRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
	ZRevRank(ctx context.Context, key, member string) *IntCmd
	ZRevRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd
	ZScore(ctx context.Context, key, member string) *FloatCmd
	ZUnionStore(ctx context.Context, dest string, store *ZStore) *IntCmd
	ZRandMember(ctx context.Context, key string, count int) *StringSliceCmd
	ZRandMemberWithScores(ctx context.Context, key string, count int) *ZSliceCmd
	ZUnion(ctx context.Context, store ZStore) *StringSliceCmd
	ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd
	ZDiff(ctx context.Context, keys ...string) *StringSliceCmd
	ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd
	ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
}

func (c cmdable) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *ZSliceWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

type ZAddArgs struct {
	NX      bool
	XX      bool
	LT      bool
	GT      bool
	Ch      bool
	Members []Z
}

func (c cmdable) zAddArgs(key string, args ZAddArgs, incr bool) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZAdd(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZAddLT(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZAddGT(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZAddNX(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZAddXX(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZCount(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZLexCount(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZInterStore(ctx context.Context, destination string, store *ZStore) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZInter(ctx context.Context, store *ZStore) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZInterWithScores(ctx context.Context, store *ZStore) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZMPop(ctx context.Context, order string, count int64, keys ...string) *ZSliceWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type ZRangeArgs struct {
	Key string

	Start interface{}
	Stop  interface{}

	ByScore bool
	ByLex   bool

	Rev bool

	Offset int64
	Count  int64
}

func (z ZRangeArgs) appendArgs(args []interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

func (c cmdable) zRangeBy(ctx context.Context, zcmd, key string, opt *ZRangeBy, withScores bool) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRank(ctx context.Context, key, member string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) zRevRangeBy(ctx context.Context, zcmd, key string, opt *ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRevRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRevRank(ctx context.Context, key, member string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRevRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZScore(ctx context.Context, key, member string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZUnion(ctx context.Context, store ZStore) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZUnionStore(ctx context.Context, dest string, store *ZStore) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRandMember(ctx context.Context, key string, count int) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZRandMemberWithScores(ctx context.Context, key string, count int) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

type Z struct {
	Score  float64
	Member interface{}
}

type ZWithKey struct {
	Z
	Key string
}

type ZStore struct {
	Keys    []string
	Weights []float64

	Aggregate string
}

func (z ZStore) len() (n int) { _ = "STUB: not implemented"; return 0 }

func (z ZStore) appendArgs(args []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }
