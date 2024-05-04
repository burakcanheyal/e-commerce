package entity

type Question struct {
	Id          int32  `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Description string `gorm:"column:description"`
	Status      int8   `gorm:"column:status"`
	Topic       int8   `gorm:"column:topic"`
}
