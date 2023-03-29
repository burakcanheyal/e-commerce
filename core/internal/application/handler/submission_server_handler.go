package handler

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/service"
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
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := a.KeyService.SubmissionUserRole(id.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInSendRequest())
}

func (a *SubmissionServerHandler) ResponseToChangeUserRole(context *gin.Context) {
	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		context.JSON(401, internal.UserNotFound)
		return
	}

	response := dto.AppOperationDto{}
	if err := context.BindJSON(&response); err != nil {
		context.JSON(http.StatusServiceUnavailable, ErrorInJson())
		return
	}

	err := a.KeyService.ResultOfUpdateUserRole(response, id.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInResponseRequest())
}
