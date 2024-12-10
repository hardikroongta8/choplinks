package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Server struct {
		BaseURL string
		Port    string
	}
	DB struct {
		URI string
	}
}

var appConfig *Config = nil

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found...")
	}
	cfg := &Config{}
	cfg.Server.BaseURL = os.Getenv("BASE_URL")
	cfg.Server.Port = os.Getenv("PORT")

	cfg.DB.URI = os.Getenv("DB_URI")

	appConfig = cfg
	return appConfig
}

func GetConfig() *Config {
	return appConfig
}
