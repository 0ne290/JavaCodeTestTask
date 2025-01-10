package domain

import (
	"fmt"
	"errors"
	"math"
)

type Wallet struct {
	Uuid []byte
	// Сумма на балансе в наименьших единицах. Если счет подразумевается, например, долларовый, то сумма на балансе будет измеряться в центах и тогда значение 7950 будет эквивалентно $79.5. Изначально я хотел использовать uint64, чтобы увеличить верхнюю границу допустимого диапазона на X2 + 1, но у PostgreSQL проблемы с беззнаковыми типами.
	Balance int64
}

func NewWallet(uuidProvider UuidProvider) *Wallet {
	return &Wallet{uuidProvider.Random(), 0}
}

func (wallet *Wallet) Deposit(amount int64) error {
	if math.MaxInt64 - wallet.Balance < amount {
	    return errors.New(fmt.Sprint("balance cannot be greater than ", math.MaxInt64))
	}
	
	wallet.Balance += amount

	return nil
}

func (wallet *Wallet) Withdraw(amount int64) {
	if amount > wallet.Balance
	    return errors.New("balance cannot be less than 0")
	
	wallet.Balance -= amount

	return nil
}
