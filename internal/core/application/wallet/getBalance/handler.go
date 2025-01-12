package getWalletBalance

import (
	"context"

	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
)

func Handle(ctx context.Context, uuidProvider domain.UuidProvider, unitOfWork domain.UnitOfWork, request *Request) (int64, error) {
	defer unitOfWork.Rollback(ctx)

	walletUuid, err := uuidProvider.FromString(request.WalletId)

	if err != nil {
		return 0, err
	}

	wallet, err := unitOfWork.Repository().GetWalletByUuid(ctx, walletUuid)

	if err != nil {
		return 0, err
	}

	return wallet.Balance, nil
}
