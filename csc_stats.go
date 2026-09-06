package redis

type CSCStats struct {
	Hits             uint64
	Misses           uint64
	Entries          int
	MemoryUsageBytes int64
}

type cacheStatsReporter interface {
	Stats() CSCStats
}

func (c *Client) CSCStats() CSCStats { _ = "STUB: not implemented"; return *new(CSCStats) }
