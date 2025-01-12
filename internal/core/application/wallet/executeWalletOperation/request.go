package executeWalletOperation

type Request struct {
	WalletId      string `json:"walletId"`
	OperationType string `json:"operationType"`
	Amount        int64  `json:"amount"`
}
