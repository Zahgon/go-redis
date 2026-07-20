package redis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9/internal/routing"
)

var (
	errInvalidCmdPointer         = errors.New("redis: invalid command pointer")
	errNoCmdsToAggregate         = errors.New("redis: no commands to aggregate")
	errNoResToAggregate          = errors.New("redis: no results to aggregate")
	errInvalidCursorCmdArgsCount = errors.New("redis: FT.CURSOR command requires at least 3 arguments")
	errInvalidCursorIdType       = errors.New("redis: invalid cursor ID type")
)

type slotResult struct {
	cmd  Cmder
	keys []string
	err  error
}

func (c *ClusterClient) routeAndRun(ctx context.Context, cmd Cmder, node *clusterNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) executeDefault(ctx context.Context, cmd Cmder, policy *routing.CommandPolicy, node *clusterNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) executeOnArbitraryNode(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) executeOnAllNodes(ctx context.Context, cmd Cmder, policy *routing.CommandPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) executeOnAllShards(ctx context.Context, cmd Cmder, policy *routing.CommandPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) executeMultiShard(ctx context.Context, cmd Cmder, policy *routing.CommandPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) executeMultiSlot(ctx context.Context, cmd Cmder, slotMap map[int][]string, keyOrder []string, policy *routing.CommandPolicy, firstKeyPos int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) createSlotSpecificCommand(ctx context.Context, originalCmd Cmder, keys []string, firstKeyPos int) Cmder {
	_ = "STUB: not implemented"
	return *new(Cmder)
}

func createCommandByType(ctx context.Context, cmdType CmdType, args ...interface{}) Cmder {
	_ = "STUB: not implemented"
	return *new(Cmder)
}

func (c *ClusterClient) executeSpecialCommand(ctx context.Context, cmd Cmder, policy *routing.CommandPolicy, node *clusterNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) executeCursorCommand(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) executeParallel(ctx context.Context, cmd Cmder, nodes []*clusterNode, policy *routing.CommandPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) aggregateMultiSlotResults(ctx context.Context, cmd Cmder, results <-chan slotResult, keyOrder []string, policy *routing.CommandPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) aggregateKeyedValues(cmd Cmder, keyedResults map[string]routing.AggregatorResErr, keyOrder []string, policy *routing.CommandPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) aggregateResponses(cmd Cmder, cmds []Cmder, policy *routing.CommandPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) createAggregator(policy *routing.CommandPolicy, cmd Cmder, isKeyed bool) routing.ResponseAggregator {
	_ = "STUB: not implemented"
	return *new(routing.ResponseAggregator)
}

func (c *ClusterClient) finishAggregation(cmd Cmder, aggregator routing.ResponseAggregator) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) pickArbitraryNode(ctx context.Context) *clusterNode {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) hasKeys(cmd Cmder) bool { _ = "STUB: not implemented"; return false }

func (c *ClusterClient) readOnlyEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ClusterClient) setCommandValue(cmd Cmder, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) setCommandValueReflection(cmd Cmder, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
