package redis

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9/internal/pool"
	"github.com/redis/go-redis/v9/internal/proto"
)

type himportFieldset struct {
	fields  []string
	version uint64
}

type himportRegistry struct {
	mu          sync.RWMutex
	nextVersion uint64
	fieldsets   map[string]himportFieldset

	tombstones map[string]struct{}

	discardAllEpoch uint64
}

func newHImportRegistry() *himportRegistry { _ = "STUB: not implemented"; return nil }

func (r *himportRegistry) register(name string, fields []string) (version, epoch uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (r *himportRegistry) lookup(name string) (himportFieldset, bool) {
	_ = "STUB: not implemented"
	return *new(himportFieldset), false
}

func (r *himportRegistry) discard(name string) bool { _ = "STUB: not implemented"; return false }

func (r *himportRegistry) discardAll() (epoch uint64, removed int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (r *himportRegistry) discardVersion(name string, version uint64) {
	_ = "STUB: not implemented"
	return
}

func (r *himportRegistry) refreshVersion(name string, version uint64) {
	_ = "STUB: not implemented"
	return
}

func (r *himportRegistry) idle() bool { _ = "STUB: not implemented"; return false }

func (r *himportRegistry) cleanupSnapshot() (epoch uint64, tombstones []string) {
	_ = "STUB: not implemented"
	return 0, nil
}

func himportNoSuchFieldset(err error) bool { _ = "STUB: not implemented"; return false }

func (c *baseClient) himportInjectedCmds(ctx context.Context, cn *pool.Conn, cmds []Cmder) []Cmder {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) himportReadInjectedReplies(ctx context.Context, cn *pool.Conn, rd *proto.Reader, injected []Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) himportAfterCmd(cn *pool.Conn, hc himportCmder) {
	_ = "STUB: not implemented"
	return
}

func (c *baseClient) himportAfterBatch(cn *pool.Conn, injected []Cmder, cmds []Cmder) {
	_ = "STUB: not implemented"
	return
}

func (c *baseClient) himportRetryFailedSets(ctx context.Context, cn *pool.Conn, cmds []Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) himportShouldRetrySet(cmd Cmder, err error) bool {
	_ = "STUB: not implemented"
	return false
}
