package service

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"TajikCareerHub/utils/errs"
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
		if err := checkVacancyBlocked(vacancy.ID); err != nil {
			continue
		}
		if err := checkUserBlocked(vacancy.UserID); err != nil {
			continue
		}
		filteredVacancies = append(filteredVacancies, vacancy)
	}

	return filteredVacancies, nil
}

func GetVacancyByID(userID uint, vacancyID uint) (vacancy models.Vacancy, err error) {
	if err := checkUserBlocked(userID); err != nil {
		return models.Vacancy{}, err
	}

	vacancy, err = repository.GetVacancyByID(vacancyID)
	if err != nil {
		return models.Vacancy{}, err
	}

	if err := checkVacancyBlocked(vacancyID); err != nil {
		return models.Vacancy{}, err
	}

	if err := repository.RecordVacancyView(userID, vacancyID); err != nil {
		return models.Vacancy{}, err
	}
	return vacancy, nil
}

func AddVacancy(userID uint, vacancy models.Vacancy) (err error) {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	vacancy.UserID = userID
	if err := vacancy.ValidateVacancy(); err != nil {
		logger.Error.Printf("[service.AddVacancy] validation error: %v\n", err)
		return err
	}
	return repository.AddVacancy(vacancy)
}

func UpdateVacancy(userID uint, vacancyID uint, updatedVacancy models.Vacancy) (err error) {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	vacancy, err := repository.GetVacancyByID(vacancyID)
	if err != nil {
		return err
	}
	if err := checkVacancyBlocked(vacancyID); err != nil {
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

	err = vacancy.ValidateVacancy()
	if err != nil {
		logger.Error.Printf("[service.UpdateVacancy] validation error: %v\n", err)
		return err
	}
	err = repository.UpdateVacancy(vacancyID, vacancy)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVacancy(userID uint, vacancyID uint) (err error) {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	if err := checkVacancyBlocked(vacancyID); err != nil {
		return err
	}
	err = repository.DeleteVacancy(vacancyID)
	if err != nil {
		return err
	}
	return nil
}

func GetVacancyReportByID(vacancyID uint) (*models.VacancyReport, error) {
	err := checkVacancyBlocked(vacancyID)
	if err != nil {
		return nil, errs.ErrVacancyBlocked
	}
	report, err := repository.GetVacancyReportByID(vacancyID)
	if err != nil {
		return nil, err
	}

	if report == nil {
		return nil, errs.ErrNoReportsFound
	}

	return report, nil
}

func BlockVacancy(userID uint, vacancyID uint, RoleID uint) (err error) {

	if RoleID != 1 {
		return errs.ErrAccessDenied
	}
	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	if err := checkVacancyBlocked(vacancyID); err != nil {
		return err
	}
	err = repository.BlockVacancy(vacancyID)
	if err != nil {
		return err
	}
	return nil
}

func UnblockVacancy(userID uint, vacancyID uint, RoleID uint) (err error) {
	if RoleID != 1 {
		return errs.ErrAccessDenied
	}

	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	if err := checkVacancyBlocked(vacancyID); err != nil {
		return err
	}
	err = repository.UnblockVacancy(vacancyID)
	if err != nil {
		return err
	}
	return nil
}
