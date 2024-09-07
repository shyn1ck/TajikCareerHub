package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllJobs(keyword, minSalary, maxSalary, location, category string) ([]models.Job, error) {
	var jobs []models.Job
	query := db.GetDBConn()

	if keyword != "" {
		query = query.Where("title ILIKE ?", "%"+keyword+"%")
	}

	if minSalary != "" && maxSalary != "" {
		query = query.Where("salary BETWEEN ? AND ?", minSalary, maxSalary)
	} else if minSalary != "" {
		query = query.Where("salary >= ?", minSalary)
	} else if maxSalary != "" {
		query = query.Where("salary <= ?", maxSalary)
	}

	if location != "" {
		query = query.Where("location = ?", location)
	}
	if category != "" {
		query = query.Joins("JOIN job_categories ON job_categories.id = jobs.job_category_id").
			Where("job_categories.name = ?", category)
	}

	return jobs, nil
}

func GetJobByID(id uint) (models.Job, error) {
	var job models.Job
	err := db.GetDBConn().
		Preload("Company").
		Preload("JobCategory").
		Where("id = ?", id).
		First(&job).Error
	if err != nil {
		logger.Error.Printf("[repository.GetJobByID]: Error retrieving job with ID %v. Error: %v\n", id, err)
		return models.Job{}, err
	}
	return job, nil
}

func AddJob(job models.Job) error {
	result := db.GetDBConn().Create(&job)
	if result.Error != nil {
		logger.Error.Printf("[repository.AddJob]: Failed to add job, error: %v\n", result.Error)
		return result.Error
	}
	return nil
}

func UpdateJob(jobID uint, job models.Job) error {
	err := db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Updates(job).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateJob]: Failed to update job with ID %v. Error: %v\n", jobID, err)
		return err
	}
	return nil
}

func DeleteJob(jobID uint) error {
	err := db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteJob]: Failed to soft delete job with ID %v. Error: %v\n", jobID, err)
		return err
	}
	return nil
}

func UpdateJobSalary(jobID uint, newSalary string) error {
	err := db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Update("salary", newSalary).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateJobSalary]: Failed to update salary for job with ID %v. Error: %v\n", jobID, err)
		return err
	}
	return nil
}
