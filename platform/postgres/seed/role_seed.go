package seed

import (
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"gorm.io/gorm"
)

func RolSeed(db *gorm.DB) {
	firstUserId := int32(1)
	secondUserId := int32(2)
	rol := []entity.Role{
		{
			0,
			firstUserId,
			enum.RoleAdmin,
			entity.User{},
		},
		{
			0,
			secondUserId,
			enum.RoleManager,
			entity.User{},
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
