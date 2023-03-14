package handler

import (
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/service"
	"attempt4/core/platform/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenticationServerHandler struct {
	authenticationService service.Authentication
}

func NewAuthenticationServerHandler(authenticationService service.Authentication) AuthenticationServerHandler {
	a := AuthenticationServerHandler{authenticationService}
	return a
}
func (u *AuthenticationServerHandler) Login(context *gin.Context) {
	user := dto.AuthDto{}
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}
	err = u.authenticationService.Login(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	accessToken, err := u.authenticationService.GenerateAccessToken(user.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}
	refreshToken, err := u.authenticationService.GenerateRefreshToken(user.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, gin.H{"access": accessToken, "refresh": refreshToken})
}
