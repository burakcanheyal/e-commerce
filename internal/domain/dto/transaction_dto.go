package dto

import "time"

type TransactionDto struct {
	OperationNumber string         `json:"operation_number"`
	Balance         float32        `json:"balance"`
	OrderId         string         `json:"order_id"`
	OrderQuantity   int32          `json:"order_quantity"`
	OperationDate   time.Time      `json:"operation_date"`
	Product         ProductTripDto `json:"product"`
}
type TransactionDtoArray struct {
	Transactions []TransactionDto `json:"transactions"`
}
