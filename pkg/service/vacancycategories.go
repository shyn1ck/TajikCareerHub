package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"errors"
)

func GetVacancyCategoryByID(id uint) (models.VacancyCategory, error) {
	category, err := repository.GetVacancyCategoryByID(id)
	if err != nil {
		return category, err
	}
	return category, nil
}

func GetAllVacancyCategories() ([]models.VacancyCategory, error) {
	categories, err := repository.GetAllVacancyCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateVacancyCategory(category models.VacancyCategory) error {
	existingCategory, err := repository.GetVacancyCategoryByID(category.ID)
	if err == nil && existingCategory.ID != 0 {
		return errors.New("vacancy category already exists")
	}
	return repository.CreateVacancyCategory(category)
}

func UpdateVacancyCategory(category models.VacancyCategory) error {
	_, err := repository.GetVacancyCategoryByID(category.ID)
	if err != nil {
		return errors.New("vacancy category does not exist")
	}
	return repository.UpdateVacancyCategory(category)
}

func DeleteVacancyCategory(id uint) error {
	_, err := repository.GetVacancyCategoryByID(id)
	if err != nil {
		return errors.New("vacancy category does not exist")
	}
	return repository.DeleteVacancyCategory(id)
}
