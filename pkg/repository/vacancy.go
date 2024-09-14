package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"gorm.io/gorm"
)

func GetAllVacancies(search string, minSalary int, maxSalary int, location string, category string, sort string) ([]models.Vacancy, error) {
	var vacancies []models.Vacancy
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

	err := query.Find(&vacancies).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllVacancies] Error fetching vacancies: %v", err)
		return nil, err
	}

	return vacancies, nil
}

func GetVacancyByID(id uint) (models.Vacancy, error) {
	var vacancy models.Vacancy
	err := db.GetDBConn().
		Preload("Company").
		Preload("VacancyCategory").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "full_name", "email")
		}).
		Where("id = ?", id).
		Where("deleted_at = false").
		First(&vacancy).Error
	if err != nil {
		logger.Error.Printf("[repository.GetVacancyByID] Error retrieving vacancy with ID %v. Error: %v\n", id, err)
		return models.Vacancy{}, TranslateError(err)
	}
	return vacancy, nil
}

func AddVacancy(vacancy models.Vacancy) error {
	if err := db.GetDBConn().Create(&vacancy).Error; err != nil {
		logger.Error.Printf("[repository.AddVacancy]: Failed to add vacancy, error: %v\n", err)
		return TranslateError(err)
	}
	return nil
}

func UpdateVacancy(vacancyID uint, vacancy models.Vacancy) error {
	err := db.GetDBConn().Model(&models.Vacancy{}).Where("id = ? AND deleted_at = false", vacancyID).Updates(vacancy).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateVacancy]: Failed to update vacancy with ID %v. Error: %v\n", vacancyID, err)
		return TranslateError(err)
	}
	return nil
}

func DeleteVacancy(vacancyID uint) error {
	err := db.GetDBConn().Model(&models.Vacancy{}).Where("id = ?", vacancyID).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteVacancy] Failed to soft delete vacancy with ID %v. Error: %v\n", vacancyID, err)
		return TranslateError(err)
	}
	return nil
}

func GetVacancyReport() ([]models.VacancyReport, error) {
	var reports []models.VacancyReport
	err := db.GetDBConn().
		Table("vacancies").
		Select("vacancies.id as vacancy_id, vacancies.title as vacancy_title, COUNT(applications.id) as application_count").
		Joins("left join applications on applications.vacancy_id = vacancies.id").
		Group("vacancies.id").
		Scan(&reports).Error
	if err != nil {
		logger.Error.Printf("[repository.GetVacancyReport] Error retrieving vacancy report: %v", err)
		return nil, err
	}
	return reports, nil
}
