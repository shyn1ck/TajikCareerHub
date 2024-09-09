package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetVacancyCategoryByID(id uint) (category models.VacancyCategory, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.GetVacancyCategoryByID]: Error retrieving vacancy category with ID %v. Error: %v\n", id, err)
		return category, err
	}
	return category, nil
}

func GetAllVacancyCategories() (categories []models.VacancyCategory, err error) {
	err = db.GetDBConn().Where("deleted_at = ?", false).Find(&categories).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllVacancyCategories]: Error retrieving all vacancy categories. Error: %v\n", err)
		return nil, err
	}
	return categories, nil
}

func CreateVacancyCategory(category models.VacancyCategory) error {
	err := db.GetDBConn().Create(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.CreateVacancyCategory]: Failed to create vacancy category. Error: %v\n", err)
		return err
	}
	return nil
}

func UpdateVacancyCategory(category models.VacancyCategory) error {
	err := db.GetDBConn().Model(&models.VacancyCategory{}).Where("id = ?", category.ID).Updates(category).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateVacancyCategory]: Failed to update vacancy category with ID %v. Error: %v\n", category.ID, err)
		return err
	}
	return nil
}

func DeleteVacancyCategory(id uint) error {
	err := db.GetDBConn().Model(&models.VacancyCategory{}).Where("id = ?", id).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteVacancyCategory]: Failed to soft delete vacancy category with ID %v. Error: %v\n", id, err)
		return err
	}
	return nil
}
