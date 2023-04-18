package cmd

import (
	"attempt4/internal/application/handler"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/service"
	"attempt4/internal/middleware"
	"attempt4/internal/server"
	"attempt4/platform/postgres"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/zap"
	"github.com/spf13/viper"
	"log"
)

func init() {
	zap.Logger.Info("Env dosyaları okunuyor")

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}
func Setup() {
	zap.Logger.Info("Setup Başlatıldı")
	config := dto.Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Println(err)
	}
	db := postgres.InitializeDatabase(config.DBURL)

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	panelRepository := repository.NewSubmissionRepository(db)
	walletOperation := repository.NewWalletOperationRepository(db)

	userService := service.NewUserService(userRepository, roleRepository, walletRepository)
	productService := service.NewProductService(productRepository, userRepository)
	orderService := service.NewOrderService(orderRepository, productRepository, userRepository)
	authenticationService := service.NewAuthentication(userRepository, config.Secret, config.Secret2)
	walletService := service.NewWalletService(userRepository, walletRepository, productRepository,
		orderRepository, walletOperation, roleRepository)
	keyService := service.NewRolService(userRepository, roleRepository, panelRepository)

	authenticationMiddleware := middleware.NewMiddleware(authenticationService, userService)

	authenticationServerHandler := handler.NewAuthenticationServerHandler(authenticationService)
	profileServerHandler := handler.NewProfileServerHandler(userService)
	productServerHandler := handler.NewProductServerHandler(productService)
	orderServerHandler := handler.NewOrderServerHandler(orderService)
	walletServerHandler := handler.NewWalletServerHandler(walletService)
	keyServerHandler := handler.NewSubmissionServerHandler(keyService)

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
