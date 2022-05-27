package state

import "fmt"

type OnState struct {
	BaseState
	// BaseState BaseState
}

// just implement only Off(sw *Switch) method in State interface
func (on OnState) Off(sw *Switch) {
	fmt.Println("Turning light off...")
	// TODO: why can i assign OnState to State?
	sw.State = NewOffState()
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{
		BaseState: BaseState{},
	}
}
