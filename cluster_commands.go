package redis

import "context"

type ClusterCmdable interface {
	ClusterMyShardID(ctx context.Context) *StringCmd
	ClusterMyID(ctx context.Context) *StringCmd
	ClusterSlots(ctx context.Context) *ClusterSlotsCmd
	ClusterShards(ctx context.Context) *ClusterShardsCmd
	ClusterLinks(ctx context.Context) *ClusterLinksCmd
	ClusterNodes(ctx context.Context) *StringCmd
	ClusterMeet(ctx context.Context, host, port string) *StatusCmd
	ClusterForget(ctx context.Context, nodeID string) *StatusCmd
	ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd
	ClusterResetSoft(ctx context.Context) *StatusCmd
	ClusterResetHard(ctx context.Context) *StatusCmd
	ClusterInfo(ctx context.Context) *StringCmd
	ClusterKeySlot(ctx context.Context, key string) *IntCmd
	ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *StringSliceCmd
	ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd
	ClusterCountKeysInSlot(ctx context.Context, slot int) *IntCmd
	ClusterDelSlots(ctx context.Context, slots ...int) *StatusCmd
	ClusterDelSlotsRange(ctx context.Context, min, max int) *StatusCmd
	ClusterSaveConfig(ctx context.Context) *StatusCmd
	ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd
	ClusterFailover(ctx context.Context) *StatusCmd
	ClusterAddSlots(ctx context.Context, slots ...int) *StatusCmd
	ClusterAddSlotsRange(ctx context.Context, min, max int) *StatusCmd
	ReadOnly(ctx context.Context) *StatusCmd
	ReadWrite(ctx context.Context) *StatusCmd
}

func (c cmdable) ClusterMyShardID(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterMyID(ctx context.Context) *StringCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ClusterSlots(ctx context.Context) *ClusterSlotsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterShards(ctx context.Context) *ClusterShardsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterLinks(ctx context.Context) *ClusterLinksCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterNodes(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterMeet(ctx context.Context, host, port string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterForget(ctx context.Context, nodeID string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterResetSoft(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterResetHard(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterInfo(ctx context.Context) *StringCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ClusterKeySlot(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterCountKeysInSlot(ctx context.Context, slot int) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterDelSlots(ctx context.Context, slots ...int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterDelSlotsRange(ctx context.Context, min, max int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterSaveConfig(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterFailover(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterAddSlots(ctx context.Context, slots ...int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ClusterAddSlotsRange(ctx context.Context, min, max int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ReadOnly(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ReadWrite(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }
