package state

import "fmt"

type OffState struct {
	// Embedded below instead of defining 'BaseState BaseState',
	// any instance of OffState can directly call On(sw *Switch) and Off(sw *Switch).
	// So, it can be seen as it implements State interface.
	BaseState
}

// Overrides On(sw *Switch) implemented by (embedded struct of) BaseState
func (of OffState) On(sw *Switch) {
	fmt.Println("Turning light on...")
	sw.State = NewOnState()
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{
		BaseState: BaseState{},
	}
}
