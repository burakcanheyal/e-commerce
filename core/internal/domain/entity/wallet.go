package entity

import "time"

type Wallet struct {
	Id        int32      `gorm:"primary_key;AUTO_INCREMENT;not null"`
	UserId    int32      `gorm:"foreign_key;not null"`
	Balance   float32    `gorm:"column:balance"`
	Status    int8       `gorm:"type:smallint;column:status;not null"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	User      User       `gorm:"foreign_key:UserId"`
}
