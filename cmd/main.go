package main

import (
	"TajikCareerHub/configs"
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/pkg/controllers"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	fmt.Println("JWT Secret Key:", os.Getenv("JWT_SECRET_KEY"))
	fmt.Println("JWT TTL Minutes:", configs.AppSettings.AuthParams.JwtTtlMinutes)

	err := godotenv.Load(".env")
	if err != nil {
		panic(errors.New(fmt.Sprintf("error loading .env file. Error is %s", err)))
	}

	err = configs.ReadSettings()
	if err != nil {
		panic(err)
	}
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
