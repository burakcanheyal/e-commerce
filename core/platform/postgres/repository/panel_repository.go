package repository

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/entity"
	"gorm.io/gorm"
)

type PanelRepository struct {
	db *gorm.DB
}

func NewPanelRepository(db *gorm.DB) PanelRepository {
	p := PanelRepository{db}
	return p
}
func (p *PanelRepository) Create(panel entity.Panel) (entity.Panel, error) {
	if err := p.db.Create(&panel).Error; err != nil {
		return panel, internal.DBNotCreated
	}
	return panel, nil
}

func (p *PanelRepository) GetById(id int32) (entity.Panel, error) {
	var panel entity.Panel
	if err := p.db.Model(&panel).Where("id=?", id).First(&panel).Error; err != nil {
		return panel, internal.DBNotFound
	}
	return panel, nil
}

func (p *PanelRepository) Update(panel entity.Panel) error {
	if err := p.db.Model(&panel).Where("id=?", panel.Id).Updates(
		entity.Panel{
			Status: panel.Status,
		}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
