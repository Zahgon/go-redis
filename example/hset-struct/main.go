package main

import (
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/redis/go-redis/v9"
)

type Model struct {
	Str1    string     `redis:"str1"`
	Str2    string     `redis:"str2"`
	Str3    *string    `redis:"str3"`
	Str4    *string    `redis:"str4"`
	Bytes   []byte     `redis:"bytes"`
	Int     int        `redis:"int"`
	Int2    *int       `redis:"int2"`
	Int3    *int       `redis:"int3"`
	Bool    bool       `redis:"bool"`
	Bool2   *bool      `redis:"bool2"`
	Bool3   *bool      `redis:"bool3"`
	Bool4   *bool      `redis:"bool4,omitempty"`
	Time    time.Time  `redis:"time"`
	Time2   *time.Time `redis:"time2"`
	Time3   *time.Time `redis:"time3"`
	Ignored struct{}   `redis:"-"`
}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	_ = rdb.FlushDB(ctx).Err()

	t := time.Date(2025, 02, 8, 0, 0, 0, 0, time.UTC)

	data := Model{
		Str1:    "hello",
		Str2:    "world",
		Str3:    ToPtr("hello"),
		Str4:    nil,
		Bytes:   []byte("this is bytes !"),
		Int:     123,
		Int2:    ToPtr(0),
		Int3:    nil,
		Bool:    true,
		Bool2:   ToPtr(false),
		Bool3:   nil,
		Time:    t,
		Time2:   ToPtr(t),
		Time3:   nil,
		Ignored: struct{}{},
	}

	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HMSet(ctx, "key", data)
		return nil
	}); err != nil {
		panic(err)
	}

	var model1, model2 Model

	if err := rdb.HGetAll(ctx, "key").Scan(&model1); err != nil {
		panic(err)
	}

	if err := rdb.HMGet(ctx, "key", "str1", "int").Scan(&model2); err != nil {
		panic(err)
	}

	spew.Dump(model1)

	spew.Dump(model2)

}

func ToPtr[T any](v T) *T { _ = "STUB: not implemented"; return nil }
