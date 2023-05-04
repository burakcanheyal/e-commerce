package cmd

import (
	"attempt4/internal/application/handler"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/service"
	"attempt4/internal/middleware"
	"attempt4/internal/server"
	"attempt4/platform/app_log"
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
	db, err := postgres.InitializeDatabase(config.DBURL)
	if err != nil {
		log.Println(err)
	}

	err = app_log.InitializeAppLogDatabase(db)
	if err != nil {
		log.Println(err)
	}

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	panelRepository := repository.NewSubmissionRepository(db)
	walletOperationRepository := repository.NewWalletOperationRepository(db)
	applicationLogRepository := app_log.NewApplicationLogRepository(db)

	appLogService := app_log.NewApplicationLogService(applicationLogRepository)
	userService := service.NewUserService(userRepository, roleRepository, walletRepository, appLogService)
	productService := service.NewProductService(productRepository, userRepository, appLogService)
	orderService := service.NewOrderService(orderRepository, productRepository, userRepository, appLogService)
	authenticationService := service.NewAuthentication(userRepository, config.Secret, config.Secret2, appLogService)
	walletService := service.NewWalletService(userRepository, walletRepository, productRepository,
		orderRepository, walletOperationRepository, roleRepository, appLogService)
	keyService := service.NewRolService(userRepository, roleRepository, panelRepository, appLogService)

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
