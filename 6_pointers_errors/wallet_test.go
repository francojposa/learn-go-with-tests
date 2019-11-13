package pointerr

import "testing"

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	want := 10
	got := wallet.Balance()

	if got != want {
		t.Errorf("\nwant: %d, got: %d", want, got)
	}
}
