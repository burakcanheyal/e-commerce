package dto

type UserDto struct {
	Id        int32
	Username  string `json:"username" validate:"required,gte=1,lte=32"`
	Password  string `json:"password" validate:"required,gte=8,lte=16"`
	Email     string `json:"email" validate:"required,email"`
	Name      string `json:"name" validate:"gte=1,lte=32"`
	Surname   string `json:"surname" validate:"gte=1,lte=32"`
	Phone     string `json:"phone" validate:"gte=9,lte=14"`
	Status    int8
	BirthDate string `json:"birthdate"`
}
