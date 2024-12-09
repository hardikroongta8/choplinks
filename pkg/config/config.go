package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port        string
		Environment string
		JWTSecret   string
	}
	Database struct {
		URI    string
		DBName string
	}
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found...")
	}
	cfg := &Config{}
	cfg.Server.Port = os.Getenv("PORT")
	cfg.Server.Environment = os.Getenv("ENVIRONMENT")
	cfg.Server.JWTSecret = os.Getenv("JWT_SECRET")
	cfg.Database.URI = os.Getenv("DB_URI")
	cfg.Database.DBName = os.Getenv("DB_NAME")
	return cfg
}
