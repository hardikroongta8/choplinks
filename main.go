package main

import (
	"github.com/hardikroongta8/choplinks/api"
	"github.com/hardikroongta8/choplinks/config"
	"github.com/hardikroongta8/choplinks/storage"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("Error while loading .env File:", err.Error())
	}
	log.Printf("Starting the server on port %s", cfg.Port)
	store, err := storage.NewMySQLStore(cfg.DatabaseURI)
	if err != nil {
		log.Fatalln("Error while initializing db storage:", err.Error())
	}
	server := api.NewServer(":"+cfg.Port, store)
	err = store.Init()
	if err != nil {
		log.Fatalln("Error initializing tables:", err.Error())
	}
	server.Run()
}
