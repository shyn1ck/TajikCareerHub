package main

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"log"
)

func main() {
	err := logger.Init()
	if err != nil {
		return
	}
	if err := db.ConnectToDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.CloseDBConn(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

}
