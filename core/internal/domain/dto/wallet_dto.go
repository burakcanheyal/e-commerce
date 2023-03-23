package dto

type WalletDto struct {
	Balance float64 `json:"balance" validate:"gte=1,number"`
}
