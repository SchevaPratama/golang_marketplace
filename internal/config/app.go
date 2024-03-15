package config

import (
	"golang-marketplace/internal/handler"
	"golang-marketplace/internal/repository"
	"golang-marketplace/internal/router"
	"golang-marketplace/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	DB       *sqlx.DB
	App      *fiber.App
	Config   *viper.Viper
	Validate *validator.Validate
	Log      *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.DB)
	productRepository := repository.NewProductRepository(config.DB)
	bankAccountRepository := repository.NewBankAccountRepository(config.DB)

	// setup services
	productService := service.NewProductService(productRepository, config.Validate, config.Log, config.DB)
	userService := service.NewUserService(userRepository, config.Validate, config.Log)
	bankAccountService := service.NewBankAccountService(bankAccountRepository, config.Validate, config.Log)

	// setup handler
	productHandler := handler.NewProductHandler(productService, config.Log)
	userHandler := handler.NewUserHandler(userService, config.Log)
	bankAccountHandler := handler.NewBankAccountHandler(bankAccountService, config.Log)

	// recover from panic
	config.App.Use(func(c *fiber.Ctx) error {

		defer func() {
			if r := recover(); r != nil {
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}
		}()
		return c.Next()
	})

	// setup error handler

	// setup route
	routeConfig := router.RouteConfig{
		App:                config.App,
		ProductHandler:     productHandler,
		UserHandler:        userHandler,
		BankAccountHandler: bankAccountHandler,
	}

	routeConfig.Setup()
}
