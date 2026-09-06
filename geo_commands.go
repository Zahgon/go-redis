package redis

import (
	"context"
)

type GeoCmdable interface {
	GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) *IntCmd
	GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *IntCmd
	GeoRadiusByMember(ctx context.Context, key, member string, query *GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *GeoRadiusQuery) *IntCmd
	GeoSearch(ctx context.Context, key string, q *GeoSearchQuery) *StringSliceCmd
	GeoSearchLocation(ctx context.Context, key string, q *GeoSearchLocationQuery) *GeoSearchLocationCmd
	GeoSearchStore(ctx context.Context, key, store string, q *GeoSearchStoreQuery) *IntCmd
	GeoDist(ctx context.Context, key string, member1, member2, unit string) *FloatCmd
	GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd
}

func (c cmdable) GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GeoRadius(
	ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery,
) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GeoRadiusStore(
	ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery,
) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GeoRadiusByMember(
	ctx context.Context, key, member string, query *GeoRadiusQuery,
) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GeoRadiusByMemberStore(
	ctx context.Context, key, member string, query *GeoRadiusQuery,
) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GeoSearch(ctx context.Context, key string, q *GeoSearchQuery) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GeoSearchLocation(
	ctx context.Context, key string, q *GeoSearchLocationQuery,
) *GeoSearchLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GeoSearchStore(ctx context.Context, key, store string, q *GeoSearchStoreQuery) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GeoDist(
	ctx context.Context, key string, member1, member2, unit string,
) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd {
	_ = "STUB: not implemented"
	return nil
}
