package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllJobs() (jobs []models.Job, err error) {
	err = db.GetDBConn().Where("deleted_at IS NULL").Find(&jobs).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllJobs]: Error retrieving all jobs. Error: %v\n", err)
		return nil, err
	}
	logger.Info.Println("[repository.GetAllJobs]: Successfully retrieved all jobs.")
	return jobs, nil
}

func GetJobByID(id uint) (job models.Job, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&job).Error
	if err != nil {
		logger.Error.Printf("[repository.GetJobByID]: Error retrieving job with ID %v. Error: %v\n", id, err)
		return models.Job{}, err
	}
	logger.Info.Printf("[repository.GetJobByID]: Successfully retrieved job with ID %v.\n", id)
	return job, nil
}

func AddJob(job models.Job) error {
	logger.Info.Println("[repository.AddJob]: Adding new job to the database")
	result := db.GetDBConn().Create(&job)
	if result.Error != nil {
		logger.Error.Printf("[repository.AddJob]: Failed to add job, error: %v\n", result.Error)
		return result.Error
	}
	logger.Info.Println("[repository.AddJob]: Job added successfully")
	return nil
}

func UpdateJob(jobID uint, job models.Job) error {
	logger.Info.Printf("[repository.UpdateJob]: Updating job with ID %v in the database...\n", jobID)
	err := db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Updates(job).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateJob]: Failed to update job with ID %v. Error: %v\n", jobID, err)
		return err
	}
	logger.Info.Printf("[repository.UpdateJob]: Job with ID %v updated successfully.\n", jobID)
	return nil
}

func DeleteJob(jobID uint) error {
	logger.Info.Printf("[repository.DeleteJob]: Soft deleting job with ID %v...\n", jobID)
	err := db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteJob]: Failed to soft delete job with ID %v. Error: %v\n", jobID, err)
		return err
	}
	logger.Info.Printf("[repository.DeleteJob]: Job with ID %v successfully soft deleted.\n", jobID)
	return nil
}

func FilterJobs(location string, category string) (jobs []models.Job, err error) {
	logger.Info.Printf("[repository.FilterJobs]: Filtering jobs by location %v and category %v...\n", location, category)
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
	logger.Info.Printf("[repository.FilterJobs]: Successfully filtered jobs by location %v and category %v.\n", location, category)
	return jobs, nil
}

func FilterJobsBySalary(minSalary, maxSalary string) (jobs []models.Job, err error) {
	logger.Info.Printf("[repository.FilterJobsBySalary]: Filtering jobs with salary between %v and %v...\n", minSalary, maxSalary)
	query := db.GetDBConn()
	if minSalary != "" && maxSalary != "" {
		query = query.Where("salary BETWEEN ? AND ?", minSalary, maxSalary)
	}
	err = query.Find(&jobs).Error
	if err != nil {
		logger.Error.Printf("[repository.FilterJobsBySalary]: Failed to filter jobs by salary. Error: %v\n", err)
		return nil, err
	}
	logger.Info.Printf("[repository.FilterJobsBySalary]: Successfully filtered jobs by salary between %v and %v.\n", minSalary, maxSalary)
	return jobs, nil
}

func UpdateJobSalary(jobID uint, newSalary string) error {
	logger.Info.Printf("[repository.UpdateJobSalary]: Updating salary for job with ID %v...\n", jobID)
	err := db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Update("salary", newSalary).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateJobSalary]: Failed to update salary for job with ID %v. Error: %v\n", jobID, err)
		return err
	}
	logger.Info.Printf("[repository.UpdateJobSalary]: Salary for job with ID %v updated successfully.\n", jobID)
	return nil
}

func FilterJobsBySalaryRange(minSalary, maxSalary string) (jobs []models.Job, err error) {
	logger.Info.Printf("[repository.FilterJobsBySalaryRange]: Filtering jobs with salary range between %v and %v...\n", minSalary, maxSalary)
	query := db.GetDBConn()
	if minSalary != "" && maxSalary != "" {
		query = query.Where("salary >= ? AND salary <= ?", minSalary, maxSalary)
	}
	err = query.Find(&jobs).Error
	if err != nil {
		logger.Error.Printf("[repository.FilterJobsBySalaryRange]: Failed to filter jobs by salary range. Error: %v\n", err)
		return nil, err
	}
	logger.Info.Printf("[repository.FilterJobsBySalaryRange]: Successfully filtered jobs by salary range between %v and %v.\n", minSalary, maxSalary)
	return jobs, nil
}
