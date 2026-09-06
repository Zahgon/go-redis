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

	if err := rdb.ConfigSet(ctx, "search-enable-unstable-features", "yes").Err(); err != nil {
		log.Fatalf("cannot enable search unstable features (COLLECT needs Redis 8.8+ with search): %v", err)
	}

	if err := rdb.FlushDB(ctx).Err(); err != nil {
		log.Fatalf("flushdb: %v", err)
	}
	createIndex(ctx, rdb)
	seedProducts(ctx, rdb)

	fmt.Println("# 1. Top-2 products per category (options struct)")
	topPerCategory(ctx, rdb)

	fmt.Println("\n# 2. Whole documents per brand (builder, LOAD * + FIELDS *)")
	documentsPerBrand(ctx, rdb)
}

func createIndex(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func seedProducts(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func topPerCategory(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func documentsPerBrand(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func printEntries(row redis.AggregateRow, alias string) { _ = "STUB: not implemented"; return }
