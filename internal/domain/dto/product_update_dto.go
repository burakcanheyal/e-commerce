package dto

type ProductUpdateDto struct {
	Name     string  `json:"name" validate:"required,gte=1,lte=32"`
	Quantity int32   `json:"quantity" validate:"lte=127,gte=0,number"`
	Price    float32 `json:"price" validate:"gte=0,number"`
}
