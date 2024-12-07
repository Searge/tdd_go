package pointers

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

// String returns a human-readable string representation of the amount of Bitcoin
// in the value.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// Deposit adds amount to the current balance of the wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// Balance returns the current balance of the wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
