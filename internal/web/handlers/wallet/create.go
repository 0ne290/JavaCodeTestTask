package walletHandlers

import (
	"encoding/json"
	"net/http"

	createWallet "github.com/0ne290/JavaCodeTestTask/internal/core/application/wallet/create"
	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
	response "github.com/0ne290/JavaCodeTestTask/internal/web"
)

func Create(uuidProviderFactory func() domain.UuidProvider, unitOfWorkFactory func() domain.UnitOfWork) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletUuid := createWallet.Handle(r.Context(), uuidProviderFactory(), unitOfWorkFactory())

		json.NewEncoder(w).Encode(response.Success(struct {
			WalletId string `json:"walletId"`
		}{walletUuid}))
	}
}
