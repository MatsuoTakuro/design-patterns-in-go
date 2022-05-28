package command

type CompositeBankAccountCommand struct {
	commands []Command
}

func (c *CompositeBankAccountCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
	}
}

func (c *CompositeBankAccountCommand) Undo() {
	// undo in reverse order
	for i := range c.commands {
		c.commands[len(c.commands)-i-1].Undo()
	}
}

func (c *CompositeBankAccountCommand) Succeeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccountCommand) setSucceeded(suceeded bool) {
	for _, cmd := range c.commands {
		cmd.setSucceeded(suceeded)
	}
}
