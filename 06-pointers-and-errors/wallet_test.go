package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		depositAmount := Bitcoin(10)
		wallet.Deposit(depositAmount)
		want := depositAmount
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(5)
		want := Bitcoin(5)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw more than what the wallet has", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(20))

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("expected error but didn't get one")
		}
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.GetBalance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("%q got want %q", got, want)
	}
}
