package pointersanderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	// In Go, the fmt package provides formatting verbs like %s that are used to format values according to their respective Stringer interfaces. When you use %s in a format string, the fmt package looks for the String() method in the value you're trying to format and uses that method to get a string representation of the value.
}
