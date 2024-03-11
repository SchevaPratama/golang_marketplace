package main

import (
	"fmt"
	"golang-marketplace/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)

	err := app.Listen(":8000")
	if err != nil {
		fmt.Errorf("Fatal error config file: %w \n", err)
	}
}
