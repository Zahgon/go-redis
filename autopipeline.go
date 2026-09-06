package redis

import (
	"context"
	"errors"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sys/cpu"

	"github.com/redis/go-redis/v9/internal"
	"github.com/redis/go-redis/v9/internal/pool"
)

type AutoPipelineOptions struct {
	MaxBatchSize int

	MaxBatchBytes int

	MaxConcurrentBatches int

	Unordered bool

	contentSharded bool

	NumShards int

	MaxFlushDelay time.Duration

	AdaptiveDelay bool
}

const autoPipelinePermitBackstop = 30 * time.Second

const autoPipelineCloseBackstop = 30 * time.Second

func numAutoPipelineShards() int { _ = "STUB: not implemented"; return 0 }

func DefaultAutoPipelineOptions() *AutoPipelineOptions { _ = "STUB: not implemented"; return nil }

func DefaultBlockingAutoPipelineOptions() *AutoPipelineOptions {
	_ = "STUB: not implemented"
	return nil
}

func (cfg *AutoPipelineOptions) Validate() error { _ = "STUB: not implemented"; return nil }

type cmdableClient interface {
	UniversalClient

	processPipelineHook(ctx context.Context, cmds []Cmder) error

	withProcessPipelineHook(ctx context.Context, cmds []Cmder, hook ProcessPipelineHook) error
	hookCount() int
	withProcessHook(ctx context.Context, cmd Cmder, hook ProcessHook) error
	processPipeline(ctx context.Context, cmds []Cmder) error
	process(ctx context.Context, cmd Cmder) error
}

type apBatch struct {
	done chan struct{}

	closed atomic.Bool

	dispGid atomic.Int64

	nodeMu   sync.Mutex
	nodeGids []int64

	nodeCount atomic.Int32
}

func (b *apBatch) enterNodeDispatch() func() { _ = "STUB: not implemented"; return nil }

func (b *apBatch) isExecutorGoroutine() bool { _ = "STUB: not implemented"; return false }

var noopUnregister = func() {}

func registerBatchExecutors(cmds []Cmder) func() { _ = "STUB: not implemented"; return nil }

func newAPBatch() *apBatch { _ = "STUB: not implemented"; return nil }

func (b *apBatch) close() { _ = "STUB: not implemented"; return }

func (ap *AutoPipeliner) armSelfDeadlockGuard() bool { _ = "STUB: not implemented"; return false }

func curGoroutineID() int64 { _ = "STUB: not implemented"; return 0 }

var queueSlicePool = sync.Pool{
	New: func() interface{} { s := make([]Cmder, 0, 100); return &s },
}

func getQueueSlice(capacity int) []Cmder { _ = "STUB: not implemented"; return nil }

func putQueueSlice(slice []Cmder) { _ = "STUB: not implemented"; return }

type AutoPipeliner struct {
	cmdable

	pipeliner cmdableClient

	pipelinePool pool.Pooler

	cscActiveFn func() bool
	config      *AutoPipelineOptions

	blocking bool

	shards []*apShard
	next   atomic.Uint32

	shardFn func(Cmder) int

	preflight func(ctx context.Context, cmd Cmder) error

	mustDivert func(ctx context.Context, cmd Cmder) bool

	sharedClosed *atomic.Bool

	expectedArrivals atomic.Int64

	execEWMA atomic.Int64

	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
	batchWg sync.WaitGroup

	divertMu sync.Mutex
	divertWg sync.WaitGroup
	closed   atomic.Bool

	closeDone chan struct{}
	closeErr  error
}

const apEnqueueStripes = 8

type apStripe struct {
	mu       sync.Mutex
	queue    []Cmder
	queueLen atomic.Int32

	queueBytes atomic.Int64
	curBatch   *apBatch

	_ cpu.CacheLinePad
}

type apShard struct {
	ap *AutoPipeliner

	next    atomic.Uint32
	stripes []apStripe
	notify  chan struct{}
	sem     *internal.FIFOSemaphore

	inFlight atomic.Int32
}

func (s *apShard) stripe() *apStripe { _ = "STUB: not implemented"; return nil }

func getOrCreateAutoPipeliner(
	mu *sync.Mutex,
	slot **AutoPipeliner,
	closed *bool,
	sharedClosed *atomic.Bool,
	override *AutoPipelineOptions,
	fallback func() *AutoPipelineOptions,
	build func(*AutoPipelineOptions) (*AutoPipeliner, error),
) (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAutoPipeliner(pipeliner cmdableClient, config *AutoPipelineOptions, blocking bool) (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *AutoPipeliner) Do(ctx context.Context, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) runOutsidePipeline(ctx context.Context, cmd Cmder) *apBatch {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) DoRaw(ctx context.Context, args ...interface{}) *RawCmd {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) DoRawWriteTo(ctx context.Context, w io.Writer, args ...interface{}) *RawWriteToCmd {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) Process(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) AddHook(hook Hook) { _ = "STUB: not implemented"; return }

func (ap *AutoPipeliner) DBSize(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (ap *AutoPipeliner) ScriptLoad(ctx context.Context, script string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) ScriptFlush(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) HImportPrepare(ctx context.Context, fieldsetName string, fields ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) HImportDiscard(ctx context.Context, fieldsetName string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) HImportDiscardAll(ctx context.Context) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) Watch(ctx context.Context, fn func(*Tx) error, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) Subscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) PSubscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) SSubscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) PoolStats() *PoolStats { _ = "STUB: not implemented"; return nil }

func (ap *AutoPipeliner) AutoPipeline() (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *AutoPipeliner) AutoPipelineWithOptions(config *AutoPipelineOptions) (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *AutoPipeliner) AsyncAutoPipeline() (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *AutoPipeliner) AsyncAutoPipelineWithOptions(config *AutoPipelineOptions) (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type AutoFuture struct {
	cmd   Cmder
	batch *apBatch
}

func (f AutoFuture) Wait() error { _ = "STUB: not implemented"; return nil }

func (f AutoFuture) WaitContext(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (f AutoFuture) Cmd() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

var outsidePipelineCommands = map[string]struct{}{
	"shutdown": {}, "monitor": {},
	"select": {}, "auth": {}, "hello": {}, "reset": {}, "quit": {},
	"multi": {}, "exec": {}, "discard": {}, "watch": {}, "unwatch": {},
	"subscribe": {}, "unsubscribe": {}, "psubscribe": {}, "punsubscribe": {},
	"ssubscribe": {}, "sunsubscribe": {},
	"client": {},

	"readonly": {}, "readwrite": {}, "asking": {},
}

func runsOutsidePipeline(name string) bool { _ = "STUB: not implemented"; return false }

var blockingCommands = map[string]struct{}{
	"blpop": {}, "brpop": {}, "brpoplpush": {},
	"blmove": {}, "blmovem": {}, "blmpop": {},
	"bzpopmin": {}, "bzpopmax": {}, "bzmpop": {},
	"wait": {}, "waitaof": {},

	"migrate": {},
}

func isBlockingCmd(cmd Cmder) bool { _ = "STUB: not implemented"; return false }

func blockingArgString(arg interface{}) string { _ = "STUB: not implemented"; return "" }

func (ap *AutoPipeliner) submit(ctx context.Context, cmd Cmder) AutoFuture {
	_ = "STUB: not implemented"
	return *new(AutoFuture)
}

var ErrSubmitBlockingFace = errors.New(
	"redis: Submit requires the deferred autopipeliner (AsyncAutoPipeline); on the blocking face use the typed methods or Do",
)

var errZeroAutoFuture = errors.New("redis: Wait on a zero AutoFuture")

var errDoNoArgs = errors.New("redis: AutoPipeliner.Do requires at least one argument")

var ErrAutoPipelineTimeout = errors.New(
	"redis: autopipeline: no batch permit within the internal backstop (engine overloaded or a batch is wedged)",
)

func (ap *AutoPipeliner) Submit(ctx context.Context, cmd Cmder) AutoFuture {
	_ = "STUB: not implemented"
	return *new(AutoFuture)
}

func (ap *AutoPipeliner) processAsync(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AutoPipeliner) processBlocking(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

var completedBatch = func() *apBatch {
	b := newAPBatch()
	b.close()
	return b
}()

func (ap *AutoPipeliner) isClosed() bool { _ = "STUB: not implemented"; return false }

func (ap *AutoPipeliner) enqueue(cmd Cmder) *apBatch { _ = "STUB: not implemented"; return nil }

func (s *apShard) wake() { _ = "STUB: not implemented"; return }

func (ap *AutoPipeliner) IsBlocking() bool { _ = "STUB: not implemented"; return false }

func (ap *AutoPipeliner) Config() AutoPipelineOptions {
	_ = "STUB: not implemented"
	return *new(AutoPipelineOptions)
}

func (ap *AutoPipeliner) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (ap *AutoPipeliner) numShards() int { _ = "STUB: not implemented"; return 0 }

func (ap *AutoPipeliner) setShardFn(fn func(Cmder) int) { _ = "STUB: not implemented"; return }

func (ap *AutoPipeliner) setPreflight(fn func(ctx context.Context, cmd Cmder) error) {
	_ = "STUB: not implemented"
	return
}

func (ap *AutoPipeliner) setMustDivert(fn func(ctx context.Context, cmd Cmder) bool) {
	_ = "STUB: not implemented"
	return
}

func (ap *AutoPipeliner) Close() error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck // handshake, not a critical section

func (ap *AutoPipeliner) WaitClosed() error { _ = "STUB: not implemented"; return nil }

func (ap *AutoPipeliner) drainAll(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *apShard) flusher() { _ = "STUB: not implemented"; return }

func (s *apShard) accumulateBatch() { _ = "STUB: not implemented"; return }

const (
	silenceGapFloor = 200 * time.Microsecond
	silenceGapCeil  = 2 * time.Millisecond
)

const coalesceMinFlush = 8

const stragglerHoldGaps = 3

func (ap *AutoPipeliner) observeBatchExec(d time.Duration) { _ = "STUB: not implemented"; return }

func (ap *AutoPipeliner) silenceGap() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (ap *AutoPipeliner) pipelineHasFreeConn() bool { _ = "STUB: not implemented"; return false }

func (s *apShard) awaitExpectedArrivals(batchSize int) { _ = "STUB: not implemented"; return }

func (ap *AutoPipeliner) dispatchCmds(ctx context.Context, queues [][]Cmder, total int) {
	_ = "STUB: not implemented"
	return
}

func (ap *AutoPipeliner) dispatchCmdsMaybeChunked(ctx context.Context, queues [][]Cmder, total int) {
	_ = "STUB: not implemented"
	return
}

func (ap *AutoPipeliner) dispatchSequential(ctx context.Context, groups [][]Cmder) {
	_ = "STUB: not implemented"
	return
}

func splitRetryRuns(cmds []Cmder) [][]Cmder { _ = "STUB: not implemented"; return nil }

func recoverDispatchPanic(cmds ...[]Cmder) { _ = "STUB: not implemented"; return }

func (s *apShard) flushBatchSlice() { _ = "STUB: not implemented"; return }

func (s *apShard) flushBatchSliceShutdown() { _ = "STUB: not implemented"; return }

func (s *apShard) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *apShard) bytesFull() bool { _ = "STUB: not implemented"; return false }

func cmdApproxBytes(cmd Cmder) int64 { _ = "STUB: not implemented"; return 0 }

func (ap *AutoPipeliner) Len() int { _ = "STUB: not implemented"; return 0 }

func (ap *AutoPipeliner) calculateDelay(queueLen int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (ap *AutoPipeliner) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (ap *AutoPipeliner) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *AutoPipeliner) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *AutoPipeliner) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

var _ Cmdable = (*AutoPipeliner)(nil)
