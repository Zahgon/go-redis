package redis

import "context"

type himportCmder interface {
	Cmder
	himportCmd()
}

var (
	_ himportCmder = (*HImportPrepareCmd)(nil)
	_ himportCmder = (*HImportSetCmd)(nil)
	_ himportCmder = (*HImportDiscardCmd)(nil)
	_ himportCmder = (*HImportDiscardAllCmd)(nil)
)

type HImportPrepareCmd struct {
	StatusCmd

	fieldsetName string
	fields       []string

	registryVersion uint64
	registryEpoch   uint64
}

func (cmd *HImportPrepareCmd) himportCmd() { _ = "STUB: not implemented"; return }

func NewHImportPrepareCmd(ctx context.Context, fieldsetName string, fields ...string) *HImportPrepareCmd {
	_ = "STUB: not implemented"
	return nil
}

type HImportSetCmd struct {
	StatusCmd

	fieldsetName string
}

func (cmd *HImportSetCmd) himportCmd() { _ = "STUB: not implemented"; return }

func NewHImportSetCmd(ctx context.Context, key, fieldsetName string, values ...interface{}) *HImportSetCmd {
	_ = "STUB: not implemented"
	return nil
}

type HImportDiscardCmd struct {
	IntCmd

	fieldsetName string
}

func (cmd *HImportDiscardCmd) himportCmd() { _ = "STUB: not implemented"; return }

func NewHImportDiscardCmd(ctx context.Context, fieldsetName string) *HImportDiscardCmd {
	_ = "STUB: not implemented"
	return nil
}

type HImportDiscardAllCmd struct {
	IntCmd

	registryEpoch uint64
}

func (cmd *HImportDiscardAllCmd) himportCmd() { _ = "STUB: not implemented"; return }

func NewHImportDiscardAllCmd(ctx context.Context) *HImportDiscardAllCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HImportPrepare(ctx context.Context, fieldsetName string, fields ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HImportSet(ctx context.Context, key, fieldsetName string, values ...interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HImportDiscard(ctx context.Context, fieldsetName string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) HImportDiscardAll(ctx context.Context) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}
