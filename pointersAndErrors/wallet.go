package pointersanderrors

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// *Wallet rather than Wallet which you can read as "a pointer to a wallet".
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	balance := w.balance

	if balance < amount {
		return errors.New("oh no")
	}
	w.balance -= amount
	return nil
}

// We get the pointer (memory address) of something by placing an & character at the beginning of the symbol.
