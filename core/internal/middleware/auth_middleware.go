package middleware

import (
	"attempt4/core/internal"
	"attempt4/core/internal/application/handler"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type AuthenticationMiddleware struct {
	authenticationService service.Authentication
}

func NewAuthenticationMiddleware(authenticationService service.Authentication) AuthenticationMiddleware {
	a := AuthenticationMiddleware{authenticationService}
	return a
}
func (a *AuthenticationMiddleware) Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authentication")
		refreshToken := context.GetHeader("Refresh")
		if tokenString == "" {
			context.JSON(401, handler.TokenError())
			context.Abort()
			return
		}
		user, err := a.authenticationService.UserService.GetUserByTokenString(tokenString)
		if err != nil {
			context.JSON(401, handler.NewHttpError(err))
			context.Abort()
			return
		}

		if user.Status == enum.UserDeletedStatus {
			context.JSON(401, handler.NewHttpError(err))
			context.Abort()
			return
		}
		if user.Status == enum.UserPassiveStatus {
			context.JSON(401, handler.NewHttpError(err))
			context.Abort()
			return
		}

		err = a.authenticationService.ValidateAccessToken(tokenString)
		if err != nil {
			err := a.authenticationService.ValidateRefreshToken(refreshToken)
			if err != nil {
				context.JSON(401, handler.NewHttpError(err))
				context.Abort()
				return
			}

			tokenString, err := a.authenticationService.GenerateAccessToken(user.Username)
			if err != nil {
				context.JSON(401, handler.NewHttpError(err))
				context.Abort()
				return
			}
			context.JSON(200, tokenString)
		}

		context.Next()
	}
}
func (a *AuthenticationMiddleware) Permission(permissionType int) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authentication")
		if tokenString == "" {
			context.JSON(401, handler.TokenError())
			context.Abort()
			return
		}

		rol, err := a.authenticationService.UserService.GetUserRoleByTokenString(tokenString)
		if err != nil {
			context.JSON(401, handler.NewHttpError(err))
			context.Abort()
			return
		}

		if permissionType != 0 {
			if rol != permissionType {
				context.AbortWithStatusJSON(401, internal.UserUnauthorized)
				return
			}
		}

		context.Next()
	}
}
