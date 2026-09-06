package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

const indexName = "idx:cities"

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	setup(ctx, rdb)
	defer cleanup(ctx, rdb)

	aliases, err := rdb.FTAliasList(ctx, indexName).Result()
	if err != nil {
		log.Fatalf("FT.ALIASLIST: %v", err)
	}
	fmt.Printf("aliases right after FT.CREATE: %d\n", len(aliases))

	for _, alias := range []string{"cities", "cities-latest"} {
		if err := rdb.FTAliasAdd(ctx, indexName, alias).Err(); err != nil {
			log.Fatalf("FT.ALIASADD %s: %v", alias, err)
		}
	}

	aliases, err = rdb.FTAliasList(ctx, indexName).Result()
	if err != nil {
		log.Fatalf("FT.ALIASLIST: %v", err)
	}
	fmt.Printf("aliases after FT.ALIASADD: %v\n", aliases)

	if err := rdb.FTAliasList(ctx, "cities").Err(); err != nil {
		fmt.Printf("FT.ALIASLIST with an alias as argument: %v\n", err)
	}
}

func setup(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }

func cleanup(ctx context.Context, rdb *redis.Client) { _ = "STUB: not implemented"; return }
