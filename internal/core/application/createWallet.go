package application

import (
	"context"

	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
)

func CreateWallet(ctx context.Context, uuidProvider domain.UuidProvider, unitOfWork domain.UnitOfWork) string {
	wallet := domain.NewWallet(uuidProvider)

	unitOfWork.Repository().AddWallet(ctx, wallet)

	unitOfWork.Save(ctx)

	return uuidProvider.ToString(wallet.Uuid)
}
