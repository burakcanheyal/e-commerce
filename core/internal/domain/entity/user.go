package entity

import (
	"time"
)

type User struct {
	Id            int32     `gorm:"AUTO_INCREMENT, primaryKey;column:id"`
	Username      string    `gorm:"unique;not null;column:username"`
	Password      string    `gorm:"not null;column:password"`
	Email         string    `gorm:"column:email"`
	Name          string    `gorm:"column:name"`
	Surname       string    `gorm:"column:surname"`
	Phone         string    `gorm:"column:phone"`
	Status        int8      `gorm:"column:status"`
	Code          string    `gorm:"column:code"`
	CodeExpiredAt time.Time `gorm:"not null; column:code_expired_at"`
	BirthDate     time.Time `gorm:"column:birth_date"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at"`
}
