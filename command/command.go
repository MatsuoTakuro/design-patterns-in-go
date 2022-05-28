package command

type Command interface {
	Call()
	Undo()
}
