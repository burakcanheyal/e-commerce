package seed

import (
	"attempt4/platform/app_log/entity"
	"gorm.io/gorm"
	"time"
)

func ApplicationLogSeed(db *gorm.DB) {
	logs := []entity.ApplicationLog{
		entity.ApplicationLog{
			Id:           0,
			UserId:       1,
			LogType:      "Error",
			Content:      "Deneme",
			RelatedTable: "Seed",
			CreatedAt:    time.Now(),
		},
		entity.ApplicationLog{
			Id:           0,
			UserId:       2,
			LogType:      "Info",
			Content:      "Kullanıcı yaratıldı",
			RelatedTable: "Seed",
			CreatedAt:    time.Now(),
		},
	}

	var size int64
	db.Model(&logs).Count(&size)
	if size == 0 {
		for _, l := range logs {
			db.Create(&l)
		}
	}
}
