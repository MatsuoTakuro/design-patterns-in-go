package state

type Switch struct {
	State State
}

func (sw *Switch) On() {
	// TODO: why can BaseState.On(sw *Switch) be executed instead if sw.State does not have On(sw *Switch) method?
	sw.State.On(sw)
}

func (sw *Switch) Off() {
	// TODO: why can BaseState.Off(sw *Switch) be executed instead if sw.State does not have Off(sw *Switch) method?
	sw.State.Off(sw)
}

func NewSwitch() *Switch {
	return &Switch{
		// TODO: why can i assign OffState to State?
		State: NewOffState(),
	}
}
