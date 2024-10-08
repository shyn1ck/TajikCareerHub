package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"TajikCareerHub/utils/errs"
)

func GetAllCompanies(userID uint) (companies []models.Company, err error) {
	err = checkUserBlocked(userID)
	if err != nil {
		return companies, err
	}
	companies, err = repository.GetAllCompanies()
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func GetCompanyByID(id uint, userID uint) (company models.Company, err error) {
	err = checkUserBlocked(userID)
	if err != nil {
		return company, err
	}
	company, err = repository.GetCompanyByID(id)
	if err != nil {
		return models.Company{}, err
	}
	if (company == models.Company{}) {
		return models.Company{}, errs.ErrCompanyNotFound
	}
	return company, nil
}

func AddCompany(userID uint, company models.Company, RoleID uint) (err error) {
	err = checkUserBlocked(userID)
	if err != nil {
		return err
	}

	if RoleID != 2 && RoleID != 1 {
		return errs.ErrAccessDenied
	}

	err = repository.AddCompany(company)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCompany(userID uint, company models.Company, RoleID uint) (err error) {
	err = checkUserBlocked(userID)
	if err != nil {
		return err
	}

	if RoleID != 2 && RoleID != 1 {
		return errs.ErrAccessDenied
	}

	err = repository.UpdateCompany(company)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCompany(id uint, userID uint, RoleID uint) (err error) {
	err = checkUserBlocked(userID)
	if err != nil {
		return err
	}

	if RoleID != 2 && RoleID != 1 {
		return errs.ErrAccessDenied
	}

	err = repository.DeleteCompany(id)
	if err != nil {
		return err
	}
	return nil
}
