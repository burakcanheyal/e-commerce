package server

import (
	"attempt4/core/internal/application/handler"
	"attempt4/core/internal/middleware"
	"github.com/gin-gonic/gin"
)

type WebServer struct {
	productServerHandler     handler.ProductServerHandler
	userServerHandler        handler.UserServerHandler
	orderServerHandler       handler.OrderServerHandler
	authentication           handler.AuthenticationServerHandler
	authenticationMiddleware middleware.AuthenticationMiddleware
}

func NewWebServer(productServerHandler handler.ProductServerHandler, userServerHandler handler.UserServerHandler,
	orderServerHandler handler.OrderServerHandler, authentication handler.AuthenticationServerHandler,
	authenticationMiddleware middleware.AuthenticationMiddleware) WebServer {
	s := WebServer{productServerHandler, userServerHandler,
		orderServerHandler, authentication, authenticationMiddleware}
	return s
}
func (s *WebServer) SetupRoot() {
	router := gin.Default()

	router.POST("/login", s.authentication.Login)
	router.POST("/user/add", s.userServerHandler.Create)
	router.POST("/activation", s.userServerHandler.ActivateUser)

	user := router.Group("/profil", s.authenticationMiddleware.Auth())
	user.PUT("/", s.userServerHandler.Update)
	user.PUT("/pass/", s.userServerHandler.UpdatePassword)
	user.DELETE("/", s.userServerHandler.Delete)
	user.GET("/", s.userServerHandler.GetByUsername)

	order := router.Group("/order", s.authenticationMiddleware.Auth())
	order.GET("/:id", s.orderServerHandler.GetById)
	order.GET("/", s.orderServerHandler.GetAllOrders)
	order.POST("/", s.orderServerHandler.Create)
	order.PUT("/", s.orderServerHandler.Update)
	order.DELETE("/", s.orderServerHandler.Delete)

	product := router.Group("/product", s.authenticationMiddleware.Auth())
	product.GET("/:name", s.productServerHandler.GetByName)
	product.GET("/", s.productServerHandler.GetAllProducts)
	product.POST("/", s.productServerHandler.Create)
	product.DELETE("/", s.productServerHandler.Delete)
	product.PUT("/", s.productServerHandler.Update)

	router.Run("localhost:8000")
}
