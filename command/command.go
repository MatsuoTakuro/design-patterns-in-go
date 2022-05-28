package command

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	setSucceeded(suceeded bool)
}
