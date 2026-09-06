package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	_ = rdb.FlushDB(ctx).Err()

	fmt.Printf("# BLOOM\n")
	bloomFilter(ctx, rdb)

	fmt.Printf("\n# CUCKOO\n")
	cuckooFilter(ctx, rdb)

	fmt.Printf("\n# COUNT-MIN\n")
	countMinSketch(ctx, rdb)

	fmt.Printf("\n# TOP-K\n")
	topK(ctx, rdb)
}

func bloomFilter(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func cuckooFilter(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func countMinSketch(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func topK(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }
