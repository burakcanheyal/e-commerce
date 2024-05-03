package entity

type Trip struct {
	Id              int32   `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Description     string  `gorm:"column:description"`
	ProductId       int32   `gorm:"foreign_key;column:product_id"`
	Name            string  `gorm:"column:name"`
	Status          int32   `gorm:"column:status"`
	Lat             float32 `gorm:"column:lat"`
	Lng             float32 `gorm:"column:lng"`
	NaturePoint     int8    `gorm:"column:nature_point"`
	HistoricalPoint int8    `gorm:"column:historical_point"`
	AdventurePoint  int8    `gorm:"column:adventure_point"`
	CulturalPoint   int8    `gorm:"column:cultural_point"`
	RelaxPoint      int8    `gorm:"column:relax_point"`
}
