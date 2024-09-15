package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"errors"
	"gorm.io/gorm"
)

func GetAllCategories() ([]models.VacancyCategory, error) {
	var categories []models.VacancyCategory
	err := db.GetDBConn().
		Where("deleted_at = ?", false).
		Find(&categories).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllCategories]: Error retrieving all Categories. Error: %v\n", err)
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(id uint) (models.VacancyCategory, error) {
	var category models.VacancyCategory
	err := db.GetDBConn().
		Where("id = ? AND deleted_at = false", id).
		First(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.GetCategoryByID]: Error retrieving category with ID %v. Error: %v\n", id, err)
		return models.VacancyCategory{}, TranslateError(err)
	}
	return category, nil
}

func AddCategory(category models.VacancyCategory) error {
	err := db.GetDBConn().Create(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.AddCategory]: Failed to add category. Error: %v\n", err)
		return TranslateError(err)
	}
	return nil
}

func UpdateCategory(category models.VacancyCategory) error {
	err := db.GetDBConn().
		Model(&models.VacancyCategory{}).
		Where("id = ? AND deleted_at = false", category.ID).
		Save(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateCategory]: Failed to update category with ID %v. Error: %v\n", category.ID, err)
		return TranslateError(err)
	}
	return nil
}

func DeleteCategory(id uint) error {
	err := db.GetDBConn().
		Model(&models.VacancyCategory{}).
		Where("id = ?", id).
		Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteCategory]: Failed to soft delete category with ID %v. Error: %v\n", id, err)
		return TranslateError(err)
	}
	return nil
}

func GetCategoryByName(categoryName string) (category models.VacancyCategory, err error) {
	err = db.GetDBConn().Where("name = ? AND deleted_at = false", categoryName).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return category, nil
		}
		logger.Error.Printf("[repository.GetCategoryByName]: Error retrieving category by name. Error: %v\n", err)
		return category, TranslateError(err)
	}
	return category, nil
}
