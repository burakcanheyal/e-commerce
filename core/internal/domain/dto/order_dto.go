package dto

type OrderDto struct {
	Id        int32 `json:"id"`
	ProductId int32 `json:"product_id" validate:"required,number,gte=1"`
	Quantity  int32 `json:"quantity" validate:"required,number,gte=1,lte=127"`
}
