package server

import (
	"attempt4/core/internal/application/handler"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/internal/middleware"
	"github.com/gin-gonic/gin"
)

type WebServer struct {
	productServerHandler handler.ProductServerHandler
	profileServerHandler handler.ProfileServerHandler
	orderServerHandler   handler.OrderServerHandler
	authentication       handler.AuthenticationServerHandler
	middleware           middleware.Middleware
}

func NewWebServer(
	productServerHandler handler.ProductServerHandler,
	profileServerHandler handler.ProfileServerHandler,
	orderServerHandler handler.OrderServerHandler,
	authentication handler.AuthenticationServerHandler,
	middleware middleware.Middleware,
) WebServer {
	s := WebServer{
		productServerHandler,
		profileServerHandler,
		orderServerHandler,
		authentication,
		middleware,
	}
	return s
}
func (s *WebServer) SetupRoot() {
	router := gin.Default()

	router.POST("/login", s.authentication.Login)
	router.POST("/user/add", s.profileServerHandler.Create)
	router.POST("/activation", s.profileServerHandler.ActivateUser)

	user := router.Group("/profil", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser, enum.RoleManager, enum.RoleAdmin}))
	user.PUT("/", s.profileServerHandler.Update)
	user.PUT("/pass/", s.profileServerHandler.UpdatePassword)
	user.DELETE("/", s.profileServerHandler.Delete)
	user.GET("/", s.profileServerHandler.GetByUsername)

	order := router.Group("/order", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser, enum.RoleManager, enum.RoleAdmin}))
	order.GET("/:id", s.orderServerHandler.GetById)
	order.GET("/", s.orderServerHandler.GetAllOrders)
	order.POST("/", s.orderServerHandler.Create)
	order.PUT("/", s.orderServerHandler.Update)
	order.DELETE("/", s.orderServerHandler.Delete)

	product := router.Group("/product", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleManager}))
	product.GET("/:name", s.productServerHandler.GetByName)
	product.GET("/", s.productServerHandler.GetAllProducts)
	product.POST("/", s.productServerHandler.Create)
	product.DELETE("/", s.productServerHandler.Delete)
	product.PUT("/", s.productServerHandler.Update)

	router.Run("localhost:8000")
}
