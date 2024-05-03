package repository

import (
	"attempt4/internal"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"gorm.io/gorm"
)

type TripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) TripRepository {
	r := TripRepository{db}
	return r
}

func (r *TripRepository) Create(question entity.Trip) (entity.Trip, error) {
	if err := r.db.Create(&question).Error; err != nil {
		return question, internal.DBNotCreated
	}
	return question, nil
}

func (r *TripRepository) Delete(question entity.Trip) error {
	if err := r.db.Model(&question).Where("id=?", question.Id).Update("status", enum.TripPassive).Error; err != nil {
		return internal.DBNotDeleted
	}
	return nil
}

func (r *TripRepository) GetById(id int32) (entity.Trip, error) {
	var key entity.Trip
	if err := r.db.Model(&key).Where("status != ", enum.TripPassive).Where("id=?", id).Scan(&key).Error; err != nil {
		return key, internal.DBNotFound
	}
	return key, nil
}

func (r *TripRepository) GetByProductId(id int32) (entity.Trip, error) {
	var key entity.Trip
	if err := r.db.Model(&key).Where("product_id=?", id).Scan(&key).Error; err != nil {
		return key, internal.DBNotFound
	}
	return key, nil
}

func (r *TripRepository) Update(question entity.Trip) error {
	if err := r.db.Model(&question).Where("id=?", question.Id).Updates(entity.Trip{
		Description: question.Description,
	}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
