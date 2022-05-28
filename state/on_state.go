package state

import "fmt"

type OnState struct {
	// defined below instead of defining 'BaseState BaseState',
	// any instance of OnState can directly call On(sw *Switch) and Off(sw *Switch).
	// So, it can be seen as it implements State interface.
	BaseState
}

// Overrides Off(sw *Switch) implemented by (embedded struct of) BaseState
func (on OnState) Off(sw *Switch) {
	fmt.Println("Turning light off...")
	sw.State = NewOffState()
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{
		BaseState: BaseState{},
	}
}
