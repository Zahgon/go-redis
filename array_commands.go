package redis

import (
	"context"
)

type ArrayCmdable interface {
	ARSet(ctx context.Context, key string, index uint64, values ...string) *IntCmd
	ARGet(ctx context.Context, key string, index uint64) *StringCmd
	ARGetRange(ctx context.Context, key string, start, end uint64) *SliceCmd
	ARMGet(ctx context.Context, key string, indexes ...uint64) *SliceCmd
	ARMSet(ctx context.Context, key string, members ...AREntry) *IntCmd
	ARInsert(ctx context.Context, key string, values ...string) *UintCmd
	ARDel(ctx context.Context, key string, indexes ...uint64) *IntCmd
	ARDelRange(ctx context.Context, key string, ranges ...ARRange) *UintCmd
	ARLen(ctx context.Context, key string) *UintCmd
	ARCount(ctx context.Context, key string) *UintCmd
	ARNext(ctx context.Context, key string) *UintCmd
	ARSeek(ctx context.Context, key string, index uint64) *IntCmd
	ARInfo(ctx context.Context, key string) *MapStringInterfaceCmd
	ARInfoFull(ctx context.Context, key string) *MapStringInterfaceCmd
	ARScan(ctx context.Context, key string, start, end uint64, args *ARScanArgs) *AREntrySliceCmd
	AROpSum(ctx context.Context, key string, start, end uint64) *StringCmd
	AROpMin(ctx context.Context, key string, start, end uint64) *StringCmd
	AROpMax(ctx context.Context, key string, start, end uint64) *StringCmd
	AROpAnd(ctx context.Context, key string, start, end uint64) *IntCmd
	AROpOr(ctx context.Context, key string, start, end uint64) *IntCmd
	AROpXor(ctx context.Context, key string, start, end uint64) *IntCmd
	AROpMatch(ctx context.Context, key string, start, end uint64, value string) *IntCmd
	AROpUsed(ctx context.Context, key string, start, end uint64) *IntCmd
	ARGrep(ctx context.Context, key string, start, end string, args *ARGrepArgs) *UintSliceCmd
	ARGrepWithValues(ctx context.Context, key string, start, end string, args *ARGrepArgs) *AREntrySliceCmd
	ARRing(ctx context.Context, key string, size uint64, values ...string) *UintCmd
	ARLastItems(ctx context.Context, key string, count uint64, rev bool) *SliceCmd
}

type AREntry struct {
	Index uint64
	Value string
}

type ARRange struct {
	Start uint64
	End   uint64
}

type ARScanArgs struct {
	Limit uint64
}

type ARGrepPredicateType string

const (
	ARGrepExact ARGrepPredicateType = "EXACT"
	ARGrepMatch ARGrepPredicateType = "MATCH"
	ARGrepGlob  ARGrepPredicateType = "GLOB"
	ARGrepRegex ARGrepPredicateType = "RE"
)

type ARGrepPredicate struct {
	Type  ARGrepPredicateType
	Value string
}

type ARGrepArgs struct {
	Predicates []ARGrepPredicate
	CombineAnd bool
	Limit      uint64
	NoCase     bool
}

func (c cmdable) ARSet(ctx context.Context, key string, index uint64, values ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARGet(ctx context.Context, key string, index uint64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARGetRange(ctx context.Context, key string, start, end uint64) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARMGet(ctx context.Context, key string, indexes ...uint64) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARMSet(ctx context.Context, key string, members ...AREntry) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARInsert(ctx context.Context, key string, values ...string) *UintCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARDel(ctx context.Context, key string, indexes ...uint64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARDelRange(ctx context.Context, key string, ranges ...ARRange) *UintCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARLen(ctx context.Context, key string) *UintCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARCount(ctx context.Context, key string) *UintCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARNext(ctx context.Context, key string) *UintCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARSeek(ctx context.Context, key string, index uint64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARInfo(ctx context.Context, key string) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARInfoFull(ctx context.Context, key string) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARScan(ctx context.Context, key string, start, end uint64, scanArgs *ARScanArgs) *AREntrySliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) AROpSum(ctx context.Context, key string, start, end uint64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) AROpMin(ctx context.Context, key string, start, end uint64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) AROpMax(ctx context.Context, key string, start, end uint64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) AROpAnd(ctx context.Context, key string, start, end uint64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) AROpOr(ctx context.Context, key string, start, end uint64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) AROpXor(ctx context.Context, key string, start, end uint64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) AROpMatch(ctx context.Context, key string, start, end uint64, value string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) AROpUsed(ctx context.Context, key string, start, end uint64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARGrep(ctx context.Context, key string, start, end string, grepArgs *ARGrepArgs) *UintSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARGrepWithValues(ctx context.Context, key string, start, end string, grepArgs *ARGrepArgs) *AREntrySliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (args *ARGrepArgs) Len() int { _ = "STUB: not implemented"; return 0 }

func (args *ARGrepArgs) Append(a []any) []any { _ = "STUB: not implemented"; return nil }

func (c cmdable) ARRing(ctx context.Context, key string, size uint64, values ...string) *UintCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ARLastItems(ctx context.Context, key string, count uint64, rev bool) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}
