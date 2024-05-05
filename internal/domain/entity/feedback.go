package entity

import "time"

type Feedback struct {
	Id          int32      `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Description string     `gorm:"type:varchar(32);not null;column:description"`
	Star        int32      `gorm:"column:star"`
	TripId      int32      `gorm:"column:trip_id"`
	UserId      int32      `gorm:"column:user_id"`
	Status      int8       `gorm:"column:status"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
	Trip        Trip       `gorm:"foreign_key:trip_id"`
	User        User       `gorm:"foreign_key:user_id"`
}
