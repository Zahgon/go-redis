package redisprometheus

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/redis/go-redis/v9"
)

type StatGetter interface {
	PoolStats() *redis.PoolStats
}

type Collector struct {
	getter      StatGetter
	hitDesc     *prometheus.Desc
	missDesc    *prometheus.Desc
	timeoutDesc *prometheus.Desc
	totalDesc   *prometheus.Desc
	idleDesc    *prometheus.Desc
	staleDesc   *prometheus.Desc
}

var _ prometheus.Collector = (*Collector)(nil)

func NewCollector(namespace, subsystem string, getter StatGetter) *Collector {
	_ = "STUB: not implemented"
	return nil
}

func (s *Collector) Describe(descs chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }

func (s *Collector) Collect(metrics chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }
