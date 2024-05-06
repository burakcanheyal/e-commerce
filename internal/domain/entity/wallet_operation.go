package entity

import "time"

type WalletOperation struct {
	Id              int32     `gorm:"primary_key;AUTO_INCREMENT;column:id;not null"`
	OperationNumber string    `gorm:"type:varchar(16);column:operation_number;not null"`
	Type            int8      `gorm:"type:smallint;column:type;not null"`
	Balance         float32   `gorm:"column:balance;not null"`
	UserId          int32     `gorm:"foreign_key;column:user_id"`
	OrderId         string    `gorm:"foreign_key;column:order_id"`
	OperationDate   time.Time `gorm:"column:operation_date"`
}
