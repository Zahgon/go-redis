package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	_ = rdb.Set(ctx, "key_with_ttl", "bar", time.Minute).Err()
	_ = rdb.Set(ctx, "key_without_ttl_1", "", 0).Err()
	_ = rdb.Set(ctx, "key_without_ttl_2", "", 0).Err()

	checker := NewKeyChecker(rdb, 100)

	start := time.Now()
	checker.Start(ctx)

	iter := rdb.Scan(ctx, 0, "", 0).Iterator()
	for iter.Next(ctx) {
		checker.Add(iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	deleted := checker.Stop()
	fmt.Println("deleted", deleted, "keys", "in", time.Since(start))
}

type KeyChecker struct {
	rdb       *redis.Client
	batchSize int
	ch        chan string
	delCh     chan string
	wg        sync.WaitGroup
	deleted   int
	logger    *zap.Logger
}

func NewKeyChecker(rdb *redis.Client, batchSize int) *KeyChecker {
	_ = "STUB: not implemented"
	return nil
}

func (c *KeyChecker) Add(key string) { _ = "STUB: not implemented"; return }

func (c *KeyChecker) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *KeyChecker) Stop() int { _ = "STUB: not implemented"; return 0 }

func (c *KeyChecker) checkKeys(ctx context.Context, keys []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *KeyChecker) del(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
