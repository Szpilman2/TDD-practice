package main

import "testing"


func TestWallet(t *testing.T){

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assureEqual(t, got, want)
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(0)

		assureEqual(t, got, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		got := wallet.Balance()
		want := startingBalance

		assureEqual(t, got, want)

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
	
}


func assureEqual(t testing.TB, got, want Bitcoin){
	t.Helper()

	if got != want {
		t.Errorf("got: %v but want: %v", got, want)
	}
}