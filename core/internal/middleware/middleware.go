package middleware

import (
	"attempt4/core/internal"
	"attempt4/core/internal/application/handler"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	authenticationService service.Authentication
	userService           service.UserService
}

func NewMiddleware(authenticationService service.Authentication, userService service.UserService) Middleware {
	a := Middleware{authenticationService, userService}
	return a
}
func (a *Middleware) Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authentication")
		refreshToken := context.GetHeader("Refresh")
		if tokenString == "" {
			context.AbortWithStatusJSON(401, handler.TokenError())
			return
		}
		user, err := a.authenticationService.GetUserByTokenString(tokenString)
		if err != nil {
			context.AbortWithStatusJSON(401, handler.NewHttpError(err))
			return
		}

		if user.Status == enum.UserDeletedStatus {
			context.AbortWithStatusJSON(401, internal.DeletedUser)
			return
		}
		if user.Status == enum.UserPassiveStatus {
			context.AbortWithStatusJSON(401, internal.PassiveUser)
			return
		}

		err = a.authenticationService.ValidateAccessToken(tokenString)
		if err != nil {
			err := a.authenticationService.ValidateRefreshToken(refreshToken)
			if err != nil {
				context.AbortWithStatusJSON(401, handler.NewHttpError(err))
				return
			}

			tokenString, err := a.authenticationService.GenerateAccessToken(user.Username)
			if err != nil {
				context.AbortWithStatusJSON(401, handler.NewHttpError(err))
				return
			}
			context.JSON(200, tokenString)
		}

		context.Set("id", dto.IdDto{Id: user.Id})
		context.Next()
	}
}
func (a *Middleware) Permission(permissionType []int) gin.HandlerFunc {
	return func(context *gin.Context) {
		//Todo:Boş olma durumunu denetle, TokenDto, TokenUser

		id, exist := context.Keys["id"].(dto.IdDto)
		if exist != true {
			context.AbortWithStatusJSON(401, internal.UserNotFound)
			return
		}

		rol, err := a.userService.GetUserRoleById(id.Id)
		if err != nil {
			context.AbortWithStatusJSON(401, handler.NewHttpError(err))
			return
		}

		for i := range permissionType {
			if permissionType[i] == rol {
				context.Next()
				break
			}
		}
		context.AbortWithStatusJSON(401, internal.UserUnauthorized)
		return

	}
}
