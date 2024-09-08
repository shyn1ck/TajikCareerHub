package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllResumes(search string, minExperienceYears int, location string, category string) ([]models.Resume, error) {
	var resumes []models.Resume
	query := db.GetDBConn().Preload("JobCategory").Model(&models.Resume{})
	if search != "" {
		query = query.Where("summary ILIKE ?", "%"+search+"%")
	}
	if location != "" {
		query = query.Where("location = ?", location)
	}
	if category != "" {
		query = query.Joins("JOIN job_categories ON job_categories.id = resumes.job_category_id").
			Where("job_categories.name = ?", category)
	}

	if minExperienceYears > 0 {
		query = query.Where("experience_years >= ?", minExperienceYears)
	}

	err := query.Find(&resumes).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllResumes] Error fetching resumes: %v", err)
		return nil, err
	}

	return resumes, nil
}

func GetResumeByID(id uint) (models.Resume, error) {
	var resume models.Resume
	err := db.GetDBConn().Where("id = ?", id).First(&resume).Error
	if err != nil {
		logger.Error.Printf("[repository.GetResumeByID] error getting resume by ID %v: %v\n", id, err)
		return resume, errs.TranslateError(err)
	}
	return resume, nil
}

func AddResume(resume models.Resume) error {
	if err := db.GetDBConn().Create(&resume).Error; err != nil {
		logger.Error.Printf("[repository.AddResume]: Failed to add resume, error: %v\n", err)
		return errs.TranslateError(err)
	}
	return nil
}

func UpdateResume(resume models.Resume) error {
	err := db.GetDBConn().Save(&resume).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateResume]: Failed to update resume: %v\n", err)
		return errs.TranslateError(err)
	}
	return nil
}

func DeleteResume(id uint) error {
	if err := db.GetDBConn().Where("id = ?", id).Delete(&models.Resume{}).Error; err != nil {
		logger.Error.Printf("[repository.DeleteResume] Failed to delete resume with ID %v: %v\n", id, err)
		return errs.TranslateError(err)
	}
	return nil
}
