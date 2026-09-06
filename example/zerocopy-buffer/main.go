package main

import (
	"context"
	"crypto/rand"
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

	defer rdb.Del(ctx, "zerocopy:greeting", "zerocopy:largeblob")

	fmt.Println("=== Example 1: Basic round-trip ===")

	writeData := []byte("Hello, zero-copy world!")
	if err := rdb.SetFromBuffer(ctx, "zerocopy:greeting", writeData).Err(); err != nil {
		log.Fatalf("SetFromBuffer failed: %v", err)
	}
	fmt.Printf("SET zerocopy:greeting (%d bytes)\n", len(writeData))

	readBuf := make([]byte, 100)
	cmd := rdb.GetToBuffer(ctx, "zerocopy:greeting", readBuf)
	if err := cmd.Err(); err != nil {
		log.Fatalf("GetToBuffer failed: %v", err)
	}

	n := cmd.Val()
	fmt.Printf("GET zerocopy:greeting -> %d bytes: %q\n", n, string(cmd.Bytes()))

	fmt.Println("\n=== Example 2: Large binary data (1 MB) ===")

	const blobSize = 1 * 1024 * 1024
	blob := make([]byte, blobSize)
	if _, err := rand.Read(blob); err != nil {
		log.Fatalf("failed to generate random data: %v", err)
	}

	if err := rdb.SetFromBuffer(ctx, "zerocopy:largeblob", blob).Err(); err != nil {
		log.Fatalf("SetFromBuffer (1MB) failed: %v", err)
	}
	fmt.Printf("SET zerocopy:largeblob (%d bytes)\n", blobSize)

	largeBuf := make([]byte, blobSize+64)
	cmd = rdb.GetToBuffer(ctx, "zerocopy:largeblob", largeBuf)
	if err := cmd.Err(); err != nil {
		log.Fatalf("GetToBuffer (1MB) failed: %v", err)
	}
	fmt.Printf("GET zerocopy:largeblob -> %d bytes\n", cmd.Val())

	retrieved := cmd.Bytes()
	if len(retrieved) != blobSize {
		log.Fatalf("size mismatch: expected %d, got %d", blobSize, len(retrieved))
	}
	for i := range blob {
		if blob[i] != retrieved[i] {
			log.Fatalf("data mismatch at byte %d", i)
		}
	}
	fmt.Println("Data integrity verified ✓")

	fmt.Println("\n=== Example 3: Handling non-existent keys ===")

	buf := make([]byte, 64)
	cmd = rdb.GetToBuffer(ctx, "zerocopy:nonexistent", buf)
	if cmd.Err() == redis.Nil {
		fmt.Println("Key does not exist (redis.Nil) — as expected")
	} else if cmd.Err() != nil {
		log.Fatalf("unexpected error: %v", cmd.Err())
	}

	runPackingExamples(ctx, rdb)
	runPipelineExamples(ctx, rdb)

	fmt.Println("\nDone!")
}
