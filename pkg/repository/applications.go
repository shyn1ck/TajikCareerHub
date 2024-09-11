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
		Where("deleted_at = false").Find(&applications).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllApplications]: Error retrieving all applications. Error: %v\n", err)
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
		logger.Error.Printf("[repository.GetApplicationByID]: Error retrieving application with ID %v. Error: %v\n", id, err)
		return application, err
	}
	return application, nil
}

func GetApplicationsByUserID(userID uint) ([]models.Application, error) {
	var applications []models.Application
	err := db.GetDBConn().
		Preload("User").
		Preload("Vacancy").
		Preload("Resume").
		Where("user_id = ? AND deleted_at = false", userID).
		Find(&applications).Error
	if err != nil {
		logger.Error.Printf("[repository.GetApplicationsByUserID]: Error retrieving applications for user ID %v. Error: %v\n", userID, err)
		return nil, err
	}
	return applications, nil
}

func GetApplicationsByVacancyID(vacancyID uint) ([]models.Application, error) {
	var applications []models.Application
	err := db.GetDBConn().
		Preload("User").
		Preload("Vacancy").
		Preload("Resume").
		Where("vacancy_id = ? AND deleted_at = false", vacancyID).
		Find(&applications).Error
	if err != nil {
		logger.Error.Printf("[repository.GetApplicationsByVacancyID]: Error retrieving applications for vacancy ID %v. Error: %v\n", vacancyID, err)
		return nil, err
	}
	return applications, nil
}

func AddApplication(application models.Application) error {
	err := db.GetDBConn().Create(&application).Error
	if err != nil {
		logger.Error.Printf("[repository.AddApplication]: Error adding application. Error: %v\n", err)
		return err
	}
	return nil
}

func UpdateApplication(application models.Application) error {
	err := db.GetDBConn().Save(&application).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateApplication]: Failed to update application with ID %v. Error: %v\n", application.ID, err)
		return err
	}
	return nil
}

func DeleteApplication(id uint) error {
	err := db.GetDBConn().
		Model(&models.Application{}).
		Where("id = ?", id).
		Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteApplication]: Failed to soft delete application with ID %v. Error: %v\n", id, err)
		return err
	}
	return nil
}
