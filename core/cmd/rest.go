package cmd

import (
	"attempt4/core/internal/application/handler"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/service"
	"attempt4/core/internal/middleware"
	"attempt4/core/internal/server"
	"attempt4/core/platform/postgres"
	"attempt4/core/platform/postgres/repository"
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

	db := postgres.InitializeDatabase(config.DBURL)

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	keyRepository := repository.NewKeyRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	panelRepository := repository.NewAppOperationRepository(db)
	walletOperation := repository.NewWalletOperationRepository(db)

	userService := service.NewUserService(userRepository, keyRepository, walletRepository)
	productService := service.NewProductService(productRepository, userRepository)
	orderService := service.NewOrderService(orderRepository, productRepository, userRepository)
	authenticationService := service.NewAuthentication(userRepository, config.Secret, config.Secret2)
	walletService := service.NewWalletService(userRepository, walletRepository, productRepository, orderRepository, walletOperation)
	keyService := service.NewRolService(userRepository, keyRepository, panelRepository)

	authenticationMiddleware := middleware.NewMiddleware(authenticationService, userService)

	authenticationServerHandler := handler.NewAuthenticationServerHandler(authenticationService)
	profileServerHandler := handler.NewProfileServerHandler(userService)
	productServerHandler := handler.NewProductServerHandler(productService)
	orderServerHandler := handler.NewOrderServerHandler(orderService)
	walletServerHandler := handler.NewWalletServerHandler(walletService)
	keyServerHandler := handler.NewAppOperationServerHandler(keyService)

	webServer := server.NewWebServer(
		productServerHandler,
		profileServerHandler,
		orderServerHandler,
		authenticationServerHandler,
		walletServerHandler,
		keyServerHandler,
		authenticationMiddleware,
	)

	webServer.SetupRoot()
}
