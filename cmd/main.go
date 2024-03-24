package main

import (
	"fmt"
	"github.com/ansrivas/fiberprometheus/v2"
	"golang-marketplace/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	prometheus := fiberprometheus.New("golang-marketplace")
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
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)
	err := app.Listen(fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatal("Failed to start server: %w \n", err)
	}
}
