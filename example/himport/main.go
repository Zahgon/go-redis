package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

const (
	fieldsetName = "user-fields"
	bulkKeys     = 10_000
	batchSize    = 500
)

func main() {
	ctx := context.Background()

	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		PoolSize: 8,
	})
	defer rdb.Close()

	if err := rdb.HImportDiscardAll(ctx).Err(); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "unknown command") {
			log.Fatalf("this server does not support HIMPORT (requires Redis 8.10+): %v", err)
		}
		log.Fatalf("connection failed: %v", err)
	}

	basicUsage(ctx, rdb)
	bulkIngestion(ctx, rdb)
	concurrentWriters(ctx, rdb)
	inspectEncoding(ctx, rdb)
	discard(ctx, rdb)
}

func basicUsage(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func bulkIngestion(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func concurrentWriters(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func inspectEncoding(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func discard(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }
