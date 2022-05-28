package state

type SwState int

const (
	Locked SwState = iota // as default
	Failed
	Unlocked
)

func NewSwState() SwState {
	return Locked
}

func (s SwState) String() string {
	switch s {
	case Locked:
		return "LOCKED"
	case Failed:
		return "FAILED"
	case Unlocked:
		return "UNLOCKED"
	}
	return "UNKNOWN"
}
