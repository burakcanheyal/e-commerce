package dto

type AuthDto struct {
	Username string `json:"username" validate:"required,gte=1,lte=32"`
	Password string `json:"password" validate:"required,gte=8,lte=16"`
}
