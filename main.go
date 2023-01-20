package main

import (
	"log"

	"final-project-backend/db"
	"final-project-backend/server"
	"final-project-backend/utils/storage"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
	}

	err = storage.Connect()
	if err != nil {
		log.Println("Failed to connect storage", err)
	}

	server.Init()
}
