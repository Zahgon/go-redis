package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	redisotel "github.com/redis/go-redis/extra/redisotel-native/v9"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	ctx := context.Background()

	exporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to create OTLP exporter: %v", err)
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(
			metric.NewPeriodicReader(exporter,
				metric.WithInterval(10*time.Second),
			),
		),
	)
	defer func() {
		if err := meterProvider.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down meter provider: %v", err)
		}
	}()

	otel.SetMeterProvider(meterProvider)

	otelInstance := redisotel.GetObservabilityInstance()
	config := redisotel.NewConfig().WithEnabled(true)
	if err := otelInstance.Init(config); err != nil {
		log.Fatalf("Failed to initialize OTel: %v", err)
	}
	defer otelInstance.Shutdown()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	log.Println("Executing Redis operations...")
	var wg sync.WaitGroup
	wg.Add(50)
	for i := range 50 {
		go func(i int) {
			defer wg.Done()

			for j := range 10 {
				if err := rdb.Set(ctx, "key"+strconv.Itoa(i*10+j), "value", 0).Err(); err != nil {
					log.Printf("Error setting key: %v", err)
				}
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
			}
		}(i)
	}
	wg.Wait()

	wg.Add(10)
	for i := range 10 {
		go func(i int) {
			defer wg.Done()

			for j := range 10 {
				if err := rdb.Set(ctx, "key"+strconv.Itoa(i*10+j), "value", 0).Err(); err != nil {
					log.Printf("Error setting key: %v", err)
				}
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
			}
		}(i)
	}
	wg.Wait()

	for j := range 10 {
		if err := rdb.Set(ctx, "key"+strconv.Itoa(j), "value", 0).Err(); err != nil {
			log.Printf("Error setting key: %v", err)
		}
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
	}

	log.Println("Operations complete. Waiting for metrics to be exported...")

	time.Sleep(15 * time.Second)

}
