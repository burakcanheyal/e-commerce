package entity

import "time"

type Order struct {
	Id        int32      `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	UserId    int32      `gorm:"foreign_key;column:user_id"`
	ProductId int32      `gorm:"foreign_key;column:product_id"`
	Quantity  int32      `gorm:"column:quantity"`
	Status    int8       `gorm:"type:smallint;column:status"`
	Price     float32    `gorm:"column:price"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	Product   Product    `gorm:"foreign_key:ProductId"`
	User      User       `gorm:"foreign_key:UserId"`
}
