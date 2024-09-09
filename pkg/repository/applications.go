package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"gorm.io/gorm"
)

func GetAllApplications() (applications []models.Application, err error) {
	err = db.GetDBConn().
		Preload("User").
		Preload("Job").
		Preload("Resume").
		Where("deleted_at is NULL").Find(&applications).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllApplications]: Error retrieving all applications. Error: %v\n", err)
		return nil, err
	}
	return applications, nil
}

func GetApplicationByID(id uint) (application models.Application, err error) {
	err = db.GetDBConn().
		Preload("User").
		Preload("Job").
		Preload("Resume").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&application).Error
	if err != nil {
		logger.Error.Printf("[repository.GetApplicationByID]: Error retrieving application with ID %v. Error: %v\n", id, err)
		return application, err
	}
	return application, nil
}

func GetApplicationsByUserID(userID uint) (applications []models.Application, err error) {
	err = db.GetDBConn().
		Preload("User").
		Preload("Job").
		Preload("Resume").
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Find(&applications).Error
	if err != nil {
		logger.Error.Printf("[repository.GetApplicationsByUserID]: Error retrieving applications for user ID %v. Error: %v\n", userID, err)
		return nil, err
	}
	return applications, nil
}

func GetApplicationsByJobID(jobID uint) (applications []models.Application, err error) {
	err = db.GetDBConn().
		Preload("User").
		Preload("Job").
		Preload("Resume").
		Where("job_id = ? AND deleted_at IS NULL", jobID).
		Find(&applications).Error
	if err != nil {
		logger.Error.Printf("[repository.GetApplicationsByJobID]: Error retrieving applications for job ID %v. Error: %v\n", jobID, err)
		return nil, err
	}
	return applications, nil
}

func AddApplication(application models.Application) error {
	err := db.GetDBConn().Create(&application).Error
	if err != nil {
		logger.Error.Printf("[repository.AddApplication]: Failed to add application. Error: %v\n", err)
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
		Update("deleted_at", gorm.Expr("NOW()")).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteApplication]: Failed to soft delete application with ID %v. Error: %v\n", id, err)
		return err
	}
	return nil
}
