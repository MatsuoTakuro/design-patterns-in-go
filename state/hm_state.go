package state

type HmState int

const (
	OffHook HmState = iota
	Connecting
	Connected
	OnHold
	OnHook
)

func (s HmState) String() string {
	switch s {
	// case OffHook:
	// 	return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}
