package entity

type Key struct {
	Id     int32 `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	UserId int32 `gorm:"foreign_key;column:user_id"`
	Rol    int   `gorm:"column:rol"`
	User   User  `gorm:"foreign_key:UserId"`
}
