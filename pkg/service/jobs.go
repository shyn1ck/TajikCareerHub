package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func GetAllJobs() ([]models.Job, error) {
	return repository.GetAllJobs()
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

	// Обновляем только изменённые поля
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

func validateSalaryRange(minSalary, maxSalary string) error {
	if minSalary == "" || maxSalary == "" {
		return errors.New("salary range must not be empty")
	}
	min, err := strconv.ParseFloat(minSalary, 64)
	if err != nil {
		return errors.New("invalid format for minSalary")
	}

	max, err := strconv.ParseFloat(maxSalary, 64)
	if err != nil {
		return errors.New("invalid format for maxSalary")
	}

	if min > max {
		return errors.New("minSalary cannot be greater than maxSalary")
	}

	return nil
}

func GetJobsBySalaryRange(minSalary, maxSalary string) ([]models.Job, error) {
	if err := validateSalaryRange(minSalary, maxSalary); err != nil {
		return nil, err
	}

	jobs, err := repository.FilterJobsBySalaryRange(minSalary, maxSalary)
	if err != nil {
		return nil, err
	}

	return jobs, nil
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

func GetJobsBySalary(minSalary, maxSalary string) ([]models.Job, error) {
	if err := validateSalaryRange(minSalary, maxSalary); err != nil {
		return nil, err
	}
	jobs, err := repository.FilterJobsBySalary(minSalary, maxSalary)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}
