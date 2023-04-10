package dto

type Filter struct {
	Name     string  `form:"name"`
	Quantity int32   `form:"quantity"`
	Price    float32 `form:"price"`
}
