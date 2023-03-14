package dto

type UserUpdateCodeDto struct {
	Username string `json:"username" validate:"required,gte=1,lte=16"`
	Code     string `json:"code" validate:"required,gte=6,lte=6"`
}
