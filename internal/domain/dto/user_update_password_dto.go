package dto

type UserUpdatePasswordDto struct {
	UserName    string `json:"username" validate:"required,gte=1,lte=16"`
	Password    string `json:"password" validate:"required,gte=8,lte=16"`
	NewPassword string `json:"new_password" validate:"required,gte=8,lte=16"`
}
