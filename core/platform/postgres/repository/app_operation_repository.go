package repository

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/entity"
	"gorm.io/gorm"
)

type AppOperationRepository struct {
	db *gorm.DB
}

func NewAppOperationRepository(db *gorm.DB) AppOperationRepository {
	p := AppOperationRepository{db}
	return p
}
func (p *AppOperationRepository) Create(panel entity.AppOperation) (entity.AppOperation, error) {
	if err := p.db.Create(&panel).Error; err != nil {
		return panel, internal.DBNotCreated
	}
	return panel, nil
}

func (p *AppOperationRepository) GetById(id int32) (entity.AppOperation, error) {
	var panel entity.AppOperation
	if err := p.db.Model(&panel).Where("id=?", id).First(&panel).Error; err != nil {
		return panel, internal.DBNotFound
	}
	return panel, nil
}
func (p *AppOperationRepository) GetByUserId(id int32) (entity.AppOperation, error) {
	var panel entity.AppOperation
	if err := p.db.Model(&panel).Where("applied_user_id=?", id).First(&panel).Error; err != nil {
		return panel, internal.DBNotFound
	}
	return panel, nil
}

func (p *AppOperationRepository) Update(panel entity.AppOperation) error {
	if err := p.db.Model(&panel).Where("id=?", panel.Id).Updates(
		entity.AppOperation{
			Status:              panel.Status,
			OperationResultDate: panel.OperationResultDate,
		}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
