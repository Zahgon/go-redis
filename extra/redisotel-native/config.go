package redisotel

import (
	"go.opentelemetry.io/otel/metric"
)

type MetricGroup string

const (
	MetricGroupCommand            MetricGroup = "command"
	MetricGroupConnectionBasic    MetricGroup = "connection-basic"
	MetricGroupResiliency         MetricGroup = "resiliency"
	MetricGroupConnectionAdvanced MetricGroup = "connection-advanced"
	MetricGroupPubSub             MetricGroup = "pubsub"
	MetricGroupStream             MetricGroup = "stream"
)

type HistogramAggregation string

const (
	HistogramAggregationExplicitBucket   HistogramAggregation = "explicit_bucket_histogram"
	HistogramAggregationBase2Exponential HistogramAggregation = "base2_exponential_bucket_histogram"
)

type config struct {
	meterProvider metric.MeterProvider
	enabled       bool

	enabledMetricGroups map[MetricGroup]bool

	includeCommands map[string]bool
	excludeCommands map[string]bool

	hidePubSubChannelNames bool
	hideStreamNames        bool

	histAggregation HistogramAggregation

	bucketsOperationDuration        []float64
	bucketsStreamProcessingDuration []float64
	bucketsConnectionCreateTime     []float64
	bucketsConnectionWaitTime       []float64
}

func (c *config) isMetricGroupEnabled(group MetricGroup) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *config) isCommandIncluded(command string) bool { _ = "STUB: not implemented"; return false }

func defaultHistogramBuckets() []float64 { _ = "STUB: not implemented"; return nil }

type MetricGroupFlags uint32

const (
	MetricGroupFlagCommand            MetricGroupFlags = 1 << 0
	MetricGroupFlagConnectionBasic    MetricGroupFlags = 1 << 1
	MetricGroupFlagResiliency         MetricGroupFlags = 1 << 2
	MetricGroupFlagConnectionAdvanced MetricGroupFlags = 1 << 3
	MetricGroupFlagPubSub             MetricGroupFlags = 1 << 4
	MetricGroupFlagStream             MetricGroupFlags = 1 << 5

	MetricGroupAll MetricGroupFlags = MetricGroupFlagCommand |
		MetricGroupFlagConnectionBasic |
		MetricGroupFlagResiliency |
		MetricGroupFlagConnectionAdvanced |
		MetricGroupFlagPubSub |
		MetricGroupFlagStream
)

type Config struct {
	Enabled       bool
	MeterProvider metric.MeterProvider

	MetricGroups MetricGroupFlags

	IncludeCommands map[string]bool
	ExcludeCommands map[string]bool

	HidePubSubChannelNames bool
	HideStreamNames        bool

	HistogramAggregation HistogramAggregation

	BucketsOperationDuration    []float64
	BucketsStreamLag            []float64
	BucketsConnectionCreateTime []float64
	BucketsConnectionWaitTime   []float64
}

func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) WithEnabled(enabled bool) *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) WithMeterProvider(provider metric.MeterProvider) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) WithMetricGroups(groups MetricGroupFlags) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) WithIncludeCommands(commands []string) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) WithExcludeCommands(commands []string) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) WithHidePubSubChannelNames(hide bool) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) WithHideStreamNames(hide bool) *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) WithHistogramAggregation(agg HistogramAggregation) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) WithHistogramBuckets(buckets []float64) *Config {
	_ = "STUB: not implemented"
	return nil
}
