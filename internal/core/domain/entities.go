package domain

type Wallet struct {
	Uuid []byte
	// Сумма на счете в наименьших единицах. Если счет подразумевается, например, долларовый, то это будет сумма в центах и тогда значение 7950 было бы эквивалентно $79.5.
	Amount int64
}

func NewWallet(uuidProvider UuidProvider) *Wallet {
	return &Wallet{uuidProvider.Random(), 0}
}
