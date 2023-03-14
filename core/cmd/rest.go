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
	"github.com/go-playground/locales/tr"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
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

	validate := validator.New()
	tr := tr.New()
	uni := ut.New(tr, tr)
	validationEntity := validation.NewValidation(validate, uni)
	validationEntity.ValidatorCustomMessages()

	db := postgres.InitializeDatabase(config.DBURL)

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	orderRepository := repository.NewOrderRepository(db)

	userService := service.NewUserService(userRepository, config.Secret)
	productService := service.NewProductService(productRepository)
	orderService := service.NewOrderService(orderRepository, productRepository, userRepository, config.Secret)
	authenticationService := service.NewAuthentication(userService, config.Secret, config.Secret2)

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(authenticationService)

	//Todo: Validation static olacak
	authenticationServerHandler := handler.NewAuthenticationServerHandler(authenticationService, validationEntity)
	userServerHandler := handler.NewUserServerHandler(userService, validationEntity)
	productServerHandler := handler.NewProductServerHandler(productService, validationEntity)
	orderServerHandler := handler.NewOrderServerHandler(orderService, validationEntity)

	webServer := server.NewWebServer(productServerHandler, userServerHandler, orderServerHandler, authenticationServerHandler, authenticationMiddleware)

	webServer.SetupRoot()
}
