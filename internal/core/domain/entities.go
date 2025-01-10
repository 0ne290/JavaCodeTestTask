package domain

import (
	"fmt"
	"errors"
	"math"
)

type Wallet struct {
	Uuid []byte
	// Сумма на балансе в наименьших единицах. Если счет подразумевается, например, долларовый, то сумма на балансе будет измеряться в центах и тогда значение 7950 будет эквивалентно $79.5.
	Balance uint64
}

func NewWallet(uuidProvider UuidProvider) *Wallet {
	return &Wallet{uuidProvider.Random(), 0}
}

func (wallet *Wallet) Deposit(amount uint64) error {
	if math.MaxUint64 - wallet.Balance < amount {
	    return errors.New(fmt.Sprint("balance cannot be greater than ", uint64(math.MaxUint64)))
	}
	
	wallet.Balance += amount

	return nil
}

func (wallet *Wallet) Withdraw(amount uint64) {
	if amount > wallet.Balance
	    return errors.New("balance cannot be less than 0")
	
	wallet.Balance -= amount

	return nil
}
