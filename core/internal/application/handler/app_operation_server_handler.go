package handler

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppOperationServerHandler struct {
	KeyService service.RolService
}

func NewAppOperationServerHandler(keyService service.RolService) AppOperationServerHandler {
	k := AppOperationServerHandler{keyService}
	return k
}
func (a *AppOperationServerHandler) UpdateUserRole(context *gin.Context) {
	id, exist := context.Keys["id"].(dto.IdDto)
	if exist != true {
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := a.KeyService.AppOperationToUpdateUserRole(id.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInSendRequest())
}
func (a *AppOperationServerHandler) ResponseToChangeUserRole(context *gin.Context) {
	id, exist := context.Keys["id"].(dto.IdDto)
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
