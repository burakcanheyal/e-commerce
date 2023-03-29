package repository

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/entity"
	"gorm.io/gorm"
)

type SubmissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	p := SubmissionRepository{db}
	return p
}
func (p *SubmissionRepository) Create(panel entity.Submission) (entity.Submission, error) {
	if err := p.db.Create(&panel).Error; err != nil {
		return panel, internal.DBNotCreated
	}
	return panel, nil
}

func (p *SubmissionRepository) GetById(id int32) (entity.Submission, error) {
	var submission entity.Submission
	if err := p.db.Model(&submission).Where("id=?", id).First(&submission).Error; err != nil {
		return submission, internal.DBNotFound
	}
	return submission, nil
}

func (p *SubmissionRepository) GetByUserId(id int32) (entity.Submission, error) {
	var submission entity.Submission
	if err := p.db.Model(&submission).Where("applied_user_id=?", id).First(&submission).Error; err != nil {
		return submission, internal.DBNotFound
	}
	return submission, nil
}

func (p *SubmissionRepository) Update(submission entity.Submission) error {
	if err := p.db.Model(&submission).Where("id=?", submission.Id).Updates(
		entity.Submission{
			Status:              submission.Status,
			OperationResultDate: submission.OperationResultDate,
		}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}
