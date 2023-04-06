package seed

import (
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"gorm.io/gorm"
	"time"
)

func WalletSeed(db *gorm.DB) {
	walletUserId := [2]int32{1, 2}
	walletBalance := [2]float32{15100, 12000}
	wallets := []entity.Wallet{
		{
			0,
			walletUserId[0],
			walletBalance[0],
			enum.WalletActive,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
		{
			0,
			walletUserId[1],
			walletBalance[1],
			enum.WalletActive,
			time.Now(),
			nil,
			nil,
			entity.User{},
		},
	}
	var size int64
	db.Model(&wallets).Count(&size)
	if size == 0 {
		for _, w := range wallets {
			db.Create(&w)
		}
	}
}
