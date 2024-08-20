package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"log"
)

func GetJobCategoryByID(id uint) (category models.JobCategory, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.GetJobCategoryByID]: Error retrieving job category with ID %v. Error: %v\n", id, err)
		return category, err
	}
	logger.Info.Printf("[repository.GetJobCategoryByID]: Successfully retrieved job category with ID %v.\n", id)
	return category, nil
}

func GetAllJobCategories() (jobs []models.JobCategory, err error) {
	err = db.GetDBConn().Where("deleted_at = ?", false).Find(&models.JobCategory{}).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllJobCategories]: Error retrieving all jobs. Error: %v\n", err)
		return nil, err
	}
	logger.Info.Println("[repository.GetAllJobCategories]: Successfully retrieved all jobs.")
	return jobs, nil
}

func CreateJobCategory(category models.JobCategory) error {
	err := db.GetDBConn().Create(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.CreateJobCategory]: Failed to create job category. Error: %v\n", err)
		return err
	}
	logger.Info.Println("[repository.CreateJobCategory]: Job category created successfully.")
	return nil
}

func UpdateJobCategory(category models.JobCategory) error {
	err := db.GetDBConn().Model(&models.JobCategory{}).Where("id = ?", category.ID).Updates(category).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateJobCategory]: Failed to update job category with ID %v. Error: %v\n", category.ID, err)
		return err
	}
	logger.Info.Printf("[repository.UpdateJobCategory]: Job category with ID %v updated successfully.\n", category.ID)
	return nil
}

func DeleteJobCategory(id uint) error {
	err := db.GetDBConn().Model(&models.JobCategory{}).Where("id = ?", id).Update("deleted_at", true).Error
	if err != nil {
		log.Printf("[repository.DeleteJobCategory]: Failed to soft delete job category with ID %v. Error: %v\n", id, err)
		return err
	}
	log.Printf("[repository.DeleteJobCategory]: Job category with ID %v successfully soft deleted.\n", id)
	return nil
}
