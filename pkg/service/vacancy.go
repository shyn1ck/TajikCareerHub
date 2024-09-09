package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
)

func GetAllVacancies(userID uint, search string, minSalary int, maxSalary int, location string, category string, sort string) ([]models.Vacancy, error) {
	if err := checkUserBlocked(userID); err != nil {
		return nil, err
	}
	vacancies, err := repository.GetAllVacancies(search, minSalary, maxSalary, location, category, sort)
	if err != nil {
		return nil, err
	}
	var filteredVacancies []models.Vacancy
	for _, vacancy := range vacancies {
		if err := checkUserBlocked(vacancy.UserID); err != nil {
			continue
		}
		filteredVacancies = append(filteredVacancies, vacancy)
	}
	return filteredVacancies, nil
}

func GetVacancyByID(userID uint, vacancyID uint) (models.Vacancy, error) {
	if err := checkUserBlocked(userID); err != nil {
		return models.Vacancy{}, err
	}

	vacancy, err := repository.GetVacancyByID(vacancyID)
	if err != nil {
		return models.Vacancy{}, err
	}

	return vacancy, nil
}

func AddVacancy(userID uint, vacancy models.Vacancy) error {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	vacancy.UserID = userID
	err := repository.AddVacancy(vacancy)
	if err != nil {
		return err
	}

	return nil
}

func UpdateVacancy(userID uint, vacancyID uint, updatedVacancy models.Vacancy) error {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}

	vacancy, err := repository.GetVacancyByID(vacancyID)
	if err != nil {
		return err
	}
	if updatedVacancy.Title != "" {
		vacancy.Title = updatedVacancy.Title
	}
	if updatedVacancy.Description != "" {
		vacancy.Description = updatedVacancy.Description
	}
	if updatedVacancy.Location != "" {
		vacancy.Location = updatedVacancy.Location
	}
	if updatedVacancy.VacancyCategoryID != 0 {
		vacancy.VacancyCategoryID = updatedVacancy.VacancyCategoryID
	}
	if updatedVacancy.Salary != 0 {
		vacancy.Salary = updatedVacancy.Salary
	}
	return repository.UpdateVacancy(vacancyID, vacancy)
}

func DeleteVacancy(userID uint, vacancyID uint) error {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	return repository.DeleteVacancy(vacancyID)
}
