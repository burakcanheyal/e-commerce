package server

import (
	handler2 "attempt4/internal/application/handler"
	"attempt4/internal/domain/enum"
	"attempt4/internal/middleware"
	"github.com/gin-gonic/gin"
)

type WebServer struct {
	productServerHandler handler2.ProductServerHandler
	profileServerHandler handler2.ProfileServerHandler
	orderServerHandler   handler2.OrderServerHandler
	authentication       handler2.AuthenticationServerHandler
	walletServerHandler  handler2.WalletServerHandler
	keyServerHandler     handler2.SubmissionServerHandler
	middleware           middleware.Middleware
}

func NewWebServer(
	productServerHandler handler2.ProductServerHandler,
	profileServerHandler handler2.ProfileServerHandler,
	orderServerHandler handler2.OrderServerHandler,
	authentication handler2.AuthenticationServerHandler,
	walletServerHandler handler2.WalletServerHandler,
	keyServerHandler handler2.SubmissionServerHandler,
	middleware middleware.Middleware,
) WebServer {
	s := WebServer{
		productServerHandler,
		profileServerHandler,
		orderServerHandler,
		authentication,
		walletServerHandler,
		keyServerHandler,
		middleware,
	}
	return s
}
func (s *WebServer) SetupRoot() {
	router := gin.Default()

	router.GET("", s.profileServerHandler.Test)
	router.POST("/login", s.authentication.Login)
	router.POST("/user/add", s.profileServerHandler.Create)
	router.POST("/activation", s.profileServerHandler.ActivateUser)

	user := router.Group("/profil", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser, enum.RoleManager, enum.RoleAdmin}))
	user.PUT("/", s.profileServerHandler.Update)
	user.PUT("/pass/", s.profileServerHandler.UpdatePassword)
	user.DELETE("/", s.profileServerHandler.Delete)
	user.GET("/", s.profileServerHandler.GetUser)

	changeUserRole := router.Group("/rol", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser}))
	changeUserRole.GET("/", s.keyServerHandler.UpdateUserRole)

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
	product.GET("/statistics", s.walletServerHandler.GetAllSellTransactions)

	wallet := router.Group("/wallet", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser, enum.RoleManager, enum.RoleAdmin}))
	wallet.PUT("/", s.walletServerHandler.Update)
	wallet.GET("/complete", s.walletServerHandler.CompletePurchase)
	wallet.GET("/", s.walletServerHandler.GetAllBuyTransactions)

	panel := router.Group("/panel", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleAdmin}))
	panel.POST("/", s.keyServerHandler.ResponseToChangeUserRole)

	router.Run("0.0.0.0:8001")
}
