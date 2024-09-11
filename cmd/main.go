package main

import (
	"TajikCareerHub/configs"
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/pkg/controllers"
	"TajikCareerHub/server"
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title TajikCareerHub API
// @version 1.0
// @description API Server for TajikCareerHub Application

// @host localhost:8181
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Error.Fatalf("Ошибка загрузки .env файла: %s", err)
	}
	err := configs.ReadSettings()
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

	mainServer := new(server.Server)
	go func() {
		if err = mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil {
			log.Fatalf("Error to start HTTP server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if sqlDB, err := db.GetDBConn().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			logger.Error.Fatalf("Error to close DB: %s", err)
		}
	} else {
		logger.Error.Fatalf("Error to get *sql.DB from GORM: %s", err)
	}

	if err = mainServer.Shutdown(context.Background()); err != nil {
		logger.Error.Fatalf("Error during server shutdown: %s", err)
	}
}
