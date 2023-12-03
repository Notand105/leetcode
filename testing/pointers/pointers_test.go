package pointers

import (
	"testing"
)

//t.fatal("comment") will stop the tests when its called

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{balance: 0}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("withdraw", func(t *testing.T) {

		wallet := Wallet{balance: 0}
		wallet.Deposit(10)
		wallet.Withdraw(7)

		got := wallet.Balance()
		want := Bitcoin(3)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("withdraw more than balance", func(t *testing.T) {

		wallet := Wallet{balance: 0}
		wallet.Deposit(10)
		err := wallet.Withdraw(Bitcoin(15))

		if err == nil {
			t.Errorf("Didnt get an error")
		}
	})
}
