package main

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	duration   = 3 * time.Second
	goroutines = 500
	window     = 200
)

func addr() string { _ = "STUB: not implemented"; return "" }

func fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

type result struct {
	name    string
	ordered string
	ops     float64
}

func main() {
	ctx := context.Background()

	probe := redis.NewClient(&redis.Options{Addr: addr()})
	if err := probe.Ping(ctx).Err(); err != nil {
		fatalf("cannot reach redis at %s: %v", addr(), err)
	}
	probe.Close()

	usageTour(ctx)
	throughputComparison(ctx)
}

func usageTour(ctx context.Context) { _ = "STUB: not implemented"; return }

func clusterTour(ctx context.Context, addrs []string) { _ = "STUB: not implemented"; return }

func throughputComparison(ctx context.Context) { _ = "STUB: not implemented"; return }

func benchNormalBlocking(ctx context.Context) float64 { _ = "STUB: not implemented"; return 0 }

func benchOrderedBlocking(ctx context.Context) float64 { _ = "STUB: not implemented"; return 0 }

func benchOrderedReadLater(ctx context.Context) float64 { _ = "STUB: not implemented"; return 0 }

func benchUnorderedReadLater(ctx context.Context) float64 { _ = "STUB: not implemented"; return 0 }

func benchReadLater(ctx context.Context, ap *redis.AutoPipeliner) float64 {
	_ = "STUB: not implemented"
	return 0
}

func run(op func(g int, doOp func())) float64 { _ = "STUB: not implemented"; return 0 }
