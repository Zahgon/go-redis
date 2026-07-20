package e2e

import (
	"context"
	"net/http"
	"time"
)

type ActionType string

type ActionListItem struct {
	JobID       string `json:"job_id"`
	ActionType  string `json:"action_type"`
	Status      string `json:"status"`
	SubmittedAt string `json:"submitted_at"`
}

const (
	ActionClusterFailover   ActionType = "cluster_failover"
	ActionClusterReshard    ActionType = "cluster_reshard"
	ActionClusterAddNode    ActionType = "cluster_add_node"
	ActionClusterRemoveNode ActionType = "cluster_remove_node"
	ActionClusterMigrate    ActionType = "cluster_migrate"

	ActionNodeRestart ActionType = "node_restart"
	ActionNodeStop    ActionType = "node_stop"
	ActionNodeStart   ActionType = "node_start"
	ActionNodeKill    ActionType = "node_kill"

	ActionNetworkPartition  ActionType = "network_partition"
	ActionNetworkLatency    ActionType = "network_latency"
	ActionNetworkPacketLoss ActionType = "network_packet_loss"
	ActionNetworkBandwidth  ActionType = "network_bandwidth"
	ActionNetworkRestore    ActionType = "network_restore"

	ActionConfigChange    ActionType = "config_change"
	ActionMaintenanceMode ActionType = "maintenance_mode"
	ActionSlotMigration   ActionType = "slot_migrate"

	ActionSequence       ActionType = "sequence_of_actions"
	ActionExecuteCommand ActionType = "execute_command"

	ActionDeleteDatabase ActionType = "delete_database"
	ActionCreateDatabase ActionType = "create_database"
	ActionFailover       ActionType = "failover"
	ActionMigrate        ActionType = "migrate"
	ActionBind           ActionType = "bind"

	ActionSlotMigrate ActionType = "slot_migrate"
)

type SlotMigrateEffect string

const (
	SlotMigrateEffectRemoveAdd SlotMigrateEffect = "remove-add"

	SlotMigrateEffectRemove SlotMigrateEffect = "remove"

	SlotMigrateEffectAdd SlotMigrateEffect = "add"

	SlotMigrateEffectSlotShuffle SlotMigrateEffect = "slot-shuffle"
)

type SlotMigrateVariant string

const (
	SlotMigrateVariantDefault SlotMigrateVariant = "default"

	SlotMigrateVariantMigrate SlotMigrateVariant = "migrate"

	SlotMigrateVariantFailover SlotMigrateVariant = "failover"
)

type SlotMigrateRequest struct {
	Effect       SlotMigrateEffect  `json:"effect"`
	BdbID        string             `json:"bdb_id"`
	ClusterIndex int                `json:"cluster_index,omitempty"`
	Trigger      SlotMigrateVariant `json:"variant,omitempty"`
	SourceNode   *int               `json:"source_node,omitempty"`
	TargetNode   *int               `json:"target_node,omitempty"`
}

type SlotMigrateTrigger struct {
	Name         string                          `json:"name"`
	Description  string                          `json:"description"`
	Requirements []SlotMigrateTriggerRequirement `json:"requirements"`
}

type SlotMigrateTriggerRequirement struct {
	DBConfig    map[string]interface{} `json:"dbconfig"`
	Cluster     map[string]interface{} `json:"cluster"`
	Description string                 `json:"description"`
}

type SlotMigrateTriggersResponse struct {
	Effect   SlotMigrateEffect      `json:"effect"`
	Cluster  map[string]interface{} `json:"cluster"`
	Triggers []SlotMigrateTrigger   `json:"triggers"`
}

type ActionStatus string

const (
	StatusPending   ActionStatus = "pending"
	StatusRunning   ActionStatus = "running"
	StatusFinished  ActionStatus = "finished"
	StatusFailed    ActionStatus = "failed"
	StatusSuccess   ActionStatus = "success"
	StatusCancelled ActionStatus = "cancelled"
)

type ActionRequest struct {
	Type       ActionType             `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

type ActionResponse struct {
	ActionID string `json:"action_id"`
	Status   string `json:"status"`
	Message  string `json:"message,omitempty"`
}

type ActionStatusResponse struct {
	ActionID  string                 `json:"action_id"`
	Status    ActionStatus           `json:"status"`
	Error     interface{}            `json:"error,omitempty"`
	Output    map[string]interface{} `json:"output,omitempty"`
	Progress  float64                `json:"progress,omitempty"`
	StartTime time.Time              `json:"start_time,omitempty"`
	EndTime   time.Time              `json:"end_time,omitempty"`
}

type SequenceAction struct {
	Type       ActionType             `json:"type"`
	Parameters map[string]interface{} `json:"params,omitempty"`
	Delay      time.Duration          `json:"delay,omitempty"`
}

type FaultInjectorClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewFaultInjectorClient(baseURL string) *FaultInjectorClient {
	_ = "STUB: not implemented"
	return nil
}

func (c *FaultInjectorClient) GetBaseURL() string { _ = "STUB: not implemented"; return "" }

type ActionsListResponse struct {
	Actions []ActionListItem `json:"actions"`
}

func (c *FaultInjectorClient) ListActions(ctx context.Context) ([]ActionListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) TriggerAction(ctx context.Context, action ActionRequest) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) TriggerSequence(ctx context.Context, bdbID int, actions []SequenceAction) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) GetActionStatus(ctx context.Context, actionID string) (*ActionStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) WaitForAction(ctx context.Context, actionID string, options ...WaitOption) (*ActionStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) TriggerClusterFailover(ctx context.Context, nodeID string, force bool) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) TriggerClusterReshard(ctx context.Context, slots []int, sourceNode, targetNode string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) TriggerSlotMigration(ctx context.Context, startSlot, endSlot int, sourceNode, targetNode string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) GetSlotMigrateTriggers(ctx context.Context, effect SlotMigrateEffect, clusterIndex int) (*SlotMigrateTriggersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) GetSlotMigrateTriggersRaw(ctx context.Context, effect SlotMigrateEffect, clusterIndex int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) TriggerSlotMigrate(ctx context.Context, req SlotMigrateRequest) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) TriggerSlotMigrateRemoveAdd(ctx context.Context, bdbID string, trigger SlotMigrateVariant) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) TriggerSlotMigrateRemove(ctx context.Context, bdbID string, trigger SlotMigrateVariant) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) TriggerSlotMigrateAdd(ctx context.Context, bdbID string, trigger SlotMigrateVariant) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) TriggerSlotMigrateSlotShuffle(ctx context.Context, bdbID string, trigger SlotMigrateVariant) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) RestartNode(ctx context.Context, nodeID string, graceful bool) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) StopNode(ctx context.Context, nodeID string, graceful bool) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) StartNode(ctx context.Context, nodeID string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) KillNode(ctx context.Context, nodeID string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) SimulateNetworkPartition(ctx context.Context, nodes []string, duration time.Duration) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) SimulateNetworkLatency(ctx context.Context, nodes []string, latency time.Duration, jitter time.Duration) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) SimulatePacketLoss(ctx context.Context, nodes []string, lossPercent float64) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) LimitBandwidth(ctx context.Context, nodes []string, bandwidth string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) RestoreNetwork(ctx context.Context, nodes []string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) ChangeConfig(ctx context.Context, nodeID string, config map[string]string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) EnableMaintenanceMode(ctx context.Context, nodeID string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) DisableMaintenanceMode(ctx context.Context, nodeID string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DatabaseConfig struct {
	Name                         string                 `json:"name"`
	Port                         int                    `json:"port"`
	MemorySize                   int64                  `json:"memory_size"`
	Replication                  bool                   `json:"replication"`
	EvictionPolicy               string                 `json:"eviction_policy"`
	Sharding                     bool                   `json:"sharding"`
	AutoUpgrade                  bool                   `json:"auto_upgrade"`
	ShardsCount                  int                    `json:"shards_count"`
	ModuleList                   []DatabaseModule       `json:"module_list,omitempty"`
	OSSCluster                   bool                   `json:"oss_cluster"`
	OSSClusterAPIPreferredIPType string                 `json:"oss_cluster_api_preferred_ip_type,omitempty"`
	ProxyPolicy                  string                 `json:"proxy_policy,omitempty"`
	ShardsPlacement              string                 `json:"shards_placement,omitempty"`
	ShardKeyRegex                []ShardKeyRegexPattern `json:"shard_key_regex,omitempty"`
}

type DatabaseModule struct {
	ModuleArgs string `json:"module_args"`
	ModuleName string `json:"module_name"`
}

type ShardKeyRegexPattern struct {
	Regex string `json:"regex"`
}

func (c *FaultInjectorClient) DeleteDatabase(ctx context.Context, clusterIndex int, bdbID int) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) CreateDatabase(ctx context.Context, clusterIndex int, databaseConfig DatabaseConfig) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) CreateDatabaseFromMap(ctx context.Context, clusterIndex int, databaseConfig map[string]interface{}) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isPortUnavailableError(err error) bool { _ = "STUB: not implemented"; return false }

func (c *FaultInjectorClient) CreateDatabaseWithPortRetry(ctx context.Context, clusterIndex int, databaseConfig map[string]interface{}, maxRetries int) (*ActionResponse, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (c *FaultInjectorClient) CreateDatabaseConfigWithPortRetry(ctx context.Context, clusterIndex int, databaseConfig DatabaseConfig, maxRetries int) (*ActionResponse, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (c *FaultInjectorClient) ExecuteSequence(ctx context.Context, actions []SequenceAction) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) ExecuteCommand(ctx context.Context, nodeID, command string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) SimulateClusterUpgrade(ctx context.Context, nodes []string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FaultInjectorClient) SimulateNetworkIssues(ctx context.Context, nodes []string) (*ActionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type waitConfig struct {
	pollInterval time.Duration
	maxWaitTime  time.Duration
}

type WaitOption func(*waitConfig)

func WithPollInterval(interval time.Duration) WaitOption {
	_ = "STUB: not implemented"
	return *new(WaitOption)
}

func WithMaxWaitTime(maxWait time.Duration) WaitOption {
	_ = "STUB: not implemented"
	return *new(WaitOption)
}

func (c *FaultInjectorClient) request(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func GetClusterNodes() []string { _ = "STUB: not implemented"; return nil }

func GetMasterNodes() []string { _ = "STUB: not implemented"; return nil }

func GetSlaveNodes() []string { _ = "STUB: not implemented"; return nil }

func ParseNodeID(nodeAddr string) string { _ = "STUB: not implemented"; return "" }

func FormatSlotRange(start, end int) string { _ = "STUB: not implemented"; return "" }

func DebugE2E() bool { _ = "STUB: not implemented"; return false }

func formatSMigratingNotification(seqID int64, slots ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func formatSMigratedNotification(seqID int64, endpoints ...string) string {
	_ = "STUB: not implemented"
	return ""
}
