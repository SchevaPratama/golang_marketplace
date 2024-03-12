package main

import (
	"fmt"
	"golang-marketplace/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	db := config.NewDatabase(viperConfig)
	validate := config.NewValidator(viperConfig)
	config.Bootstrap(&config.BootstrapConfig{
		App:      app,
		DB:       db,
		Config:   viperConfig,
		Validate: validate,
	})
	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		fmt.Errorf("Failed to start server: %w \n", err)
	}
}
