package entity

import "time"

type Wallet struct {
	Id        int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId    int32     `json:"user_id" gorm:"foreign_key"`
	Balance   float64   `json:"balance"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
	User      User      `gorm:"foreign_key:UserId"`
}
