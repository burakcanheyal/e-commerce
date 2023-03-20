package handler

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/service"
	"attempt4/core/platform/validation"
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

	context.JSON(http.StatusOK, SuccessInCreate())
}

func (p *ProfileServerHandler) Update(context *gin.Context) {
	user := dto.UserDto{}
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = p.UserService.UpdateUser(user)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInUpdate())
}

func (p *ProfileServerHandler) Delete(context *gin.Context) {
	id, exist := context.Get("id")
	if exist != true {
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := p.UserService.DeleteUser(id.(int32))
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInDelete())
}

func (p *ProfileServerHandler) GetByUsername(context *gin.Context) {
	id, exist := context.Get("id")
	if exist != true {
		context.JSON(http.StatusBadRequest, internal.UserNotFound)
		return
	}

	user, err := p.UserService.GetUserById(id.(int32))
	if err != nil {
		context.JSON(http.StatusNotFound, NonExistItem())
		return
	}

	context.JSON(http.StatusOK, user)
}

func (p *ProfileServerHandler) UpdatePassword(context *gin.Context) {
	user := dto.UserUpdatePasswordDto{}
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = p.UserService.UpdateUserPassword(user)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInUpdate())
}

func (p *ProfileServerHandler) ActivateUser(context *gin.Context) {
	code := dto.UserUpdateCodeDto{}
	if err := context.BindJSON(&code); err != nil {
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

	context.JSON(http.StatusOK, SuccessInActivate())
}
