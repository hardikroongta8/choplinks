package main

import (
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
	db.Connect(cfg.Database.URI)
	defer db.Disconnect()

	database := db.GetDatabase(cfg.Database.DBName)

	err := http.ListenAndServe(":"+cfg.Server.Port, routes.SetupRoutes(database))
	if err != nil {
		log.Fatal("Some error occurred while listening to the server!")
	}
}
