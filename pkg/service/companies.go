package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"errors"
)

func GetAllCompanies() ([]models.Company, error) {
	companies, err := repository.GetAllCompanies()
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func GetCompanyByID(id uint) (models.Company, error) {
	companies, err := repository.GetCompanyByID(id)
	if err != nil {
		return models.Company{}, err
	}
	if len(companies) == 0 {
		return models.Company{}, errors.New("company not found")
	}
	return companies[0], nil
}

func AddCompany(company models.Company) error {
	return repository.AddCompany(company)
}

func UpdateCompany(company models.Company) error {
	return repository.UpdateCompany(company)
}

func DeleteCompany(id uint) error {
	return repository.DeleteCompany(id)
}
