package pointer_and_error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {

	t.Run("should increase balance in wallet after deposited money amount", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Money(50.5))

		expected := Money(50.5)
		actual := wallet.balance

		assert.Equal(t, expected, actual)
	})
}

func TestBalance(t *testing.T) {

	t.Run("should return zero from empty wallet", func(t *testing.T) {
		wallet := Wallet{}

		expected := Money(0.0)
		actual := wallet.Balance()

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct balance from wallet after deposited money amount", func(t *testing.T) {
		wallet := Wallet{}
		depositAmount := Money(50.5)
		wallet.Deposit(depositAmount)

		expected := Money(50.5)
		actual := wallet.Balance()

		assert.Equal(t, expected, actual)
	})
}

func TestWithdraw(t *testing.T) {
	t.Run("should return withdrawed amount if balance is sufficient", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Money(25.0))

		expected := Money(15.0)
		actual, err := wallet.Withdraw(Money(10.0))

		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})

	t.Run("should return error if balance is NOT sufficient", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Money(25.0))

		expected := Money(0.0)
		actual, err := wallet.Withdraw(Money(50.0))

		assert.Equal(t, expected, actual)
		assert.EqualError(t, err, "insufficient balance")
	})
}
