package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"errors"
)

func GetJobCategoryByID(id uint) (models.JobCategory, error) {
	category, err := repository.GetJobCategoryByID(id)
	if err != nil {
		return category, err
	}
	return category, nil
}

func GetAllJobCategories() ([]models.JobCategory, error) {
	categories, err := repository.GetAllJobCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateJobCategory(category models.JobCategory) error {
	// Check if the category already exists
	existingCategory, err := repository.GetJobCategoryByID(category.ID)
	if err == nil && existingCategory.ID != 0 {
		return errors.New("job category already exists")
	}
	return repository.CreateJobCategory(category)
}

func UpdateJobCategory(category models.JobCategory) error {
	_, err := repository.GetJobCategoryByID(category.ID)
	if err != nil {
		return errors.New("job category does not exist")
	}
	return repository.UpdateJobCategory(category)
}

func DeleteJobCategory(id uint) error {
	_, err := repository.GetJobCategoryByID(id)
	if err != nil {
		return errors.New("job category does not exist")
	}
	return repository.DeleteJobCategory(id)
}
