package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	keys := seed(ctx, rdb)
	defer rdb.Del(ctx, keys...)

	labels, err := rdb.TSQueryLabels(ctx, []string{"type=sensor"}).Result()
	if err != nil {
		log.Fatalf("TS.QUERYLABELS LABELS: %v", err)
	}
	fmt.Printf("label names for type=sensor: %v\n", labels)

	locations, err := rdb.TSQueryLabelValues(ctx, "location", []string{"type=sensor"}).Result()
	if err != nil {
		log.Fatalf("TS.QUERYLABELS VALUES: %v", err)
	}
	fmt.Printf("locations for type=sensor:   %v\n", locations)

	series, err := rdb.TSQueryIndex(ctx, []string{"type=sensor", "location=kitchen"}).Result()
	if err != nil {
		log.Fatalf("TS.QUERYINDEX: %v", err)
	}
	fmt.Printf("kitchen sensor series:       %v\n", series)

	missing, err := rdb.TSQueryLabelValues(ctx, "rack", []string{"type=sensor"}).Result()
	if err != nil {
		log.Fatalf("TS.QUERYLABELS VALUES rack: %v", err)
	}
	fmt.Printf("values of an absent label:   %d\n", len(missing))

	all, err := rdb.TSQueryLabels(ctx, nil).Result()
	if err != nil {
		log.Fatalf("TS.QUERYLABELS LABELS (no filter): %v", err)
	}
	fmt.Printf("label names, all series:     %v\n", all)
}

func seed(ctx context.Context, rdb *redis.Client) []string { _ = "STUB: not implemented"; return nil }
