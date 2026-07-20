package redis

import (
	"context"
	"errors"
)

type HotKeysMetric string

const (
	HotKeysMetricCPU HotKeysMetric = "CPU"

	HotKeysMetricNET HotKeysMetric = "NET"
)

type HotKeysStartArgs struct {
	Metrics []HotKeysMetric

	Count uint8

	Duration int64

	Sample int64

	Slots []uint16
}

var ErrHotKeysNoMetrics = errors.New("redis: at least one metric must be specified for HOTKEYS START")

func (c *Client) HotKeysStart(ctx context.Context, args *HotKeysStartArgs) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) HotKeysStop(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Client) HotKeysReset(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) HotKeysGet(ctx context.Context) *HotKeysCmd { _ = "STUB: not implemented"; return nil }
