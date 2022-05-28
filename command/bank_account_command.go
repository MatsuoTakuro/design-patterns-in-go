package command

type BankAccountCommand struct {
	account   *BankAccount
	action    BankAccountAction
	amount    int
	succeeded bool
}

type BankAccountAction int

const (
	Deposit BankAccountAction = iota + 1
	Withdraw
)

func NewBankAccountCommand(account *BankAccount, action BankAccountAction, amount int) *BankAccountCommand {
	return &BankAccountCommand{
		account:   account,
		action:    action,
		amount:    amount,
		succeeded: false,
	}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true

	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}

	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)

	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

// additional members
func (b *BankAccountCommand) Succeeded() bool {
	return b.succeeded
}

func (b *BankAccountCommand) setSucceeded(succeeded bool) {
	b.succeeded = succeeded
}
