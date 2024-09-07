package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllJobs(keyword, location, category string) (jobs []models.Job, err error) {

	query := db.GetDBConn().Model(&models.Job{}).Where("deleted_at IS NULL")
	if keyword != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if location != "" {
		query = query.Where("location = ?", location)
	}
	if category != "" {
		query = query.Joins("JOIN job_categories ON jobs.job_category_id = job_categories.id").Where("job_categories.name = ?", category)
	}

	err = query.Find(&jobs).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllJobs]: Error retrieving jobs. Error: %v\n", err)
		return nil, err
	}
	return jobs, nil
}

func GetJobByID(id uint) (job models.Job, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&job).Error
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

func FilterJobs(location string, category string) (jobs []models.Job, err error) {
	query := db.GetDBConn()
	if location != "" {
		query = query.Where("location = ?", location)
	}
	if category != "" {
		query = query.Joins("JOIN job_categories ON jobs.job_category_id = job_categories.id").
			Where("job_categories.name = ?", category)
	}
	err = query.Find(&jobs).Error
	if err != nil {
		logger.Error.Printf("[repository.FilterJobs]: Failed to filter jobs. Error: %v\n", err)
		return nil, err
	}
	return jobs, nil
}

func UpdateJobSalary(jobID uint, newSalary string) error {
	err := db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Update("salary", newSalary).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateJobSalary]: Failed to update salary for job with ID %v. Error: %v\n", jobID, err)
		return err
	}
	return nil
}
