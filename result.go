package redis

import "time"

func NewCmdResult(val interface{}, err error) *Cmd { _ = "STUB: not implemented"; return nil }

func NewSliceResult(val []interface{}, err error) *SliceCmd { _ = "STUB: not implemented"; return nil }

func NewStatusResult(val string, err error) *StatusCmd { _ = "STUB: not implemented"; return nil }

func NewIntResult(val int64, err error) *IntCmd { _ = "STUB: not implemented"; return nil }

func NewDurationResult(val time.Duration, err error) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewBoolResult(val bool, err error) *BoolCmd { _ = "STUB: not implemented"; return nil }

func NewStringResult(val string, err error) *StringCmd { _ = "STUB: not implemented"; return nil }

func NewFloatResult(val float64, err error) *FloatCmd { _ = "STUB: not implemented"; return nil }

func NewStringSliceResult(val []string, err error) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewBoolSliceResult(val []bool, err error) *BoolSliceCmd { _ = "STUB: not implemented"; return nil }

func NewFloatSliceResult(val []float64, err error) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewMapStringStringResult(val map[string]string, err error) *MapStringStringCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewMapStringIntCmdResult(val map[string]int64, err error) *MapStringIntCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewTimeCmdResult(val time.Time, err error) *TimeCmd { _ = "STUB: not implemented"; return nil }

func NewZSliceCmdResult(val []Z, err error) *ZSliceCmd { _ = "STUB: not implemented"; return nil }

func NewZWithKeyCmdResult(val *ZWithKey, err error) *ZWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewScanCmdResult(keys []string, cursor uint64, err error) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterSlotsCmdResult(val []ClusterSlot, err error) *ClusterSlotsCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoLocationCmdResult(val []GeoLocation, err error) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoPosCmdResult(val []*GeoPos, err error) *GeoPosCmd { _ = "STUB: not implemented"; return nil }

func NewCommandsInfoCmdResult(val map[string]*CommandInfo, err error) *CommandsInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewXMessageSliceCmdResult(val []XMessage, err error) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewXStreamSliceCmdResult(val []XStream, err error) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewXPendingResult(val *XPending, err error) *XPendingCmd {
	_ = "STUB: not implemented"
	return nil
}
