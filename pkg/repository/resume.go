package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"errors"
	"gorm.io/gorm"
)

func GetAllResumes(search string, minExperienceYears int, location string, category string) (resumes []models.Resume, err error) {
	query := db.GetDBConn().
		Preload("VacancyCategory").
		Model(&models.Resume{}).
		Where("deleted_at = false")
	if search != "" {
		query = query.Where("summary ILIKE ?", "%"+search+"%")
	}
	if location != "" {
		query = query.Where("location = ?", location)
	}
	if category != "" {
		query = query.Joins("JOIN vacancy_categories ON vacancy_categories.id = resumes.vacancy_category_id").
			Where("vacancy_categories.name = ?", category)
	}
	if minExperienceYears > 0 {
		query = query.Where("experience_years >= ?", minExperienceYears)
	}
	err = query.Find(&resumes).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllResumes] Error fetching resumes: %v", err)
		return nil, TranslateError(err)
	}
	return resumes, nil
}

func GetResumeByID(id uint) (resume models.Resume, err error) {
	err = db.GetDBConn().
		Where("id = ?", id).
		Where("deleted_at = false").
		First(&resume).Error
	if err != nil {
		logger.Error.Printf("[repository.GetResumeByID] Error getting resume by ID %v: %v\n", id, err)
		return resume, TranslateError(err)
	}
	return resume, nil
}

func AddResume(resume models.Resume) (err error) {
	if err := db.GetDBConn().Create(&resume).Error; err != nil {
		logger.Error.Printf("[repository.AddResume]: Failed to add resume, error: %v\n", err)
		return TranslateError(err)
	}
	return nil
}

func UpdateResume(resumeID uint, resume models.Resume) (err error) {
	err = db.GetDBConn().
		Model(&models.Resume{}).
		Where("id = ? AND deleted_at = false", resumeID).
		Updates(resume).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateResume]: Failed to update resume with ID %v. Error: %v\n", resumeID, err)
		return TranslateError(err)
	}
	return nil
}

func DeleteResume(id uint) (err error) {
	err = db.GetDBConn().
		Model(&models.Resume{}).
		Where("id = ?", id).
		Update("deleted_at", true).
		Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteResume] Failed to delete resume with ID %v: %v\n", id, err)
		return TranslateError(err)
	}
	return nil
}

func RecordResumeView(userID uint, resumeID uint) (err error) {
	var resumeView models.ResumeView
	err = db.GetDBConn().
		Model(&models.ResumeView{}).
		Where("user_id = ? AND resume_id = ?", userID, resumeID).
		First(&resumeView).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.GetDBConn().
				Create(&models.ResumeView{
					UserID:   userID,
					ResumeID: resumeID,
					Count:    1,
				}).Error
			if err != nil {
				logger.Error.Printf("[repository.RecordResumeView] Error creating view record for resume ID %v. Error: %v\n", resumeID, err)
				return TranslateError(err)
			}
			logger.Info.Printf("[repository.RecordResumeView] Created new view record for resume ID %v by user ID %v.\n", resumeID, userID)
		} else {
			logger.Error.Printf("[repository.RecordResumeView] Error checking view record for resume ID %v. Error: %v\n", resumeID, err)
			return TranslateError(err)
		}
	} else {
		logger.Info.Printf("[repository.RecordResumeView] User ID %v already viewed resume ID %v.\n", userID, resumeID)
	}
	return nil
}

func GetResumeReportByID(resumeID uint) (*models.ResumeReport, error) {
	var report models.ResumeReport
	err := db.GetDBConn().
		Table("resumes").
		Select(`
			resumes.id AS resume_id, 
			resumes.title AS resume_title, 
			COUNT(DISTINCT resume_views.user_id) AS views_count, 
			COUNT(DISTINCT applications.user_id) AS applications_count`).
		Joins("LEFT JOIN resume_views ON resume_views.resume_id = resumes.id").
		Joins("LEFT JOIN applications ON applications.resume_id = resumes.id").
		Where("resumes.id = ? AND resumes.deleted_at = false", resumeID).
		Group("resumes.id").
		Scan(&report).Error
	if err != nil {
		logger.Error.Printf("[repository.GetResumeReportByID] Error retrieving resume report: %v", err)
		return nil, TranslateError(err)
	}

	logger.Info.Printf("[repository.GetResumeReportByID] Successfully retrieved data: %v", report)
	return &report, nil
}

func updateBlockStatusResume(id uint, isBlocked bool) (err error) {
	err = db.GetDBConn().Model(&models.Resume{}).Where("id = ?", id).Update("is_blocked", isBlocked).Error
	if err != nil {
		action := "block"
		if !isBlocked {
			action = "unblock"
		}
		logger.Error.Printf("[repository.updateBlockStatus] Failed to %s user with ID %v: %v\n", action, id, err)
		return TranslateError(err)
	}
	return nil
}

func BlockResume(id uint) (err error) {
	return updateBlockStatusResume(id, true)
}

func UnblockResume(id uint) (err error) {
	return updateBlockStatusResume(id, false)
}
