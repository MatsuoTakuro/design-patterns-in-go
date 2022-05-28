package state

type Switch struct {
	State State
}

func (sw *Switch) On() {
	sw.State.On(sw)
}

func (sw *Switch) Off() {
	sw.State.Off(sw)
}

func NewSwitch() *Switch {
	return &Switch{
		State: NewOffState(),
	}
}
