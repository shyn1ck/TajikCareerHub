package db

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"errors"
)

func Migrate() error {
	if dbConn == nil {
		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(
		&models.Job{},
		&models.User{},
		&models.Application{},
		&models.Company{},
		&models.Favorite{},
		&models.JobCategory{},
	)
	if err != nil {
		return errors.New("failed to migrate database schema: " + err.Error())
	}

	logger.Info.Println("Database migration completed successfully")
	return nil
}
