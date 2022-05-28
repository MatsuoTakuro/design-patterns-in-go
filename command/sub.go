package command

import "fmt"

func Sub() {
	command()
}

func command() {
	b := BankAccount{
		balance: 0,
	}
	var cmd Command = NewBankAccountCommand(&b, Deposit, 100)
	cmd.Call()
	fmt.Println(b)

	var cmd2 Command = NewBankAccountCommand(&b, Withdraw, 600)
	cmd2.Call()
	fmt.Println(b)

	cmd2.Undo()
	fmt.Println(b)
}
