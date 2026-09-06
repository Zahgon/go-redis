package redis

import "context"

type himportForEach func(ctx context.Context, fn func(ctx context.Context, client *Client) error) error

func (c *ClusterClient) himportRequeueFailedSets(ctx context.Context, cmds []Cmder, failedCmds *cmdsMap) {
	_ = "STUB: not implemented"
	return
}

func himportFanOutPrepare(ctx context.Context, registry *himportRegistry, forEach himportForEach, cmd *HImportPrepareCmd) {
	_ = "STUB: not implemented"
	return
}

func himportFanOutDiscard(ctx context.Context, registry *himportRegistry, forEach himportForEach, cmd *HImportDiscardCmd) {
	_ = "STUB: not implemented"
	return
}

func himportFanOutDiscardAll(ctx context.Context, registry *himportRegistry, forEach himportForEach, cmd *HImportDiscardAllCmd) {
	_ = "STUB: not implemented"
	return
}

func (c *ClusterClient) HImportPrepare(ctx context.Context, fieldsetName string, fields ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) HImportDiscard(ctx context.Context, fieldsetName string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) HImportDiscardAll(ctx context.Context) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ring) HImportPrepare(ctx context.Context, fieldsetName string, fields ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ring) HImportDiscard(ctx context.Context, fieldsetName string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ring) HImportDiscardAll(ctx context.Context) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}
