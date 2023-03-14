package handler

import (
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/service"
	"attempt4/core/platform/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserServerHandler struct {
	userService service.UserService
	validation  validation.Validation
}

func NewUserServerHandler(userService service.UserService, validation validation.Validation) UserServerHandler {
	u := UserServerHandler{userService, validation}
	return u
}
func (u *UserServerHandler) Create(context *gin.Context) {
	user := dto.UserDto{}
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := u.validation.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = u.userService.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, UserExist())
		return
	}

	context.JSON(http.StatusOK, SuccessInCreate())
}
func (u *UserServerHandler) Update(context *gin.Context) {
	user := dto.UserDto{}
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := u.validation.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = u.userService.UpdateUser(user)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInUpdate())
}
func (u *UserServerHandler) Delete(context *gin.Context) {
	tokenString := context.GetHeader("Authentication")

	if tokenString == "" {
		context.JSON(401, TokenError())
		return
	}
	err := u.userService.DeleteUser(tokenString)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInDelete())
}
func (u *UserServerHandler) GetByUsername(context *gin.Context) {
	tokenString := context.GetHeader("Authentication")

	if tokenString == "" {
		context.JSON(401, TokenError())
		return
	}

	user, err := u.userService.GetUserByTokenString(tokenString)
	if err != nil {
		context.JSON(http.StatusNotFound, NonExistItem())
		return
	}

	context.JSON(http.StatusOK, user)
}
func (u *UserServerHandler) UpdatePassword(context *gin.Context) {
	user := dto.UserUpdatePasswordDto{}
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := u.validation.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = u.userService.UpdateUserPassword(user)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInUpdate())
}
func (u *UserServerHandler) ActivateUser(context *gin.Context) {
	code := dto.UserUpdateCodeDto{}
	if err := context.BindJSON(&code); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := u.validation.ValidateStruct(code)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = u.userService.ActivateUser(code)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInActivate())
}
