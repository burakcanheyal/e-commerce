package dto

import "time"

type SellStaticsDto struct {
	OperationNumber string
	Balance         float64
	OrderId         int32
	ProductName     string
	OrderQuantity   int32
	BuyerName       string
	OperationDate   time.Time
}
