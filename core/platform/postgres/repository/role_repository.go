package repository

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	r := RoleRepository{db}
	return r
}
func (r *RoleRepository) Create(key entity.Key) (entity.Key, error) {
	if err := r.db.Create(&key).Error; err != nil {
		return key, internal.DBNotCreated
	}
	return key, nil
}

func (r *RoleRepository) Delete(key entity.Key) error {
	if err := r.db.Model(&key).Where("status != ?", enum.RoleDeleted).Where("key_id=?", key.Id).Update("status", enum.RoleDeleted).Error; err != nil {
		return internal.DBNotDeleted
	}
	return nil
}
func (r *RoleRepository) GetById(id int32) (entity.Key, error) {
	var key entity.Key
	if err := r.db.Model(&key).Where("status != ", enum.RoleDeleted).Where("key_id=?", id).Scan(&key).Error; err != nil {
		return key, internal.DBNotFound
	}
	return key, nil
}
func (r *RoleRepository) GetByUserId(id int32) (entity.Key, error) {
	var key entity.Key
	if err := r.db.Model(&key).Where("user_id=?", id).Scan(&key).Error; err != nil {
		return key, internal.DBNotFound
	}
	return key, nil
}
func (r *RoleRepository) Update(key entity.Key) error {
	if err := r.db.Model(&key).Where("key_id=?", key.Id).Updates(entity.Key{
		Rol: key.Rol,
	}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
