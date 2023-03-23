package handler

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type KeyServerHandler struct {
	KeyService service.KeyService
}

func NewKeyServerHandler(keyService service.KeyService) KeyServerHandler {
	k := KeyServerHandler{keyService}
	return k
}
func (k *KeyServerHandler) UpdateUserRole(context *gin.Context) {
	id, exist := context.Keys["id"].(dto.IdDto)
	if exist != true {
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := k.KeyService.SendRequestToUpdateUserRole(id.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInSendRequest())
}
func (k *KeyServerHandler) ResponseToChangeUserRole(context *gin.Context) {
	response := dto.PanelDto{}
	if err := context.BindJSON(&response); err != nil {
		context.JSON(http.StatusServiceUnavailable, ErrorInJson())
		return
	}

	err := k.KeyService.ResponseToUpdateUserRole(response)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInResponseRequest())
}
