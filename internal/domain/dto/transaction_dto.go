package dto

import "time"

type TransactionDto struct {
	OperationNumber string
	Balance         float32
	OrderId         int32
	ProductName     string
	OrderQuantity   int32
	SellerName      string
	OperationDate   time.Time
}
