package db

import (
	"TajikCareerHub/models"
	"errors"
	"log"
)

func Migrate() error {
	if dbConn == nil {
		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(&models.Job{}, &models.User{})
	if err != nil {
		return errors.New("failed to migrate database schema: " + err.Error())
	}

	log.Println("Database migration completed successfully")
	return nil
}
