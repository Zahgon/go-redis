package maintnotifications

type State int

const (
	StateIdle State = iota

	StateMoving
)

func (s State) String() string { _ = "STUB: not implemented"; return "" }
