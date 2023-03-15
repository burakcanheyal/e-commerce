package cmd

import (
	"attempt4/core/internal/application/handler"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/service"
	"attempt4/core/internal/middleware"
	"attempt4/core/internal/server"
	"attempt4/core/platform/postgres"
	"attempt4/core/platform/postgres/repository"
	"attempt4/core/platform/validation"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}
func Setup() {
	config := dto.Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Println(err)
	}

	validation.ValidatorCustomMessages()

	db := postgres.InitializeDatabase(config.DBURL)

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	keyRepository := repository.NewKeyRepository(db)

	userService := service.NewUserService(userRepository, keyRepository, config.Secret)
	productService := service.NewProductService(productRepository, userRepository, config.Secret)
	orderService := service.NewOrderService(orderRepository, productRepository, userRepository, config.Secret)
	authenticationService := service.NewAuthentication(userService, config.Secret, config.Secret2)

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(authenticationService)

	authenticationServerHandler := handler.NewAuthenticationServerHandler(authenticationService)
	profileServerHandler := handler.NewProfileServerHandler(userService)
	productServerHandler := handler.NewProductServerHandler(productService)
	orderServerHandler := handler.NewOrderServerHandler(orderService)

	webServer := server.NewWebServer(
		productServerHandler,
		profileServerHandler,
		orderServerHandler,
		authenticationServerHandler,
		authenticationMiddleware,
	)

	webServer.SetupRoot()
}
