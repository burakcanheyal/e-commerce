package repository

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/entity"
	"gorm.io/gorm"
)

type WalletRepository struct {
	Db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	w := WalletRepository{db}
	return w
}
func (w *WalletRepository) Create(wallet entity.Wallet) (entity.Wallet, error) {
	if err := w.Db.Create(&wallet).Error; err != nil {
		return wallet, internal.DBNotCreated
	}
	return wallet, nil
}

func (w *WalletRepository) Delete(wallet entity.Wallet) error {

	if err := w.Db.Where("id = ?", wallet.Id).Updates(wallet).Error; err != nil {
		return internal.DBNotDeleted
	}
	return nil
}
func (w *WalletRepository) GetById(id int32) (entity.Wallet, error) {
	var wallet entity.Wallet
	if err := w.Db.Model(&wallet).Where("id=?", id).First(&wallet).Error; err != nil {
		return wallet, internal.DBNotFound
	}
	return wallet, nil
}
func (w *WalletRepository) GetByUserId(id int32) (entity.Wallet, error) {
	var wallet entity.Wallet
	if err := w.Db.Model(&wallet).Where("user_id=?", id).First(&wallet).Error; err != nil {
		return wallet, internal.DBNotFound
	}
	return wallet, nil
}
func (w *WalletRepository) Update(wallet entity.Wallet) error {
	if err := w.Db.Model(&wallet).Where("id=?", wallet.Id).Updates(wallet).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}

func (w *WalletRepository) Begin() *gorm.DB {
	return w.Db.Begin()
}
func (w *WalletRepository) Rollback(db *gorm.DB) {
	db.Rollback()
}
func (w *WalletRepository) Commit(db *gorm.DB) {
	db.Commit()
}
