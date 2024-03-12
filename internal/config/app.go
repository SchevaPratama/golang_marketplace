package config

import (
	"golang-marketplace/internal/handler"
	"golang-marketplace/internal/repository"
	"golang-marketplace/internal/router"
	"golang-marketplace/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sagikazarmark/slog-shim"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	DB       *sqlx.DB
	App      *fiber.App
	Config   *viper.Viper
	Validate *validator.Validate
	Log      *slog.Logger
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	productRepository := repository.NewProductRepository(config.DB)

	// setup services
	productService := service.NewProductService(productRepository, config.Validate, config.Log, config.DB)

	// setup handler
	productHandler := handler.NewProductHandler(productService, config.Log)

	// setup route
	routeConfig := router.RouteConfig{
		App:            config.App,
		ProductHandler: productHandler,
	}

	routeConfig.Setup()
}
