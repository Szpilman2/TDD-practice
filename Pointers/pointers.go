package main

import (
	"errors"
	"fmt"
)

// important tip: if you want to change state of the struct you should use pointers.
/*
obviously this implementation is incorrect:

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}
*/

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// create new type from existing ones:
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin{
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}