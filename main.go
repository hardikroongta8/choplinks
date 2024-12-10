package main

import (
	"github.com/hardikroongta8/choplinks/internal/model"
	"github.com/hardikroongta8/choplinks/internal/routes"
	"github.com/hardikroongta8/choplinks/pkg/config"
	"github.com/hardikroongta8/choplinks/pkg/db"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	log.Printf(
		"Starting the server on port %s in %s mode",
		cfg.Server.Port,
		cfg.Server.Environment,
	)
	database := db.Connect(cfg.DB.URI)
	err := database.AutoMigrate(&model.UrlMap{})
	if err != nil {
		log.Fatal("Error while migrating the DB:", err.Error())
	}
	err = http.ListenAndServe(":"+cfg.Server.Port, routes.SetupRoutes(database))
	if err != nil {
		log.Fatal("Some error occurred while listening to the server!")
	}
}
