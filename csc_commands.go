package redis

var defaultCacheableCommands = map[string]struct{}{

	"get": {}, "mget": {}, "getbit": {}, "getrange": {},
	"strlen": {}, "substr": {},

	"hget": {}, "hgetall": {}, "hmget": {},
	"hkeys": {}, "hvals": {}, "hlen": {},
	"hexists": {}, "hstrlen": {},

	"lindex": {}, "llen": {}, "lpos": {}, "lrange": {},

	"scard": {}, "sismember": {}, "smembers": {}, "smismember": {},
	"sdiff": {}, "sinter": {}, "sintercard": {}, "sunion": {},

	"zcard": {}, "zcount": {}, "zlexcount": {}, "zmscore": {},
	"zrange": {}, "zrangebylex": {}, "zrangebyscore": {},
	"zrank": {}, "zrevrange": {}, "zrevrangebylex": {},
	"zrevrangebyscore": {}, "zrevrank": {}, "zscore": {},
	"zdiff": {}, "zinter": {}, "zunion": {},

	"bitcount": {}, "bitfield_ro": {}, "bitpos": {},

	"exists": {}, "type": {}, "sort_ro": {}, "lcs": {},

	"geodist": {}, "geohash": {}, "geopos": {}, "geosearch": {},
	"georadiusbymember_ro": {}, "georadius_ro": {},

	"xlen": {}, "xrange": {}, "xrevrange": {},

	"json.get": {}, "json.mget": {}, "json.arrindex": {}, "json.arrlen": {},
	"json.objkeys": {}, "json.objlen": {}, "json.resp": {},
	"json.strlen": {}, "json.type": {},

	"ts.get": {}, "ts.info": {}, "ts.range": {}, "ts.revrange": {},
}

func isCacheable(cmd Cmder) bool { _ = "STUB: not implemented"; return false }

func sortROHasByGet(cmd Cmder) bool { _ = "STUB: not implemented"; return false }

func isClientTrackingCmd(cmd Cmder) bool { _ = "STUB: not implemented"; return false }

func isSelectCmd(cmd Cmder) bool { _ = "STUB: not implemented"; return false }

func isAuthCmd(cmd Cmder) bool { _ = "STUB: not implemented"; return false }

func isProtocolChangingHelloCmd(cmd Cmder) bool { _ = "STUB: not implemented"; return false }

func isResetCmd(cmd Cmder) bool { _ = "STUB: not implemented"; return false }

func isSubscribeCmd(cmd Cmder) bool { _ = "STUB: not implemented"; return false }

func buildCacheKey(cmd Cmder) (string, bool) { _ = "STUB: not implemented"; return "", false }

func keyArg(cmd Cmder, pos int) (string, bool) { _ = "STUB: not implemented"; return "", false }

func extractRedisKeys(cmd Cmder) []string { _ = "STUB: not implemented"; return nil }
