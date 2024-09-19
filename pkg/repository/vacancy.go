package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"errors"
	"gorm.io/gorm"
)

func GetAllVacancies(search string, minSalary int, maxSalary int, location string, category string, sort string) (vacancies []models.Vacancy, err error) {
	query := db.GetDBConn().
		Preload("Company").
		Preload("VacancyCategory").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "full_name", "email")
		}).
		Model(&models.Vacancy{}).
		Where("deleted_at = false")
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
		query = query.Joins("JOIN vacancy_categories ON vacancy_categories.id = vacancies.vacancy_category_id").
			Where("vacancy_categories.name = ?", category)
	}
	if sort == "asc" {
		query = query.Order("salary ASC")
	} else if sort == "desc" {
		query = query.Order("salary DESC")
	}
	err = query.Find(&vacancies).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllVacancies] Error fetching vacancies: %v", err)
		return nil, TranslateError(err)
	}
	return vacancies, nil
}

func GetVacancyByID(id uint) (vacancy models.Vacancy, err error) {
	err = db.GetDBConn().
		Preload("Company").
		Preload("VacancyCategory").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "full_name", "email")
		}).
		Where("id = ? AND deleted_at = false", id).
		First(&vacancy).Error

	if err != nil {
		logger.Error.Printf("[repository.GetVacancyByID] Error retrieving vacancy with ID %v. Error: %v\n", id, err)
		return models.Vacancy{}, TranslateError(err)
	}
	return vacancy, nil
}

func AddVacancy(vacancy models.Vacancy) (err error) {
	if err = db.GetDBConn().Create(&vacancy).Error; err != nil {
		logger.Error.Printf("[repository.AddVacancy]: Failed to add vacancy, error: %v\n", err)
		return TranslateError(err)
	}
	return nil
}

func UpdateVacancy(vacancyID uint, vacancy models.Vacancy) (err error) {
	err = db.GetDBConn().Model(&models.Vacancy{}).Where("id = ? AND deleted_at = false", vacancyID).Updates(vacancy).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateVacancy]: Failed to update vacancy with ID %v. Error: %v\n", vacancyID, err)
		return TranslateError(err)
	}
	return nil
}

func DeleteVacancy(vacancyID uint) (err error) {
	err = db.GetDBConn().Model(&models.Vacancy{}).Where("id = ?", vacancyID).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteVacancy] Failed to soft delete vacancy with ID %v. Error: %v\n", vacancyID, err)
		return TranslateError(err)
	}
	return nil
}

func GetVacancyReportByID(vacancyID uint) (*models.VacancyReport, error) {
	var report models.VacancyReport
	err := db.GetDBConn().
		Table("vacancies").
		Select(`
			vacancies.id AS vacancy_id, 
			vacancies.title AS vacancy_title, 
			COUNT(DISTINCT vacancy_views.user_id) AS views_count, 
			COUNT(DISTINCT applications.user_id) AS applications_count`).
		Joins("LEFT JOIN vacancy_views ON vacancy_views.vacancy_id = vacancies.id").
		Joins("LEFT JOIN applications ON applications.vacancy_id = vacancies.id").
		Where("vacancies.id = ? AND vacancies.deleted_at = false", vacancyID).
		Group("vacancies.id").
		Scan(&report).Error
	if err != nil {
		logger.Error.Printf("[repository.GetVacancyReportByID] Error retrieving vacancy report: %v", err)
		return nil, TranslateError(err)
	}
	logger.Info.Printf("[repository.GetVacancyReportByID] Successfully retrieved data: %v", report)
	return &report, nil
}

func RecordVacancyView(userID uint, vacancyID uint) (err error) {
	var vacancyView models.VacancyView
	err = db.GetDBConn().
		Model(&models.VacancyView{}).
		Where("user_id = ? AND vacancy_id = ?", userID, vacancyID).
		First(&vacancyView).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.GetDBConn().
				Create(&models.VacancyView{
					UserID:    userID,
					VacancyID: vacancyID,
					Count:     1,
				}).Error
			if err != nil {
				logger.Error.Printf("[repository.RecordView] Error creating view record for vacancy ID %v. Error: %v\n", vacancyID, err)
				return TranslateError(err)
			}
		} else {
			logger.Error.Printf("[repository.RecordView] Error checking view record for vacancy ID %v. Error: %v\n", vacancyID, err)
			return TranslateError(err)
		}
	} else {
		logger.Info.Printf("[repository.RecordView] User ID %v already viewed vacancy ID %v.\n", userID, vacancyID)
	}
	return nil
}

func updateBlockStatusJob(id uint, isBlocked bool) (err error) {
	err = db.GetDBConn().Model(&models.Vacancy{}).Where("id = ?", id).Update("is_blocked", isBlocked).Error
	if err != nil {
		action := "block"
		if !isBlocked {
			action = "unblock"
		}
		logger.Error.Printf("[repository.updateBlockStatusJob] Failed to %s job with ID %v: %v\n", action, id, err)
		return TranslateError(err)
	}
	return nil
}

func BlockVacancy(id uint) (err error) {
	return updateBlockStatusJob(id, true)
}

func UnblockVacancy(id uint) (err error) {
	return updateBlockStatusJob(id, false)
}
