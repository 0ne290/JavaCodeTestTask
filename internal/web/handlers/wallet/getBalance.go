package walletHandlers

import (
	"encoding/json"
	"net/http"

	getWalletBalance "github.com/0ne290/JavaCodeTestTask/internal/core/application/wallet/getBalance"
	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
	response "github.com/0ne290/JavaCodeTestTask/internal/web"
	"github.com/go-chi/chi/v5"
)

func GetBalance(uuidProviderFactory func() domain.UuidProvider, unitOfWorkFactory func() domain.UnitOfWork) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		request := &getWalletBalance.Request{WalletId: chi.URLParam(r, "walletId")}

		walletBalance, err := getWalletBalance.Handle(r.Context(), uuidProviderFactory(), unitOfWorkFactory(), request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(response.Fail(struct {
				Message string `json:"message"`
			}{err.Error()}))

			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(response.Success(struct {
			WalletBalance int64 `json:"walletBalance"`
		}{walletBalance}))
	}
}
