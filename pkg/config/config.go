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
	DB struct {
		username string
		password string
		dbName   string
		URI      string
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
	cfg.DB.username = os.Getenv("DB_USERNAME")
	cfg.DB.password = os.Getenv("DB_PASSWORD")
	cfg.DB.dbName = os.Getenv("DB_NAME")

	cfg.DB.URI = cfg.DB.username + ":" +
		cfg.DB.password + "@tcp(127.0.0.1:3306)/" +
		cfg.DB.dbName + "?charset=utf8&parseTime=True&loc=Local"

	return cfg
}
