package state

import "fmt"

type OffState struct {
	BaseState
	// BaseState BaseState
}

// just implement only On(sw *Switch) method in State interface
func (of OffState) On(sw *Switch) {
	fmt.Println("Turning light on...")
	// TODO: why can i assign OnState to State?
	sw.State = NewOnState()
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{
		BaseState: BaseState{},
	}
}
