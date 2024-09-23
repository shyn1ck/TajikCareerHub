package db

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"errors"
	"fmt"
)

func Migrate() error {
	if dbConn == nil {
		return errors.New("database connection is not initialized")
	}
	migrateModels := []interface{}{
		&models.Vacancy{},
		&models.User{},
		&models.Application{},
		&models.Company{},
		&models.VacancyCategory{},
		&models.ResumeView{},
		&models.Resume{},
		&models.VacancyView{},
		&models.ApplicationStatus{},
	}
	for _, model := range migrateModels {
		err := dbConn.AutoMigrate(model)
		if err != nil {
			return fmt.Errorf("failed to migrate %T: %v", model, err)
		}
		logger.Info.Printf("Migrated model: %T\n", model)
	}
	initialStatuses := []models.ApplicationStatus{
		{Name: "applied"},
		{Name: "under_review"},
		{Name: "rejected"},
		{Name: "interview"},
	}

	var count int64
	err := dbConn.Model(&models.ApplicationStatus{}).Count(&count).Error
	if err != nil {
		return fmt.Errorf("failed to count application statuses: %v", err)
	}

	if count == 0 {
		if err := dbConn.Create(&initialStatuses).Error; err != nil {
			return fmt.Errorf("failed to insert initial application statuses: %v", err)
		}
		logger.Info.Println("Initial application statuses inserted successfully")
	}

	logger.Info.Println("Database migration completed successfully")
	return nil
}
