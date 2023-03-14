package dto

type ProductDto struct {
	Name     string  `json:"name" validate:"required,gte=1,lte=32"`
	Quantity int32   `json:"quantity" validate:"lte=127,gte=1,number"`
	Price    float32 `json:"price" validate:"gte=1,number,lte=2500"`
}
