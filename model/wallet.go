package model

type WalletResponse struct {
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}
type BalanceResponse struct {
	Balance int `json:"balance"`
}

type WalletsResponse []WalletResponse
