package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllApplications() (applications []models.Application, err error) {
	err = db.GetDBConn().Where("deleted_at = ?", false).Find(&applications).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllApplications]: Error retrieving all applications. Error: %v\n", err)
		return nil, err
	}
	logger.Info.Println("[repository.GetAllApplications]: Successfully retrieved all applications.")
	return applications, nil
}

func GetApplicationByID(id uint) (application models.Application, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&application).Error
	if err != nil {
		logger.Error.Printf("[repository.GetApplicationByID]: Error retrieving application with ID %v. Error: %v\n", id, err)
		return application, err
	}
	logger.Info.Printf("[repository.GetApplicationByID]: Successfully retrieved application with ID %v.\n", id)
	return application, nil
}

func GetApplicationsByUserID(userID uint) (applications []models.Application, err error) {
	err = db.GetDBConn().Where("user_id = ? AND deleted_at = ?", userID, false).Find(&applications).Error
	if err != nil {
		logger.Error.Printf("[repository.GetApplicationsByUserID]: Error retrieving applications for user ID %v. Error: %v\n", userID, err)
		return nil, err
	}
	logger.Info.Printf("[repository.GetApplicationsByUserID]: Successfully retrieved applications for user ID %v.\n", userID)
	return applications, nil
}

func GetApplicationsByJobID(jobID uint) (applications []models.Application, err error) {
	err = db.GetDBConn().Where("job_id = ? AND deleted_at = ?", jobID, false).Find(&applications).Error
	if err != nil {
		logger.Error.Printf("[repository.GetApplicationsByJobID]: Error retrieving applications for job ID %v. Error: %v\n", jobID, err)
		return nil, err
	}
	logger.Info.Printf("[repository.GetApplicationsByJobID]: Successfully retrieved applications for job ID %v.\n", jobID)
	return applications, nil
}

func AddApplication(application models.Application) error {
	logger.Info.Printf("[repository.AddApplication]: Adding new application for user ID %v and job ID %v...\n", application.UserID, application.JobID)
	err := db.GetDBConn().Create(&application).Error
	if err != nil {
		logger.Error.Printf("[repository.AddApplication]: Failed to add application. Error: %v\n", err)
		return err
	}
	logger.Info.Printf("[repository.AddApplication]: Successfully added application for user ID %v and job ID %v.\n", application.UserID, application.JobID)
	return nil
}

func UpdateApplication(application models.Application) error {
	logger.Info.Printf("[repository.UpdateApplication]: Updating application with ID %v...\n", application.ID)
	err := db.GetDBConn().Save(&application).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateApplication]: Failed to update application with ID %v. Error: %v\n", application.ID, err)
		return err
	}
	logger.Info.Printf("[repository.UpdateApplication]: Successfully updated application with ID %v.\n", application.ID)
	return nil
}

func DeleteApplication(id uint) error {
	logger.Info.Printf("[repository.DeleteApplication]: Soft deleting application with ID %v...\n", id)
	err := db.GetDBConn().Model(&models.Application{}).Where("id = ?", id).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteApplication]: Failed to soft delete application with ID %v. Error: %v\n", id, err)
		return err
	}
	logger.Info.Printf("[repository.DeleteApplication]: Successfully soft deleted application with ID %v.\n", id)
	return nil
}
