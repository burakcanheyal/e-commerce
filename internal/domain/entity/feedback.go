package entity

import "time"

type Feedback struct {
	Id          int32      `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Description string     `gorm:"type:varchar(32);not null;column:description"`
	Star        int32      `gorm:"column:star"`
	ProductId   int32      `gorm:"column:product_id"`
	UserId      int32      `gorm:"column:user_id"`
	Status      int8       `gorm:"column:status"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
	Product     Product    `gorm:"foreign_key:product_id"`
	User        User       `gorm:"foreign_key:user_id"`
}
