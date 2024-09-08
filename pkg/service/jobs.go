package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
)

func GetAllJobs(search string, minSalary int, maxSalary int, location string, category string, sort string) (jobs []models.Job, err error) {
	jobs, err = repository.GetAllJobs(search, minSalary, maxSalary, location, category, sort)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func GetJobByID(id uint) (models.Job, error) {
	return repository.GetJobByID(id)
}

func AddJob(job models.Job) error {
	err := repository.AddJob(job)
	if err != nil {
		return err
	}
	return nil
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
	if updatedJob.Salary != 0 {
		job.Salary = updatedJob.Salary
	}

	return repository.UpdateJob(jobID, job)
}

func DeleteJob(jobID uint) error {
	return repository.DeleteJob(jobID)
}
