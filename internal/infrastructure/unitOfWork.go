package infrastructure

import (
	"context"

	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UnitOfWork struct {
	repository *Repository
}

func NewUnitOfWork(ctx context.Context, pool *pgxpool.Pool) *UnitOfWork {
	transaction, err := pool.Begin(ctx)

	if err != nil {
		panic(err.Error())
	}

	return &UnitOfWork{newRepository(transaction)}
}

func (unitOfWork *UnitOfWork) Repository() domain.Repository {
	return unitOfWork.repository
}

func (unitOfWork *UnitOfWork) Save(ctx context.Context) {
	if err := unitOfWork.repository.transaction.Commit(ctx); err != nil {
		panic(err.Error())
	}
}

func (unitOfWork *UnitOfWork) Rollback(ctx context.Context) {
	if err := unitOfWork.repository.transaction.Rollback(ctx); err != nil {
		panic(err.Error())
	}
}
