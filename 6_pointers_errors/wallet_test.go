package pointerr

import "testing"

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	want := Bitcoin(10)
	got := wallet.Balance()

	if got != want {
		t.Errorf("\nwant: %s, got: %s", want, got)
	}
}
