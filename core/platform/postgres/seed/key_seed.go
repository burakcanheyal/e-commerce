package seed

import (
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"gorm.io/gorm"
)

func RolSeed(db *gorm.DB) {
	rol := []entity.Key{
		{0, 1, enum.RoleAdmin, enum.ApprovedKeyStatus, entity.User{}},
		{0, 2, enum.RoleManager, enum.ApprovedKeyStatus, entity.User{}},
	}
	var size int64
	db.Model(&rol).Count(&size)
	if size == 0 {
		for _, p := range rol {
			db.Create(&p)
		}
	}
}
