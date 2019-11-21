package pointerr

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("\nwant: %s, got: %s", want, got)
		}
	}

	assertError := func(t *testing.T, want error, got error) {
		t.Helper()
		if got == nil {
			t.Fatal("Expected err, did not get one")
		}

		if got != want {
			t.Errorf("\nwant: %q, got: %q", want, got)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(15)
		assertBalance(t, wallet, want)
	})

	t.Run("Widthdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})

}
