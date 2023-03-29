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
			Balance:       walletOperation.Balance,
			OperationDate: walletOperation.OperationDate,
		}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}

func (w *WalletOperationRepository) GetAllTransactions(id int32) ([]entity.WalletOperation, int64, error) {
	var transactionList []entity.WalletOperation
	var total int64
	listQuery := w.db.Find(&transactionList).Where("user_id = ?", id)

	if err := listQuery.Count(&total).Find(&transactionList).Error; err != nil {
		return transactionList, 0, err
	}
	return transactionList, total, nil
}
