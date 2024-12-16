package main

import (
	"github.com/joho/godotenv"
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

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	cfg.Server.BaseURL = os.Getenv("BASE_URL")
	cfg.Server.Port = os.Getenv("PORT")

	cfg.DB.URI = os.Getenv("DB_URI")

	appConfig = cfg
	return appConfig, nil
}

func GetConfig() *Config {
	return appConfig
}
