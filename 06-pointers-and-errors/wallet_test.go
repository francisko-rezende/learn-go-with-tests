package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	depositAmount := Bitcoin(10)

	wallet.Deposit(depositAmount)
	got := wallet.GetBalance()
	want := depositAmount

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
