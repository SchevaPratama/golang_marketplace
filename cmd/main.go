package main

import (
	"fmt"
	"golang-marketplace/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	db := config.NewDatabase(viperConfig)
	aws := config.NewAws(viperConfig)
	validate := config.NewValidator(viperConfig)
	log := config.NewLogger(viperConfig)
	config.Bootstrap(&config.BootstrapConfig{
		App:      app,
		DB:       db,
		Config:   viperConfig,
		Validate: validate,
		Log:      log,
		Aws:      aws,
	})
	//webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatal("Failed to start server: %w \n", err)
	}
}
