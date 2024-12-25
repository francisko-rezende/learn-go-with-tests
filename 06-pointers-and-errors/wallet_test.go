package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.GetBalance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted error but didn't get one")
		}
	}

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

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("expected error but didn't get one")
		}
	})
}
