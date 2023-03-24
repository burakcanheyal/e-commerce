package repository

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/entity"
	"gorm.io/gorm"
)

type WalletOperationRepository struct {
	db *gorm.DB
}

func NewWalletOperationRepository(db *gorm.DB) WalletOperationRepository {
	w := WalletOperationRepository{db}
	return w
}
func (w *WalletOperationRepository) Create(walletOperation entity.WalletOperation) (entity.WalletOperation, error) {
	if err := w.db.Create(&walletOperation).Error; err != nil {
		return walletOperation, internal.DBNotCreated
	}
	return walletOperation, nil
}

func (w *WalletOperationRepository) GetById(id int32) (entity.WalletOperation, error) {
	var walletOperation entity.WalletOperation
	if err := w.db.Model(&walletOperation).Where("id=?", id).First(&walletOperation).Error; err != nil {
		return walletOperation, internal.DBNotFound
	}
	return walletOperation, nil
}
func (w *WalletOperationRepository) GetByUserId(id int32) (entity.WalletOperation, error) {
	var walletOperation entity.WalletOperation
	if err := w.db.Model(&walletOperation).Where("user_id=?", id).First(&walletOperation).Error; err != nil {
		return walletOperation, internal.DBNotFound
	}
	return walletOperation, nil
}

func (w *WalletOperationRepository) Update(walletOperation entity.WalletOperation) error {
	if err := w.db.Model(&walletOperation).Where("id=?", walletOperation.Id).Updates(
		entity.WalletOperation{
			Price:         walletOperation.Price,
			Products:      walletOperation.Products,
			OperationDate: walletOperation.OperationDate,
		}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
