package seed

import (
	"attempt4/core/internal/domain/entity"
	"gorm.io/gorm"
)

func ProductSeed(db *gorm.DB) {
	products := []entity.Product{
		{0, "Domates", 8, 13.2},
		{0, "Patates", 13, 33.2},
		{0, "Soğan", 7, 22.2},
		{0, "Sarımsak", 26, 11.1},
	}
	var size int64
	db.Model(&products).Count(&size)
	if size == 0 {
		for _, p := range products {
			db.Create(&p)
		}
	}
}