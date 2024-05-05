package service

import (
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/postgres/repository"
	"errors"
	"time"
)

type FeedbackService struct {
	userRepository     repository.UserRepository
	tripRepository     repository.TripRepository
	feedbackRepository repository.FeedbackRepository
}

func NewFeedbackService(
	userRepository repository.UserRepository,
	tripRepository repository.TripRepository,
	feedbackRepository repository.FeedbackRepository) FeedbackService {
	u := FeedbackService{
		userRepository,
		tripRepository,
		feedbackRepository,
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
			TripId:      feedbackT.TripId,
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
		TripId:      feedback.TripId,
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
func (f *FeedbackService) Deleteback(id int32) error {
	err := f.feedbackRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
