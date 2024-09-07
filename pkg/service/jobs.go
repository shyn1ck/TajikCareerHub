package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func GetAllJobs(keyword, minSalary, maxSalary, location, category string) (jobs []models.Job, err error) {
	jobs, err = repository.GetAllJobs(keyword, minSalary, maxSalary, location, category)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func GetJobByID(id uint) (models.Job, error) {
	return repository.GetJobByID(id)
}

func AddJob(job models.Job) error {
	_, err := repository.GetJobByID(job.ID)
	if err == nil {
		return errors.New("job with the same ID already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return repository.AddJob(job)
}

func UpdateJob(jobID uint, updatedJob models.Job) error {
	job, err := repository.GetJobByID(jobID)
	if err != nil {
		return err
	}
	if updatedJob.Title != "" {
		job.Title = updatedJob.Title
	}
	if updatedJob.Description != "" {
		job.Description = updatedJob.Description
	}
	if updatedJob.Location != "" {
		job.Location = updatedJob.Location
	}
	if updatedJob.JobCategory.ID != 0 {
		job.JobCategory = updatedJob.JobCategory
	}
	if updatedJob.Salary != "" {
		job.Salary = updatedJob.Salary
	}

	return repository.UpdateJob(jobID, job)
}

func DeleteJob(jobID uint) error {
	return repository.DeleteJob(jobID)
}

func FilterJobs(location string, category string) ([]models.Job, error) {
	return repository.FilterJobs(location, category)
}

func UpdateJobSalary(jobID uint, newSalary string) error {
	if _, err := strconv.ParseFloat(newSalary, 64); err != nil {
		return errors.New("invalid salary format")
	}
	err := repository.UpdateJobSalary(jobID, newSalary)
	if err != nil {
		return err
	}

	return nil
}
