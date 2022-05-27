package state

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}
