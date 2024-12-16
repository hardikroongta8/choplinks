package main

import (
	"log"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalln("Error while loading .env File:", err.Error())
	}
	log.Printf("Starting the server on port %s", cfg.Server.Port)
	store, err := NewMySQLStore()
	if err != nil {
		log.Fatalln("Error while initializing db store:", err.Error())
	}
	server := NewAPIServer(":"+cfg.Server.Port, store)
	err = store.Init()
	if err != nil {
		log.Fatalln("Error initializing tables:", err.Error())
	}
	server.Run()
}
