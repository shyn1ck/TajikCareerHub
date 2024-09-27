package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"TajikCareerHub/utils/errs"
	"errors"
	"gorm.io/gorm"
)

func GetAllCategories() (categories []models.VacancyCategory, err error) {
	categories, err = repository.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(id uint) (category models.VacancyCategory, err error) {
	category, err = repository.GetCategoryByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return models.VacancyCategory{}, nil
		}
		return models.VacancyCategory{}, err
	}
	return category, nil
}

func AddCategory(category models.VacancyCategory, RoleID uint) (err error) {
	if RoleID != 1 {
		return errs.ErrPermissionDenied
	}
	existingCategory, err := repository.GetCategoryByName(category.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		} else {
			return err
		}
	}
	if existingCategory.ID != 0 {
		return errs.ErrCategoryAlreadyExist
	}
	return repository.AddCategory(category)
}

func UpdateCategory(category models.VacancyCategory, RoleID uint) (err error) {
	if RoleID != 1 {
		return errs.ErrPermissionDenied
	}
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
	err = repository.UpdateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(id uint, RoleID uint) (err error) {
	if RoleID != 1 {
		return errs.ErrPermissionDenied
	}
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
	err = repository.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}
