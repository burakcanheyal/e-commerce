package seed

import (
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"gorm.io/gorm"
)

func WalletSeed(db *gorm.DB) {
	wallets := []entity.Wallet{
		{0, 1, 3522.10, enum.WalletActive, entity.User{}},
		{0, 2, 1347.9, enum.WalletActive, entity.User{}},
	}
	var size int64
	db.Model(&wallets).Count(&size)
	if size == 0 {
		for _, w := range wallets {
			db.Create(&w)
		}
	}
}