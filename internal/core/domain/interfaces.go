package domain

import "context"

type Repository interface {
	AddWallet(ctx context.Context, wallet *Wallet)
	GetWalletByUuid(ctx context.Context, uuid []byte) (*Wallet, error)
	UpdateWallet(ctx context.Context, wallet *Wallet)
}

type UnitOfWork interface {
	Repository() Repository
	Save(ctx context.Context)
	Rollback(ctx context.Context)
}

type UuidProvider interface {
	Random() []byte
	ToString(uuid []byte) string
	FromString(uuid string) ([]byte, error)
}
