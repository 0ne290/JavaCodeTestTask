package infrastructure

import (
	"context"
	"errors"

	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	transaction pgx.Tx
}

func newRepository(transaction pgx.Tx) *Repository {
	return &Repository{transaction}
}

func (repository *Repository) AddWallet(ctx context.Context, wallet *domain.Wallet) {
	if _, err := repository.transaction.Exec(ctx, "INSERT INTO wallets VALUES ($1, $2)", wallet.Uuid, wallet.Balance); err != nil {
		panic(err.Error())
	}
}

func (repository *Repository) GetWalletByUuid(ctx context.Context, uuid []byte) (*domain.Wallet, error) {
	wallet := &domain.Wallet{}

	err := repository.transaction.QueryRow(ctx, "SELECT * FROM wallets WHERE uuid = $1", uuid).Scan(&wallet.Uuid, &wallet.Balance)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("uuid is invalid")
		}

		panic(err.Error())
	}

	return wallet, nil
}

func (repository *Repository) UpdateWallet(ctx context.Context, wallet *domain.Wallet) {
	if _, err := repository.transaction.Exec(ctx, "UPDATE wallets SET balance = $1 WHERE uuid = $2", wallet.Balance, wallet.Uuid); err != nil {
		panic(err.Error())
	}
}
