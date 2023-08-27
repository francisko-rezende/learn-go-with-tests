package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
		// In Go, the fmt package provides formatting verbs like %s that are used to format values according to their respective Stringer interfaces. When you use %s in a format string, the fmt package looks for the String() method in the value you're trying to format and uses that method to get a string representation of the value.
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(5)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {

		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(20))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}
