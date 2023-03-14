package repository

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	u := UserRepository{db}
	return u
}
func (p *UserRepository) Create(user entity.User) (entity.User, error) {
	if err := p.db.Create(&user).Error; err != nil {
		return user, internal.DBNotCreated
	}
	return user, nil
}

func (p *UserRepository) Delete(user entity.User) error {
	if err := p.db.Model(&user).Where("status != 9").Where("id=?", user.Id).Updates(user).Error; err != nil {
		return internal.DBNotDeleted
	}
	return nil
}
func (p *UserRepository) GetById(id int32) (entity.User, error) {
	var user entity.User
	if err := p.db.Model(&user).Where("status != 9").Where("id=?", id).Scan(&user).Error; err != nil {
		return user, internal.DBNotFound
	}
	return user, nil
}
func (p *UserRepository) GetByName(userName string) (entity.User, error) {
	var user entity.User
	if err := p.db.Model(&user).Where("status != 9").Where("username=?", userName).Scan(&user).Error; err != nil {
		return user, internal.DBNotFound
	}
	return user, nil
}
func (p *UserRepository) Update(user entity.User) error {
	if err := p.db.Model(&user).Where("status != 9").Where("id=?", user.Id).Updates(user).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
