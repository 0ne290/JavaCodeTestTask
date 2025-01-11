package createWallet

import (
	"context"

	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
)

func Handle(ctx context.Context, uuidProvider domain.UuidProvider, unitOfWork domain.UnitOfWork) string {
	defer unitOfWork.Save(ctx)

	wallet := domain.NewWallet(uuidProvider)

	unitOfWork.Repository().AddWallet(ctx, wallet)

	return uuidProvider.ToString(wallet.Uuid)
}
