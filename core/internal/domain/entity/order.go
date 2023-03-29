package entity

import "time"

// Todo:Order Products diye çoklu data girişi
type Order struct {
	Id        int32     `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	UserId    int32     `gorm:"foreign_key;column:user_id"`
	ProductId int32     `gorm:"foreign_key;column:product_id"`
	Quantity  int32     `gorm:"column:quantity"`
	Status    int8      `gorm:"column:status"`
	Price     float64   `gorm:"column:price"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
	Product   Product   `gorm:"foreign_key:ProductId"`
	User      User      `gorm:"foreign_key:UserId"`
}
