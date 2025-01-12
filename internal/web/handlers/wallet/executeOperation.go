package walletHandlers

import (
	"encoding/json"
	"net/http"

	"github.com/0ne290/JavaCodeTestTask/internal/core/application/wallet/executeWalletOperation"
	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
	response "github.com/0ne290/JavaCodeTestTask/internal/web"
)

func ExecuteOperation(uuidProviderFactory func() domain.UuidProvider, unitOfWorkFactory func() domain.UnitOfWork) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		request := &executeWalletOperation.Request{}

		err := json.NewDecoder(r.Body).Decode(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(response.Fail(struct {
				Message string `json:"message"`
			}{"request body format is invalid"}))

			return
		}

		walletBalance, err := executeWalletOperation.Handle(r.Context(), uuidProviderFactory(), unitOfWorkFactory(), request)

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
