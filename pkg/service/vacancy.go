package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"errors"
)

func GetAllVacancies(search string, minSalary int, maxSalary int, location string, category string, sort string) (vacancies []models.Vacancy, err error) {
	vacancies, err = repository.GetAllVacancies(search, minSalary, maxSalary, location, category, sort)
	if err != nil {
		return nil, err
	}

	return vacancies, nil
}

func GetVacancyByID(id uint) (models.Vacancy, error) {
	return repository.GetVacancyByID(id)
}

func AddVacancy(vacancy models.Vacancy) error {
	if vacancy.UserID == 0 {
		return errors.New("user_id must be provided")
	}
	err := repository.AddVacancy(vacancy)
	if err != nil {
		return err
	}
	return nil
}

func UpdateVacancy(vacancyID uint, updatedVacancy models.Vacancy) error {
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

func DeleteVacancy(vacancyID uint) error {
	return repository.DeleteVacancy(vacancyID)
}
