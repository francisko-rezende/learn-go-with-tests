package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	checkBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.GetBalance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		depositAmount := Bitcoin(10)
		wallet.Deposit(depositAmount)
		want := depositAmount
		checkBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(5)
		want := Bitcoin(5)
		checkBalance(t, wallet, want)
	})
}
