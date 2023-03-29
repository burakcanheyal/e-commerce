package entity

import "time"

type WalletOperation struct {
	Id              int32     `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	OperationNumber string    `gorm:"column:operation_number"`
	Type            int8      `gorm:"column:type"`
	Balance         float64   `gorm:"column:balance"`
	UserId          int32     `gorm:"foreign_key;column:user_id"`
	ProductId       int32     `gorm:"foreign_key;column:product_id"`
	OrderId         int32     `gorm:"foreign_key;column:order_id"`
	OperationDate   time.Time `gorm:"column:operation_date"`
	Order           Order     `gorm:"foreign_key:OrderId"`
	User            User      `gorm:"foreign_key:UserId"`
	Product         Product   `gorm:"foreign_key:ProductId"`
}
