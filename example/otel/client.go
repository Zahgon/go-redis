package main

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/uptrace/uptrace-go/uptrace"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

var tracer = otel.Tracer("github.com/redis/go-redis/example/otel")

func main() {
	ctx := context.Background()

	uptrace.ConfigureOpentelemetry(

		uptrace.WithDSN("http://project1_secret@localhost:14318/2?grpc=14317"),

		uptrace.WithServiceName("myservice"),
		uptrace.WithServiceVersion("v1.0.0"),
	)
	defer uptrace.Shutdown(ctx)

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}

	for i := 0; i < 1e6; i++ {
		ctx, rootSpan := tracer.Start(ctx, "handleRequest")

		if err := handleRequest(ctx, rdb); err != nil {
			rootSpan.RecordError(err)
			rootSpan.SetStatus(codes.Error, err.Error())
		}

		rootSpan.End()

		if i == 0 {
			fmt.Printf("view trace: %s\n", uptrace.TraceURL(rootSpan))
		}

		time.Sleep(time.Second)
	}
}

func handleRequest(ctx context.Context, rdb *redis.Client) error {
	_ = "STUB: not implemented"
	return nil
}
