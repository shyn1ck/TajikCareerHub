package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetJobCategoryByID(id uint) (category models.JobCategory, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.GetJobCategoryByID]: Error retrieving job category with ID %v. Error: %v\n", id, err)
		return category, err
	}
	return category, nil
}

func GetAllJobCategories() (jobs []models.JobCategory, err error) {
	err = db.GetDBConn().Where("deleted_at = ?", false).Find(&models.JobCategory{}).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllJobCategories]: Error retrieving all jobs. Error: %v\n", err)
		return nil, err
	}
	return jobs, nil
}

func CreateJobCategory(category models.JobCategory) error {
	err := db.GetDBConn().Create(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.CreateJobCategory]: Failed to create job category. Error: %v\n", err)
		return err
	}
	return nil
}

func UpdateJobCategory(category models.JobCategory) error {
	err := db.GetDBConn().Model(&models.JobCategory{}).Where("id = ?", category.ID).Updates(category).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateJobCategory]: Failed to update job category with ID %v. Error: %v\n", category.ID, err)
		return err
	}
	return nil
}

func DeleteJobCategory(id uint) error {
	err := db.GetDBConn().Model(&models.JobCategory{}).Where("id = ?", id).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteJobCategory]: Failed to soft delete job category with ID %v. Error: %v\n", id, err)
		return err
	}
	return nil
}
