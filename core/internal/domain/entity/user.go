package entity

import (
	"time"
)

type User struct {
	Id            int32      `gorm:"AUTO_INCREMENT;primaryKey;column:id"`
	Username      string     `gorm:"type:varchar(32);unique;not null;column:username"`
	Password      string     `gorm:"not null;column:password"`
	Email         string     `gorm:"column:email;not null"`
	Name          string     `gorm:"type:varchar(32);column:name;not null"`
	Surname       string     `gorm:"type:varchar(32);column:surname;not null"`
	Phone         string     `gorm:"type:varchar(14);column:phone"`
	Status        int8       `gorm:"type:smallint;column:status;not null"`
	Code          *string    `gorm:"column:code"`
	CodeExpiredAt *time.Time `gorm:"column:code_expired_at"`
	BirthDate     *time.Time `gorm:"column:birth_date"`
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at"`
}
