package main

import (
	"TajikCareerHub/configs"
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/pkg/controllers"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Error.Fatalf("Ошибка загрузки .env файла: %s", err)
	}
	err := configs.ReadSettings()
	if err != nil {
		panic(err)
	}
	fmt.Print("New commit for chevrons ")
	err = logger.Init()
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
