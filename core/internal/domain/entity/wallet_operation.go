package entity

import "time"

type WalletOperation struct {
	Id              int32 `json:"key_id" gorm:"primary_key;AUTO_INCREMENT"`
	OperationNumber string
	Products        string
	Price           float64
	UserId          int32
	OperationDate   time.Time
}
