package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	BaseURL     string
	Port        string
	DatabaseURI string
	JWTSecret   string
}

var appConfig *Config = nil

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	appConfig = &Config{
		BaseURL:     os.Getenv("BASE_URL"),
		Port:        os.Getenv("PORT"),
		DatabaseURI: os.Getenv("DB_URI"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
	return appConfig, nil
}

func Get() *Config {
	return appConfig
}
