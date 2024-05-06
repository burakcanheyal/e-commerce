package server

import (
	handler2 "attempt4/internal/application/handler"
	"attempt4/internal/domain/enum"
	"attempt4/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type WebServer struct {
	productServerHandler  handler2.ProductServerHandler
	profileServerHandler  handler2.ProfileServerHandler
	orderServerHandler    handler2.OrderServerHandler
	authentication        handler2.AuthenticationServerHandler
	walletServerHandler   handler2.WalletServerHandler
	keyServerHandler      handler2.SubmissionServerHandler
	middleware            middleware.Middleware
	tripServerHandler     handler2.TripServerHandler
	adminPanelHandler     handler2.AdminPanelHandler
	feedbackServerHandler handler2.FeedbackServerHandler
}

func NewWebServer(
	productServerHandler handler2.ProductServerHandler,
	profileServerHandler handler2.ProfileServerHandler,
	orderServerHandler handler2.OrderServerHandler,
	authentication handler2.AuthenticationServerHandler,
	walletServerHandler handler2.WalletServerHandler,
	keyServerHandler handler2.SubmissionServerHandler,
	middleware middleware.Middleware,
	tripServerHandler handler2.TripServerHandler,
	adminPanelHandler handler2.AdminPanelHandler,
	feedbackServerHandler handler2.FeedbackServerHandler,
) WebServer {
	s := WebServer{
		productServerHandler,
		profileServerHandler,
		orderServerHandler,
		authentication,
		walletServerHandler,
		keyServerHandler,
		middleware,
		tripServerHandler,
		adminPanelHandler,
		feedbackServerHandler,
	}
	return s
}
func (s *WebServer) SetupRoot() {
	router := gin.Default()
	cors.Default()
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authentication"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	corsConfig.AllowAllOrigins = true
	c := cors.New(corsConfig)
	router.Use(c)
	router.POST("/login", s.authentication.Login)
	router.POST("/user/add", s.profileServerHandler.Create)
	router.POST("/activation", s.profileServerHandler.ActivateUser)

	user := router.Group("/profil/", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser, enum.RoleManager, enum.RoleAdmin}))
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
	//product.GET("/statistics", s.walletServerHandler.GetAllSellTransactions)

	wallet := router.Group("/wallet", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser, enum.RoleManager, enum.RoleAdmin}))
	wallet.PUT("/", s.walletServerHandler.Update)
	wallet.GET("/complete", s.walletServerHandler.CompletePurchase)
	wallet.GET("/", s.walletServerHandler.Get)
	wallet.GET("/completedOrders", s.walletServerHandler.GetAllBuyTransactions)

	trip := router.Group("/trips", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser, enum.RoleManager, enum.RoleAdmin}))
	trip.GET("/", s.tripServerHandler.GetQuestions)
	trip.POST("/", s.tripServerHandler.AnswerQuestion)
	trip.GET("/recommendation/", s.tripServerHandler.GetAIRecommendation)

	feedback := router.Group("/feedback", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleUser, enum.RoleManager, enum.RoleAdmin}))
	feedback.POST("/", s.feedbackServerHandler.GetFeedback)
	feedback.DELETE("/", s.feedbackServerHandler.DeleteFeedback)
	feedback.POST("/add/", s.feedbackServerHandler.CreateFeedback)

	panel := router.Group("/panel", s.middleware.Auth(), s.middleware.Permission([]int{enum.RoleAdmin}))
	panel.POST("/", s.keyServerHandler.ResponseToChangeUserRole)
	panel.GET("/user/", s.adminPanelHandler.GetUser)
	panel.GET("/trip/", s.adminPanelHandler.GetTrip)
	panel.DELETE("/user/", s.adminPanelHandler.DeleteUser)
	panel.DELETE("/trip/", s.adminPanelHandler.DeleteTrip)
	panel.POST("/trip/", s.adminPanelHandler.CreateTrip)

	router.Run("0.0.0.0:8001")
}
