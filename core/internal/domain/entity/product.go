package entity

type Product struct {
	Id       int32   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name     string  `json:"name" gorm:"unique;not null"`
	Quantity int32   `json:"quantity"`
	Price    float32 `json:"price"`
	Status   int8    `json:"status"`
	UserId   int32   `json:"user_id" gorm:"foreign_key"`
	User     User    `gorm:"foreign_key:UserId"`
}
