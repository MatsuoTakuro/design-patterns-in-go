package state

import "fmt"

// implement state interface
type BaseState struct{}

func (bs *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (bs *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}
