package handler

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/service"
	"attempt4/platform/validation"
	"attempt4/platform/zap"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileServerHandler struct {
	UserService service.UserService
}

func NewProfileServerHandler(userService service.UserService) ProfileServerHandler {
	u := ProfileServerHandler{userService}
	return u
}

func (p *ProfileServerHandler) Create(context *gin.Context) {
	user := dto.UserDto{}
	if err := context.BindJSON(&user); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = p.UserService.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	zap.Logger.Info("Kullanıcı oluşturma başarılı")
	context.JSON(http.StatusOK, SuccessInCreate())
}

func (p *ProfileServerHandler) Update(context *gin.Context) {
	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	user := dto.UserDto{}
	if err := context.BindJSON(&user); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = p.UserService.UpdateUser(id.Id, user)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Kullanıcı güncelleme başarılı")
	context.JSON(http.StatusOK, SuccessInUpdate())
}

func (p *ProfileServerHandler) Delete(context *gin.Context) {
	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := p.UserService.DeleteUser(id.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Kullanıcı silme başarılı")
	context.JSON(http.StatusOK, SuccessInDelete())
}

func (p *ProfileServerHandler) GetUser(context *gin.Context) {
	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(http.StatusBadRequest, internal.UserNotFound)
		return
	}

	user, err := p.UserService.GetUserById(id.Id)
	if err != nil {
		context.JSON(http.StatusNotFound, NonExistItem())
		return
	}

	zap.Logger.Info("Kullanıcı bilgileri görüntüleme başarılı")
	context.JSON(http.StatusOK, user)
}

func (p *ProfileServerHandler) UpdatePassword(context *gin.Context) {
	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	user := dto.UserUpdatePasswordDto{}
	if err := context.BindJSON(&user); err != nil {
		zap.Logger.Error(internal.FailInTokenParse)

		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = p.UserService.UpdateUserPassword(id.Id, user)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}
	zap.Logger.Info("Kullanıcı şifre değiştirme başarılı")
	context.JSON(http.StatusOK, SuccessInUpdate())
}

func (p *ProfileServerHandler) ActivateUser(context *gin.Context) {
	code := dto.UserUpdateCodeDto{}
	if err := context.BindJSON(&code); err != nil {
		zap.Logger.Error(internal.FailInTokenParse)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(code)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = p.UserService.ActivateUser(code)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Kullanıcı aktive etme başarılı")
	context.JSON(http.StatusOK, SuccessInActivate())
}

func (p *ProfileServerHandler) Test(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
