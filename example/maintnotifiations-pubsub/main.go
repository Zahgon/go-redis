package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/logging"
	"github.com/redis/go-redis/v9/maintnotifications"
)

var ctx = context.Background()
var cntErrors atomic.Int64
var cntSuccess atomic.Int64
var startTime = time.Now()

func main() {
	startTime = time.Now()
	wg := &sync.WaitGroup{}
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode:                       maintnotifications.ModeEnabled,
			EndpointType:               maintnotifications.EndpointTypeExternalIP,
			HandoffTimeout:             10 * time.Second,
			RelaxedTimeout:             10 * time.Second,
			PostHandoffRelaxedDuration: 10 * time.Second,
		},
	})
	_ = rdb.FlushDB(ctx).Err()
	maintnotificationsManager := rdb.GetMaintNotificationsManager()
	if maintnotificationsManager == nil {
		panic("maintnotifications manager is nil")
	}
	loggingHook := maintnotifications.NewLoggingHook(int(logging.LogLevelDebug))
	maintnotificationsManager.AddNotificationHook(loggingHook)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Printf("pool stats: %+v\n", rdb.PoolStats())
		}
	}()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	if err := rdb.Set(ctx, "publishers", "0", 0).Err(); err != nil {
		panic(err)
	}
	if err := rdb.Set(ctx, "subscribers", "0", 0).Err(); err != nil {
		panic(err)
	}
	if err := rdb.Set(ctx, "published", "0", 0).Err(); err != nil {
		panic(err)
	}
	if err := rdb.Set(ctx, "received", "0", 0).Err(); err != nil {
		panic(err)
	}
	fmt.Println("published", rdb.Get(ctx, "published").Val())
	fmt.Println("received", rdb.Get(ctx, "received").Val())
	subCtx, cancelSubCtx := context.WithCancel(ctx)
	pubCtx, cancelPublishers := context.WithCancel(ctx)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go subscribe(subCtx, rdb, "test", i, wg)
	}
	time.Sleep(time.Second)
	cancelSubCtx()
	time.Sleep(time.Second)
	subCtx, cancelSubCtx = context.WithCancel(ctx)
	for i := 0; i < 10; i++ {
		if err := rdb.Incr(ctx, "publishers").Err(); err != nil {
			fmt.Println("incr error:", err)
			cntErrors.Add(1)
		}
		wg.Add(1)
		go floodThePool(pubCtx, rdb, wg)
	}

	for i := 0; i < 500; i++ {
		if err := rdb.Incr(ctx, "subscribers").Err(); err != nil {
			fmt.Println("incr error:", err)
			cntErrors.Add(1)
		}

		wg.Add(1)
		go subscribe(subCtx, rdb, "test2", i, wg)
	}
	time.Sleep(120 * time.Second)
	fmt.Println("canceling publishers")
	cancelPublishers()
	time.Sleep(10 * time.Second)
	fmt.Println("canceling subscribers")
	cancelSubCtx()
	wg.Wait()
	published, err := rdb.Get(ctx, "published").Result()
	received, err := rdb.Get(ctx, "received").Result()
	publishers, err := rdb.Get(ctx, "publishers").Result()
	subscribers, err := rdb.Get(ctx, "subscribers").Result()
	fmt.Printf("publishers: %s\n", publishers)
	fmt.Printf("published: %s\n", published)
	fmt.Printf("subscribers: %s\n", subscribers)
	fmt.Printf("received: %s\n", received)
	publishedInt, err := rdb.Get(ctx, "published").Int()
	subscribersInt, err := rdb.Get(ctx, "subscribers").Int()
	fmt.Printf("if drained = published*subscribers: %d\n", publishedInt*subscribersInt)

	time.Sleep(2 * time.Second)
	fmt.Println("errors:", cntErrors.Load())
	fmt.Println("success:", cntSuccess.Load())
	fmt.Println("time:", time.Since(startTime))
}

func floodThePool(ctx context.Context, rdb *redis.Client, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

func subscribe(ctx context.Context, rdb *redis.Client, topic string, subscriberId int, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}
