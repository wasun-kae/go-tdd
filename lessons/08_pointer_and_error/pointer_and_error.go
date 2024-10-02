package pointer_and_error

import "errors"

type Money float64

type Wallet struct {
	balance Money
}

func (wallet *Wallet) Deposit(amount Money) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Money {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Money) (Money, error) {
	if wallet.balance < amount {
		return Money(0.0), errors.New("insufficient balance")
	}

	return wallet.balance - amount, nil
}
