package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllJobs(search string, minSalary int, maxSalary int, location string, category string, sort string) ([]models.Job, error) {
	var jobs []models.Job
	query := db.GetDBConn().Preload("Company").Preload("JobCategory").Model(&models.Job{})

	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	if minSalary > 0 && maxSalary > 0 {
		query = query.Where("salary BETWEEN ? AND ?", minSalary, maxSalary)
	} else if minSalary > 0 {
		query = query.Where("salary >= ?", minSalary)
	} else if maxSalary > 0 {
		query = query.Where("salary <= ?", maxSalary)
	}

	if location != "" {
		query = query.Where("location = ?", location)
	}
	if category != "" {
		query = query.Joins("JOIN job_categories ON job_categories.id = jobs.job_category_id").
			Where("job_categories.name = ?", category)
	}

	if sort == "asc" {
		query = query.Order("salary ASC")
	} else if sort == "desc" {
		query = query.Order("salary DESC")
	}

	err := query.Find(&jobs).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllJobs] Error fetching jobs: %v", err)
		return nil, err
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
		return models.Job{}, errs.TranslateError(err)
	}
	return job, nil
}

func AddJob(job models.Job) error {
	if err := db.GetDBConn().Create(&job).Error; err != nil {
		logger.Error.Printf("[repository.AddJob]: Failed to add job, error: %v\n", err)
		return errs.TranslateError(err)
	}
	return nil
}

func UpdateJob(jobID uint, job models.Job) error {
	err := db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Updates(job).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateJob]: Failed to update job with ID %v. Error: %v\n", jobID, err)
		return errs.TranslateError(err)
	}
	return nil
}

func DeleteJob(jobID uint) error {
	err := db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteJob]: Failed to soft delete job with ID %v. Error: %v\n", jobID, err)
		return errs.TranslateError(err)
	}
	return nil
}

func UpdateJobSalary(jobID uint, newSalary string) error {
	err := db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Update("salary", newSalary).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateJobSalary]: Failed to update salary for job with ID %v. Error: %v\n", jobID, err)
		return errs.TranslateError(err)
	}
	return nil
}
