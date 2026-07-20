package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

const indexName = "idx:products"

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Protocol: 2,
	})
	defer rdb.Close()

	if err := rdb.FlushDB(ctx).Err(); err != nil {
		log.Fatalf("flushdb: %v", err)
	}
	createIndex(ctx, rdb)
	seedProducts(ctx, rdb)

	fmt.Println("# 1. Trim before GROUPBY (real-world use case)")
	trimBeforeGroup(ctx, rdb)

	fmt.Println("\n# 2. Multi-stage pipeline (LOAD -> APPLY -> SORTBY -> GROUPBY -> APPLY -> SORTBY)")
	multiStagePipeline(ctx, rdb)

	fmt.Println("\n# 3. Same as #1 but using AggregateBuilder")
	withBuilder(ctx, rdb)
}

func createIndex(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func seedProducts(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func trimBeforeGroup(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func multiStagePipeline(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func withBuilder(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func printRows(res *redis.FTAggregateResult, err error) { _ = "STUB: not implemented"; return }
