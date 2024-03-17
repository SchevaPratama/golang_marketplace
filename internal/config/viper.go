package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func NewViper() *viper.Viper {
	config := viper.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")
	err = config.ReadInConfig()

	if err != nil {
		log.Fatalf("Error read config: %v", err)
	}

	return config
}
