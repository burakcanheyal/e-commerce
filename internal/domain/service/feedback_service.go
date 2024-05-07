package service

import (
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/regexp"
	"errors"
	"time"
)

type FeedbackService struct {
	userRepository     repository.UserRepository
	tripRepository     repository.TripRepository
	feedbackRepository repository.FeedbackRepository
	roleRepository     repository.RoleRepository
}

func NewFeedbackService(
	userRepository repository.UserRepository,
	tripRepository repository.TripRepository,
	feedbackRepository repository.FeedbackRepository,
	roleRepository repository.RoleRepository) FeedbackService {
	u := FeedbackService{
		userRepository,
		tripRepository,
		feedbackRepository,
		roleRepository,
	}
	return u
}
func (f *FeedbackService) GetFeedback(productID int32) ([]dto.FeedbackDto, error) {
	var feedback []dto.FeedbackDto
	feedbacks, total, err := f.feedbackRepository.GetByProductId(productID)
	if total == 0 {
		return feedback, errors.New("error getting feedback")
	}
	if err != nil {
		return feedback, errors.New("error getting feedback")
	}

	for _, feedbackT := range feedbacks {
		username, err := f.userRepository.GetByIdFeedbackService(feedbackT.UserId)
		if username.Id == 0 {
			return feedback, errors.New("error getting user")
		}
		if err != nil {
			return feedback, err
		}

		feedback = append(feedback, dto.FeedbackDto{
			Id:          feedbackT.Id,
			Description: feedbackT.Description,
			Star:        feedbackT.Star,
			ProductId:   feedbackT.ProductId,
			UserId:      feedbackT.UserId,
			Status:      feedbackT.Status,
			UserName:    username.Username,
		})
	}
	return feedback, nil
}
func (f *FeedbackService) CreateFeedback(feedback dto.FeedbackDto) error {
	controlledDescription := regexp.InspectFeedback(feedback.Description)
	feedbackE := entity.Feedback{
		Id:          0,
		Description: controlledDescription,
		Star:        feedback.Star,
		ProductId:   feedback.ProductId,
		UserId:      feedback.UserId,
		Status:      enum.TripActive,
		CreatedAt:   time.Now(),
		DeletedAt:   nil,
		UpdatedAt:   nil,
	}
	feedbackTest, err := f.feedbackRepository.Create(feedbackE)
	if feedbackTest.Id == 0 {
		return errors.New("error creating feedback")
	}
	if err != nil {
		return errors.New("error creating feedback")
	}

	return nil
}
func (f *FeedbackService) Deleteback(id int32, userId int32) error {
	feedback, err := f.feedbackRepository.GetById(id)
	if feedback.Id == 0 {
		return errors.New("Feedback cannot be found")
	}
	if err != nil {
		return err
	}

	role, err := f.roleRepository.GetByIdFeedbackService(userId)
	if role.Id == 0 {
		return errors.New("Role cannot be found")
	}
	if err != nil {
		return err
	}

	if userId != feedback.UserId {
		if role.Rol != enum.RoleAdmin {
			return errors.New("Bu kaydı silmeye yetkili değilsiniz.")
		}
	}
	err = f.feedbackRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
func (f *FeedbackService) GetAllFeedbacks() ([]dto.FeedbackDto, error) {
	var feedback []dto.FeedbackDto
	feedbacks, total, err := f.feedbackRepository.GetAllFeedbacks()
	if total == 0 {
		return feedback, errors.New("error getting feedback")
	}
	if err != nil {
		return feedback, errors.New("error getting feedback")
	}

	for _, feedbackT := range feedbacks {
		username, err := f.userRepository.GetByIdFeedbackService(feedbackT.UserId)
		if username.Id == 0 {
			return feedback, errors.New("error getting user")
		}
		if err != nil {
			return feedback, err
		}

		feedback = append(feedback, dto.FeedbackDto{
			Id:          feedbackT.Id,
			Description: feedbackT.Description,
			Star:        feedbackT.Star,
			ProductId:   feedbackT.ProductId,
			UserId:      feedbackT.UserId,
			Status:      feedbackT.Status,
			UserName:    username.Username,
		})
	}
	return feedback, nil
}
