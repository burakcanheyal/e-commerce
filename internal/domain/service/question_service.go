package service

import (
	"attempt4/platform/postgres/repository"
)

type QuestionService struct {
	userRepository     repository.UserRepository
	questionRepository repository.QuestionRepository
}

func NewQuestionService(
	userRepository repository.UserRepository,
	questionRepository repository.QuestionRepository) QuestionService {

	q := QuestionService{
		userRepository,
		questionRepository,
	}
	return q
}
