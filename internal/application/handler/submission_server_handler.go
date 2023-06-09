package handler

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/service"
	"attempt4/platform/zap"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SubmissionServerHandler struct {
	KeyService service.RolService
}

func NewSubmissionServerHandler(keyService service.RolService) SubmissionServerHandler {
	k := SubmissionServerHandler{keyService}
	return k
}

func (a *SubmissionServerHandler) UpdateUserRole(context *gin.Context) {
	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := a.KeyService.SubmissionUserRole(id.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("User rol değiştirme isteği başarılıyla eklendi")
	context.JSON(http.StatusOK, SuccessInSendRequest())
}

func (a *SubmissionServerHandler) ResponseToChangeUserRole(context *gin.Context) {
	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	response := dto.AppOperationDto{}
	if err := context.BindJSON(&response); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusServiceUnavailable, ErrorInJson())
		return
	}

	err := a.KeyService.ResultOfUpdateUserRole(response, id.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("User rol değiştirme isteğine cevap başarılı")
	context.JSON(http.StatusOK, SuccessInResponseRequest())
}
