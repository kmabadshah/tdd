package main

import (
	"testing"
)

func checkBitcoin(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("GOT %s WANT %s", got, want)
	}
}

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		want := Bitcoin(20)

		checkBitcoin(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		t.Run("Withdraw < Deposit ", func(t *testing.T) {
			_ = wallet.Withdraw(Bitcoin(10))
			want := Bitcoin(10)
			checkBitcoin(t, wallet, want)
		})

		t.Run("Withdraw > Deposit", func(t *testing.T) {
			want := wallet.Balance()
			err := wallet.Withdraw(Bitcoin(50))

			checkBitcoin(t, wallet, want)
			if err == nil {
				t.Errorf("WANT error but didn't get one")
			}
		})
	})
}
