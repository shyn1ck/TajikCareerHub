package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllApplications() ([]models.Application, error) {
	var applications []models.Application
	err := db.GetDBConn().
		Preload("User").
		Preload("Vacancy").
		Preload("Resume").
		Where("deleted_at = false").
		Find(&applications).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllApplications] Error fetching applications: %v", err)
		return nil, err
	}
	return applications, nil
}

func GetApplicationByID(id uint) (models.Application, error) {
	var application models.Application
	err := db.GetDBConn().
		Preload("User").
		Preload("Vacancy").
		Preload("Resume").
		Where("id = ? AND deleted_at = false", id).
		First(&application).Error
	if err != nil {
		logger.Error.Printf("[repository.GetApplicationByID] Error getting application by ID %v: %v\n", id, err)
		return application, TranslateError(err)
	}
	return application, nil
}

func AddApplication(application models.Application) error {
	if err := db.GetDBConn().Create(&application).Error; err != nil {
		logger.Error.Printf("[repository.AddApplication]: Failed to add application, error: %v\n", err)
		return TranslateError(err)
	}
	return nil
}

func UpdateApplication(applicationID uint, application models.Application) error {
	err := db.GetDBConn().
		Model(&models.Application{}).
		Where("id = ? AND deleted_at = false", applicationID).
		Updates(application).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateApplication]: Failed to update application with ID %v. Error: %v\n", applicationID, err)
		return TranslateError(err)
	}
	return nil
}

func DeleteApplication(id uint) error {
	err := db.GetDBConn().
		Model(&models.Application{}).
		Where("id = ?", id).
		Update("deleted_at", true).
		Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteApplication] Failed to delete application with ID %v: %v\n", id, err)
		return TranslateError(err)
	}
	return nil
}
