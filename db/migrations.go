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
		&models.Vacancy{},
		&models.User{},
		&models.Application{},
		&models.Company{},
		&models.VacancyCategory{},
		&models.ResumeView{},
		&models.Resume{},
		&models.VacancyView{},
		&models.ApplicationStatus{},
	)
	if err != nil {
		return errors.New("failed to migrate database schema: " + err.Error())
	}
	initialStatuses := []models.ApplicationStatus{
		{Name: "applied"},
		{Name: "under_review"},
		{Name: "rejected"},
		{Name: "interview"},
	}
	var count int64
	err = dbConn.Model(&models.ApplicationStatus{}).Count(&count).Error
	if err != nil {
		return errors.New("failed to count application statuses: " + err.Error())
	}
	if count == 0 {
		if err := dbConn.Create(&initialStatuses).Error; err != nil {
			return errors.New("failed to insert initial application statuses: " + err.Error())
		}
		logger.Info.Println("Initial application statuses inserted successfully")
	}
	logger.Info.Println("Database migration completed successfully")
	return nil
}
