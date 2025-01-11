package walletHandlers

import (
	"encoding/json"
	"net/http"

	createWallet "github.com/0ne290/JavaCodeTestTask/internal/core/application/wallet/create"
	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
	response "github.com/0ne290/JavaCodeTestTask/internal/web"
)

func Create(uuidProvider domain.UuidProvider, unitOfWork domain.UnitOfWork) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletUuid := createWallet.Handle(r.Context(), uuidProvider, unitOfWork)

		json.NewEncoder(w).Encode(response.Success(struct {
			WalletId string `json:"walletId"`
		}{walletUuid}))
	}
}
