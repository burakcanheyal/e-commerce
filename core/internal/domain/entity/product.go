package entity

import "time"

type Product struct {
	Id        int32     `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Name      string    `gorm:"unique;not null;column:name"`
	Quantity  int32     `gorm:"column:quantity"`
	Price     float32   `gorm:"column:price"`
	Status    int8      `gorm:"column:status"`
	UserId    int32     `gorm:"foreign_key;column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
	User      User      `gorm:"foreign_key:UserId"`
}
