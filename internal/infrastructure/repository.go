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
	if _, err := repository.transaction.Exec(ctx, "INSERT INTO wallets VALUES ($1, $2)", wallet.Uuid, wallet.Amount); err != nil {
		panic("transaction.Exec() error. Detail :" + err.Error())
	}
}

func (repository *Repository) TryGetWalletByUuid(ctx context.Context, uuid []byte) (*domain.Wallet, error) {
	wallet := &domain.Wallet{}

	err := repository.transaction.QueryRow(ctx, "SELECT * FROM wallets WHERE uuid = $1", uuid).Scan(&wallet.Uuid, &wallet.Amount)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("uuid is invalid")
		}

		panic("row.Scan() error. Detail: " + err.Error())
	}

	return wallet, nil
}
