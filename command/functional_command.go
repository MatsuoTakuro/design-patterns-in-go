package command

import "fmt"

func Deposit2(b *BankAccount, amount int) {
	fmt.Println("Deposting", amount)
	b.balance += amount
}

func Withdraw2(b *BankAccount, amount int) {
	fmt.Println("Withdrawing", amount)
	b.balance -= amount
}

type MoneyTransferCommand2 []func(int)

func NewMoneyTransferCommand2(from *BankAccount, to *BankAccount) MoneyTransferCommand2 {
	var mtc MoneyTransferCommand2
	mtc = append(mtc, func(amount int) {
		Withdraw2(from, amount)
	})
	mtc = append(mtc, func(amount int) {
		Deposit2(to, amount)
	})
	return mtc
}

func (mtc MoneyTransferCommand2) Call(amount int) {
	for _, cmd := range mtc {
		cmd(amount)
	}
}
