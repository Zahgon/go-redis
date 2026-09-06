package redis

import (
	"context"
)

func (c *ClusterClient) DBSize(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c *ClusterClient) ScriptLoad(ctx context.Context, script string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) ScriptFlush(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}
