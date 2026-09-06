//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	minDataLength     = 4
	redisAddr         = ":6379"
	dialTimeout       = 10 * time.Second
	readTimeout       = 10 * time.Second
	writeTimeout      = 10 * time.Second
	poolSize          = 10
	poolTimeout       = 10 * time.Second
	scanCount         = 10
	maxIterPercentage = 256
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

type redisOperation func(key, value string)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		DialTimeout:  dialTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		PoolSize:     poolSize,
		PoolTimeout:  poolTimeout,
	})
}

func Fuzz(data []byte) int { _ = "STUB: not implemented"; return 0 }
