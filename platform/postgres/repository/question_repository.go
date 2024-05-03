package repository

import (
	"attempt4/internal"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	r := QuestionRepository{db}
	return r
}

func (q *QuestionRepository) Create(question entity.Question) (entity.Question, error) {
	if err := q.db.Create(&question).Error; err != nil {
		return question, internal.DBNotCreated
	}
	return question, nil
}

func (q *QuestionRepository) Delete(question entity.Question) error {
	if err := q.db.Model(&question).Where("id=?", question.Id).Update("status", enum.QuestionDeleted).Error; err != nil {
		return internal.DBNotDeleted
	}
	return nil
}

func (q *QuestionRepository) GetById(id int32) (entity.Question, error) {
	var key entity.Question
	if err := q.db.Model(&key).Where("status != ", enum.QuestionDeleted).Where("id=?", id).Scan(&key).Error; err != nil {
		return key, internal.DBNotFound
	}
	return key, nil
}

func (q *QuestionRepository) Update(question entity.Question) error {
	if err := q.db.Model(&question).Where("id=?", question.Id).Updates(entity.Question{
		Description: question.Description,
	}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
func (q *QuestionRepository) GetAllProducts() ([]entity.Question, int64, error) {
	var productList []entity.Question
	var total int64
	listQuery := q.db.Find(&productList).Where("status != ?", enum.QuestionDeleted)

	if err := listQuery.Count(&total).Find(&productList).Error; err != nil {
		return productList, 0, err
	}
	return productList, total, nil
}
