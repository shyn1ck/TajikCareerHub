package service

import (
	"TajikCareerHub/errs"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"errors"
)

func GetAllCategories() ([]models.VacancyCategory, error) {
	categories, err := repository.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(id uint) (models.VacancyCategory, error) {
	categoryList, err := repository.GetCategoryByID(id)
	if err != nil {
		return models.VacancyCategory{}, err
	}
	if len(categoryList) == 0 {
		return models.VacancyCategory{}, errors.New("category not found")
	}
	return categoryList[0], nil
}

func AddCategory(category models.VacancyCategory) error {
	existingCategory, err := repository.GetCategoryByID(category.ID)
	if err != nil {
		if !errors.Is(err, errs.ErrRecordNotFound) {
			return err
		}
	}

	if len(existingCategory) > 0 {
		return errors.New("category already exists")
	}

	return repository.AddCategory(category)
}

func UpdateCategory(category models.VacancyCategory) error {
	existingCategory, err := repository.GetCategoryByID(category.ID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errors.New("category does not exist")
		}
		return err
	}
	if len(existingCategory) == 0 {
		return errors.New("category does not exist")
	}

	return repository.UpdateCategory(category)
}

func DeleteCategory(id uint) error {
	existingCategory, err := repository.GetCategoryByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errors.New("category does not exist")
		}
		return err
	}
	if len(existingCategory) == 0 {
		return errors.New("category does not exist")
	}

	return repository.DeleteCategory(id)
}
