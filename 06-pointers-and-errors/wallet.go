package pointersanderrors

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) GetBalance() int {
	return w.balance
}
