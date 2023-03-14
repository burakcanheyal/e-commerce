package entity

import (
	"time"
)

type User struct {
	Id            int32     `json:"id" gorm:"AUTO_INCREMENT, primaryKey"`
	Username      string    `json:"userName" gorm:"unique;not null"`
	Password      string    `json:"pass" gorm:"not null"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	Surname       string    `json:"surname"`
	Role          int8      `json:"role"`
	Status        int8      `json:"status"`
	Code          string    `json:"code"`
	CodeExpiredAt time.Time `gorm:"not null"`
	BirthDate     time.Time `json:"birthDate"`
}
