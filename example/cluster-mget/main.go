package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"localhost:16600",
			"localhost:16601",
			"localhost:16602",
			"localhost:16603",
			"localhost:16604",
			"localhost:16605",
		},
	})
	defer rdb.Close()

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis cluster: %v", err))
	}

	fmt.Println("✓ Connected to Redis cluster")

	keys := make([]string, 10)
	values := make([]string, 10)
	for i := 0; i < 10; i++ {
		keys[i] = fmt.Sprintf("key%d", i)
		values[i] = fmt.Sprintf("value%d", i)
	}

	fmt.Println("\n=== Setting 10 keys ===")
	for i := 0; i < 10; i++ {
		err := rdb.Set(ctx, keys[i], values[i], 0).Err()
		if err != nil {
			panic(fmt.Sprintf("Failed to set %s: %v", keys[i], err))
		}
		fmt.Printf("✓ SET %s = %s\n", keys[i], values[i])
	}

	fmt.Println("\n=== Cleaning up ===")
	for _, key := range keys {
		if err := rdb.Del(ctx, key).Err(); err != nil {
			fmt.Printf("Warning: Failed to delete %s: %v\n", key, err)
		}
	}
	fmt.Println("✓ Cleanup complete")

	err := rdb.Set(ctx, "{tag}exists", "asdf", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "{tag}nilkeykey1").Result()
	fmt.Printf("\nval: %+v err: %+v\n", val, err)
	valm, err := rdb.MGet(ctx, "{tag}nilkeykey1", "{tag}exists").Result()
	fmt.Printf("\nval: %+v err: %+v\n", valm, err)
}
