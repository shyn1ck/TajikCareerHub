package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/models"
	"log"
)

func GetAllJobs() (jobs []models.Job, err error) {
	err = db.GetDBConn().Where("deleted_at = false").Find(&jobs).Error
	if err != nil {
		log.Printf("repository.GetAllJobs: Error retrieving all jobs. Error: %v\n", err)
		return nil, err
	}
	log.Println("repository.GetAllJobs: Successfully retrieved all jobs.")
	return jobs, nil
}

func GetJobByID(id uint) (job []models.Job, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&job).Error
	if err != nil {
		log.Printf("repository.GetJob: Error retrieving job. Error: %v\n", err)
		return nil, err
	}
	log.Println("repository.GetJob: Successfully retrieved job.")
	return job, nil
}

func AddJob(job []models.Job) error {
	log.Println("repository.AddJob: Adding new job to the database")
	result := db.GetDBConn().Create(&job)
	if result.Error != nil {
		log.Printf("repository.AddJob: Failed to add job, error: %v\n", result.Error)
		return result.Error
	}
	log.Println("repository.AddJob: Job added successfully")
	return nil
}

func UpdateJob(jobID uint, job models.Job) (err error) {
	log.Printf("repository.UpdateJob: Updating job with ID %v in the database...\n", jobID)
	err = db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Updates(job).Error
	if err != nil {
		log.Printf("repository.UpdateJob: Failed to update job with ID %v. Error: %v\n", jobID, err)
		return err
	}
	log.Printf("repository.UpdateJob: Job with ID %v updated successfully.\n", jobID)
	return nil
}

func DeleteJob(jobID uint) (err error) {
	log.Printf("repository.DeleteJob: Soft deleting job with ID %v...\n", jobID)
	err = db.GetDBConn().Model(&models.Job{}).Where("id = ?", jobID).Update("is_deleted", true).Error
	if err != nil {
		log.Printf("repository.DeleteJob: Failed to soft delete job with ID %v. Error: %v\n", jobID, err)
		return err
	}
	log.Printf("repository.DeleteJob: Job with ID %v successfully soft deleted.\n", jobID)
	return nil
}

func FilterJobs(location string, category string) (jobs []models.Job, err error) {
	log.Printf("repository.FilterJobs: Filtering jobs by location %v and category %v...\n", location, category)
	query := db.GetDBConn()
	if location != "" {
		query = query.Where("location = ?", location)
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}
	err = query.Find(&jobs).Error
	if err != nil {
		log.Printf("repository.FilterJobs: Failed to filter jobs. Error: %v\n", err)
		return nil, err
	}
	log.Printf("repository.FilterJobs: Successfully filtered jobs by location %v and category %v.\n", location, category)
	return jobs, nil
}
