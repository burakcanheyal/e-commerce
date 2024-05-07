package repository

import (
	"attempt4/internal"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"gorm.io/gorm"
	"time"
)

type FeedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	p := FeedbackRepository{db}
	return p
}

func (f *FeedbackRepository) Create(feedback entity.Feedback) (entity.Feedback, error) {
	if err := f.db.Create(&feedback).Error; err != nil {
		return feedback, internal.DBNotCreated
	}

	return feedback, nil
}

func (f *FeedbackRepository) Delete(id int32) error {
	var feedback entity.Feedback
	if err := f.db.Model(&feedback).Where("id = ?", id).Update("status", enum.ProductDeleted).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return internal.DBNotFound
		}
		return internal.DBNotDeleted
	}
	time := time.Now()
	if err := f.db.Model(&feedback).Where("id = ?", id).Update("deleted_at", time).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return internal.DBNotFound
		}
		return internal.DBNotDeleted
	}

	return nil
}

func (f *FeedbackRepository) GetByProductId(id int32) ([]entity.Feedback, int64, error) {
	var productList []entity.Feedback
	var total int64
	listQuery := f.db.Find(&productList).Where("status != ?", enum.ProductDeleted).Where("product_id = ?", id)

	if err := listQuery.Count(&total).Find(&productList).Error; err != nil {
		return productList, 0, err
	}
	return productList, total, nil
}
func (f *FeedbackRepository) GetById(id int32) (entity.Feedback, error) {
	var productList entity.Feedback
	if err := f.db.Model(&productList).Where("status != ?", enum.ProductDeleted).Where("id=?", id).Scan(&productList).Error; err != nil {
		return productList, internal.DBNotFound
	}
	return productList, nil
}
func (f *FeedbackRepository) GetAllFeedbacks() ([]entity.Feedback, int64, error) {
	var productList []entity.Feedback
	var total int64
	listQuery := f.db.Find(&productList).Where("status != ?", enum.ProductDeleted)

	if err := listQuery.Count(&total).Find(&productList).Error; err != nil {
		return productList, 0, err
	}
	return productList, total, nil
}
