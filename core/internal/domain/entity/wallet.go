package entity

type Wallet struct {
	Id      int32   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId  int32   `json:"user_id" gorm:"foreign_key"`
	Balance float64 `json:"balance"`
	Status  int8
	User    User `gorm:"foreign_key:UserId"`
}
