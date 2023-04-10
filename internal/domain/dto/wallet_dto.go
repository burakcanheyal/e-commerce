package dto

type WalletDto struct {
	Balance float32 `json:"balance" validate:"gte=1,number"`
}
