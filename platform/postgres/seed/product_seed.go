package seed

import (
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"gorm.io/gorm"
	"time"
)

func ProductSeed(db *gorm.DB) {
	seedUserId := int32(2)
	productNames := [9]string{"İzmir Tour", "Antalya Tour", "Cappadocia Tour", "İstanbul Tour",
		"Ankara Tour", "Çanakkale Tour", "Balıkesir Tour", "Samsun Tour", "Eskişehir Tour"}
	productQuantity := [4]int32{45, 84, 95, 115}
	productPrice := [4]float32{13.2, 8.3, 14.6, 20.5}
	products := []entity.Product{
		{
			0,
			productNames[0],
			productQuantity[0],
			productPrice[0],
			enum.ProductAvailable,
			seedUserId,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
		{
			0,
			productNames[1],
			productQuantity[1],
			productPrice[1],
			enum.ProductAvailable,
			seedUserId,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
		{
			0,
			productNames[2],
			productQuantity[2],
			productPrice[2],
			enum.ProductAvailable,
			seedUserId,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
		{
			0,
			productNames[3],
			productQuantity[3],
			productPrice[3],
			enum.ProductAvailable,
			seedUserId,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
		{
			0,
			productNames[4],
			productQuantity[3],
			productPrice[3],
			enum.ProductAvailable,
			seedUserId,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
		{
			0,
			productNames[5],
			productQuantity[3],
			productPrice[3],
			enum.ProductAvailable,
			seedUserId,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
		{
			0,
			productNames[6],
			productQuantity[3],
			productPrice[3],
			enum.ProductAvailable,
			seedUserId,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
		{
			0,
			productNames[7],
			productQuantity[3],
			productPrice[3],
			enum.ProductAvailable,
			seedUserId,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
		{
			0,
			productNames[8],
			productQuantity[3],
			productPrice[3],
			enum.ProductAvailable,
			seedUserId,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
	}
	var size int64
	db.Model(&products).Count(&size)
	if size == 0 {
		for _, p := range products {
			db.Create(&p)
		}
	}
}
