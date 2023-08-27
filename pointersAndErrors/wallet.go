package pointersanderrors

import "fmt"

type Wallet struct {
	balance int
}

// *Wallet rather than Wallet which you can read as "a pointer to a wallet".
func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// We get the pointer (memory address) of something by placing an & character at the beginning of the symbol.
