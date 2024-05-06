package handler

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/service"
	"attempt4/platform/zap"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TripServerHandler struct {
	questionService service.QuestionService
}

func NewTripServerHandler(questionService service.QuestionService) TripServerHandler {
	a := TripServerHandler{questionService}
	return a
}

func (t *TripServerHandler) GetQuestions(context *gin.Context) {
	_, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(http.StatusBadRequest, internal.UserNotFound)
		return
	}

	questions, err := t.questionService.GetQuestions()
	if err != nil {
		context.JSON(http.StatusNotFound, NonExistItem())
		return
	}

	zap.Logger.Info("Soruları görüntüleme başarılı")
	context.JSON(http.StatusOK, questions)
}
func (t *TripServerHandler) AnswerQuestion(context *gin.Context) {
	user, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(http.StatusBadRequest, internal.UserNotFound)
		return
	}
	var questions []dto.TopicPointsDto
	if err := context.BindJSON(&questions); err != nil {
		zap.Logger.Error(internal.FailInTokenParse)

		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}
	err := t.questionService.CalculatePoints(user.Id, questions)
	if err != nil {
		context.JSON(http.StatusInternalServerError, AnswerQuestion())
		return
	}
	zap.Logger.Info("Soruları cevaplama başarılı")
	context.JSON(http.StatusOK, SuccessInCreate())
}
func (t *TripServerHandler) GetAIRecommendation(context *gin.Context) {
	user, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(http.StatusBadRequest, internal.UserNotFound)
		return
	}

	trips, err := t.questionService.GetAIRecommendation(user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, AnswerQuestion())
	}
	if trips.Name == "" {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(http.StatusBadRequest, internal.UserNotFound)
		return
	}
	zap.Logger.Info("AI Recommendation başarılı")
	context.JSON(http.StatusOK, trips)
}
