package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		depositAmount := Bitcoin(10)

		wallet.Deposit(depositAmount)
		got := wallet.GetBalance()
		want := depositAmount

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(5)
		got := wallet.GetBalance()
		want := Bitcoin(5)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
