package handler

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/service"
	"attempt4/platform/zap"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedbackServerHandler struct {
	feedbackService service.FeedbackService
}

func NewFeedbackServerHandler(feedbackService service.FeedbackService) FeedbackServerHandler {
	u := FeedbackServerHandler{feedbackService}
	return u
}
func (f *FeedbackServerHandler) GetFeedback(context *gin.Context) {
	code := dto.IdDto{}
	if err := context.BindJSON(&code); err != nil {
		zap.Logger.Error(internal.FailInTokenParse)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	feedback, err := f.feedbackService.GetFeedback(code.Id)
	if err != nil {
		context.JSON(http.StatusNotFound, NonExistItem())
		return
	}

	zap.Logger.Info("Feedbackleri görüntüleme başarılı")
	context.JSON(http.StatusOK, feedback)
}
func (f *FeedbackServerHandler) DeleteFeedback(context *gin.Context) {
	userDto, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}
	code := dto.IdDto{}
	if err := context.BindJSON(&code); err != nil {
		zap.Logger.Error(internal.FailInTokenParse)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := f.feedbackService.Deleteback(code.Id, userDto.Id)
	if err != nil {
		context.JSON(http.StatusNotFound, err.Error())
		return
	}

	zap.Logger.Info("Feedbackleri görüntüleme başarılı")
	context.JSON(http.StatusOK, SuccessInDelete())
}
func (f *FeedbackServerHandler) CreateFeedback(context *gin.Context) {
	userDto, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}
	feedback := dto.FeedbackDto{}
	if err := context.BindJSON(&feedback); err != nil {
		zap.Logger.Error(internal.FailInTokenParse)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}
	feedback.UserId = userDto.Id
	err := f.feedbackService.CreateFeedback(feedback)
	if err != nil {
		context.JSON(http.StatusNotFound, NonExistItem())
		return
	}

	zap.Logger.Info("Feedback yaratma başarılı")
	context.JSON(http.StatusOK, SuccessInCreate())
}
