package dto

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserToken struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
