package main

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/pkg/controllers"
)

func main() {
	err := logger.Init()
	if err != nil {
		return
	}
	if err := db.ConnectToDB(); err != nil {
		logger.Error.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.CloseDBConn(); err != nil {
			logger.Error.Printf("Error closing database connection: %v", err)
		}
	}()
	if err := db.Migrate(); err != nil {
		logger.Error.Fatalf("Failed to run database migrations: %v", err)
	}
	err = controllers.RunRoutes()
	if err != nil {
		logger.Error.Fatalf("Failed to run routes: %v", err)
	}

}
