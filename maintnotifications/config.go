package maintnotifications

import (
	"context"
	"net"
	"time"
)

type Mode string

const (
	ModeDisabled Mode = "disabled"
	ModeEnabled  Mode = "enabled"
	ModeAuto     Mode = "auto"
)

func (m Mode) IsValid() bool { _ = "STUB: not implemented"; return false }

func (m Mode) String() string { _ = "STUB: not implemented"; return "" }

type EndpointType string

const (
	EndpointTypeAuto         EndpointType = "auto"
	EndpointTypeInternalIP   EndpointType = "internal-ip"
	EndpointTypeInternalFQDN EndpointType = "internal-fqdn"
	EndpointTypeExternalIP   EndpointType = "external-ip"
	EndpointTypeExternalFQDN EndpointType = "external-fqdn"
	EndpointTypeNone         EndpointType = "none"
)

func (e EndpointType) IsValid() bool { _ = "STUB: not implemented"; return false }

func (e EndpointType) String() string { _ = "STUB: not implemented"; return "" }

type Config struct {
	Mode Mode

	EndpointType EndpointType

	RelaxedTimeout time.Duration

	HandoffTimeout time.Duration

	MaxWorkers int

	HandoffQueueSize int

	PostHandoffRelaxedDuration time.Duration

	CircuitBreakerFailureThreshold int

	CircuitBreakerResetTimeout time.Duration

	CircuitBreakerMaxRequests int

	MaxHandoffRetries int
}

func (c *Config) IsEnabled() bool { _ = "STUB: not implemented"; return false }

func DefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *Config) ApplyDefaults() *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) ApplyDefaultsWithPoolSize(poolSize int) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) ApplyDefaultsWithPoolConfig(poolSize int, maxActiveConns int) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) Clone() *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) applyWorkerDefaults(poolSize int) { _ = "STUB: not implemented"; return }

const endpointDetectResolveTimeout = 2 * time.Second

var cgnatNet = &net.IPNet{IP: net.IPv4(100, 64, 0, 0), Mask: net.CIDRMask(10, 32)}

func isPrivateIP(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func DetectEndpointType(addr string, tlsEnabled bool) EndpointType {
	_ = "STUB: not implemented"
	return *new(EndpointType)
}

func isInternalHostname(ctx context.Context, hostname string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
