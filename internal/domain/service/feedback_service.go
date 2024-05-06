package service

import (
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/postgres/repository"
	"errors"
	"log"
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
	feedbacks, total, error := f.feedbackRepository.GetByProductId(productID)
	if total == 0 {
		return feedback, errors.New("error getting feedback")
	}
	if error != nil {
		return feedback, errors.New("error getting feedback")
	}

	for _, feedbackT := range feedbacks {
		username, err := f.userRepository.GetById(feedbackT.UserId)
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
	feedbackE := entity.Feedback{
		Id:          0,
		Description: feedback.Description,
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

	role, err := f.roleRepository.GetByUserId(userId)
	if role.Id == 0 {
		return errors.New("Role cannot be found")
	}
	if err != nil {
		return err
	}

	log.Println(id)
	log.Println(feedback.UserId)
	log.Println(userId)

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
