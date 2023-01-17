package main

import (
	"log"

	"final-project-backend/db"
	"final-project-backend/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
	}
	server.Init()
}
