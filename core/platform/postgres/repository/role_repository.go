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

func (r *RoleRepository) Create(role entity.Role) (entity.Role, error) {
	if err := r.db.Create(&role).Error; err != nil {
		return role, internal.DBNotCreated
	}
	return role, nil
}

func (r *RoleRepository) Delete(role entity.Role) error {
	if err := r.db.Model(&role).Where("rol != ?", enum.RoleDeleted).Where("id=?", role.Id).Update("rol", enum.RoleDeleted).Error; err != nil {
		return internal.DBNotDeleted
	}
	return nil
}

func (r *RoleRepository) GetById(id int32) (entity.Role, error) {
	var key entity.Role
	if err := r.db.Model(&key).Where("status != ", enum.RoleDeleted).Where("key_id=?", id).Scan(&key).Error; err != nil {
		return key, internal.DBNotFound
	}
	return key, nil
}

func (r *RoleRepository) GetByUserId(id int32) (entity.Role, error) {
	var key entity.Role
	if err := r.db.Model(&key).Where("user_id=?", id).Scan(&key).Error; err != nil {
		return key, internal.DBNotFound
	}
	return key, nil
}

func (r *RoleRepository) Update(role entity.Role) error {
	if err := r.db.Model(&role).Where("key_id=?", role.Id).Updates(entity.Role{
		Rol: role.Rol,
	}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
