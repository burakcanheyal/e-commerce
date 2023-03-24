package entity

type Order struct {
	OrderId   int32 `json:"order_id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId    int32 `json:"user_id" gorm:"foreign_key"`
	ProductId int32 `json:"product_id" gorm:"foreign_key"`
	Quantity  int32 `json:"quantity"`
	Status    int8
	Price     float64
	Product   Product `gorm:"foreign_key:ProductId"`
	User      User    `gorm:"foreign_key:UserId"`
}
