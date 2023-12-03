package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

// Permite modificar la forma en la que se imprimirá al imprimirlo con %s
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("cant withdraw that amount of money")
	}
	w.balance -= amount
	return nil
}
