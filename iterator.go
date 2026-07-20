package redis

import (
	"context"
)

type ScanIterator struct {
	cmd *ScanCmd
	pos int
}

func (it *ScanIterator) Err() error { _ = "STUB: not implemented"; return nil }

func (it *ScanIterator) Next(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (it *ScanIterator) Val() string { _ = "STUB: not implemented"; return "" }
