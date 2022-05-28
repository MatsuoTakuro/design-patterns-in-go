package command

import "fmt"

func Sub() {
	// command()
	// compositeCommand()
	functionalCommand()
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

func compositeCommand() {
	from := BankAccount{
		balance: 100,
	}
	to := BankAccount{
		balance: 0,
	}
	var mtc Command = NewMoneyTransferCommand(&from, &to, 600)
	mtc.Call()
	fmt.Println("from=", from, "to=", to)

	fmt.Println("Undoing...")
	mtc.Undo()
	fmt.Println("from=", from, "to=", to)
}

func functionalCommand() {
	from := BankAccount{
		balance: 100,
	}
	to := BankAccount{
		balance: 0,
	}
	fm := NewMoneyTransferCommand2(&from, &to)
	fm.Call(1000)
	fmt.Println("from=", from, "to=", to)
}
