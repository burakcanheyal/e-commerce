package seed

import (
	entity2 "attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"gorm.io/gorm"
)

func RolSeed(db *gorm.DB) {
	firstUserId := int32(1)
	secondUserId := int32(2)
	rol := []entity2.Role{
		{
			0,
			firstUserId,
			enum.RoleAdmin,
			entity2.User{},
		},
		{
			0,
			secondUserId,
			enum.RoleManager,
			entity2.User{},
		},
	}
	var size int64
	db.Model(&rol).Count(&size)
	if size == 0 {
		for _, p := range rol {
			db.Create(&p)
		}
	}
}
