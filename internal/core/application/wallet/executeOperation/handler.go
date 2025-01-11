package executeOperation

import (
	"context"
	"errors"

	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
)

func Handle(ctx context.Context, uuidProvider domain.UuidProvider, unitOfWork domain.UnitOfWork, request *Request) (int64, error) {
	walletUuid, err := uuidProvider.FromString(request.WalletId)

	if err != nil {
		unitOfWork.Rollback(ctx)

		return 0, err
	}

	repository := unitOfWork.Repository()

	wallet, err := repository.GetWalletByUuid(ctx, walletUuid)

	if err != nil {
		unitOfWork.Rollback(ctx)

		return 0, err
	}

	switch request.OperationType {
	case "DEPOSIT":
		err := wallet.Deposit(request.Amount)

		if err != nil {
			unitOfWork.Rollback(ctx)

			return 0, err
		}

		repository.UpdateWallet(ctx, wallet)

		unitOfWork.Save(ctx)

		return wallet.Balance, nil

	case "WITHDRAW":
		err := wallet.Withdraw(request.Amount)

		if err != nil {
			unitOfWork.Rollback(ctx)

			return 0, err
		}

		repository.UpdateWallet(ctx, wallet)

		unitOfWork.Save(ctx)

		return wallet.Balance, nil

	default:
		unitOfWork.Rollback(ctx)

		return 0, errors.New("operation is unsupported")
	}
}
