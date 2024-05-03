package repository

import (
	"attempt4/internal"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"gorm.io/gorm"
)

type PointRepository struct {
	db *gorm.DB
}

func NewPointRepository(db *gorm.DB) PointRepository {
	p := PointRepository{db}
	return p
}

func (p *PointRepository) Create(points entity.TopicPoints) (entity.TopicPoints, error) {
	if err := p.db.Create(&points).Error; err != nil {
		return points, internal.DBNotCreated
	}
	return points, nil
}

func (p *PointRepository) Delete(points entity.TopicPoints) error {
	if err := p.db.Model(&points).Where("status != ?", enum.UserDeletedStatus).Where("id=?", points.Id).Update("status", enum.UserDeletedStatus).Error; err != nil {
		return internal.DBNotDeleted
	}
	return nil
}

func (p *PointRepository) GetById(id int32) (entity.TopicPoints, error) {
	var key entity.TopicPoints
	if err := p.db.Model(&key).Where("status != ", enum.UserDeletedStatus).Where("key_id=?", id).Scan(&key).Error; err != nil {
		return key, internal.DBNotFound
	}
	return key, nil
}

func (p *PointRepository) GetByUserId(id int32) (entity.TopicPoints, error) {
	var points entity.TopicPoints
	if err := p.db.Model(&points).Where("user_id=?", id).Scan(&points).Error; err != nil {
		return points, internal.DBNotFound
	}
	return points, nil
}

func (p *PointRepository) Update(role entity.TopicPoints) error {
	if err := p.db.Model(&role).Where("key_id=?", role.Id).Updates(entity.TopicPoints{
		Points: role.Points,
	}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
func (p *PointRepository) GetAllByUserIDProducts(id int32) ([]entity.TopicPoints, int64, error) {
	var pointList []entity.TopicPoints
	var total int64
	listQuery := p.db.Find(&pointList).Where("user_id = ?", id)

	if err := listQuery.Count(&total).Find(&pointList).Error; err != nil {
		return pointList, 0, err
	}
	return pointList, total, nil
}
