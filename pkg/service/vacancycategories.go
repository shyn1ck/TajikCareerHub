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
	category, err := repository.GetCategoryByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return models.VacancyCategory{}, nil
		}
		return models.VacancyCategory{}, err
	}
	return category, nil
}

func AddCategory(category models.VacancyCategory) error {
	existingCategory, err := repository.GetCategoryByID(category.ID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return repository.AddCategory(category)
		}
		return err
	}
	if existingCategory.ID != 0 {
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
	if existingCategory.ID == 0 {
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
	if existingCategory.ID == 0 {
		return errors.New("category does not exist")
	}
	return repository.DeleteCategory(id)
}
