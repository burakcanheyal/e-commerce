package entity

type Key struct {
	KeyId  int32 `json:"key_id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId int32 `json:"user_id" gorm:"foreign_key"`
	Rol    int   `json:"rol"`
	Status int8
	User   User `gorm:"foreign_key:UserId"`
}
