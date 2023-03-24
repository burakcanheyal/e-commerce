package repository

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"gorm.io/gorm"
)

type KeyRepository struct {
	db *gorm.DB
}

func NewKeyRepository(db *gorm.DB) KeyRepository {
	r := KeyRepository{db}
	return r
}
func (r *KeyRepository) Create(key entity.Key) (entity.Key, error) {
	if err := r.db.Create(&key).Error; err != nil {
		return key, internal.DBNotCreated
	}
	return key, nil
}

func (r *KeyRepository) Delete(key entity.Key) error {
	//Todo:Silinmişlerde arama
	if err := r.db.Model(&key).Where("key_id=?", key.KeyId).Update("status", enum.KeyDeleted).Error; err != nil {
		return internal.DBNotDeleted
	}
	return nil
}
func (r *KeyRepository) GetById(id int32) (entity.Key, error) {
	var key entity.Key
	if err := r.db.Model(&key).Where("status != ", enum.KeyDeleted).Where("key_id=?", id).Scan(&key).Error; err != nil {
		return key, internal.DBNotFound
	}
	return key, nil
}
func (r *KeyRepository) GetByUserId(id int32) (entity.Key, error) {
	var key entity.Key
	if err := r.db.Model(&key).Where("user_id=?", id).Scan(&key).Error; err != nil {
		return key, internal.DBNotFound
	}
	return key, nil
}
func (r *KeyRepository) Update(key entity.Key) error {
	if err := r.db.Model(&key).Where("key_id=?", key.KeyId).Updates(entity.Key{
		Rol: key.Rol,
	}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
