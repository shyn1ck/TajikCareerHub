package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllCompanies() (companies []models.Company, err error) {
	err = db.GetDBConn().Where("deleted_at = ?", false).Find(&companies).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllCompanies]: Error retrieving all companies. Error: %v\n", err)
		return nil, err
	}
	logger.Info.Println("[repository.GetAllCompanies]: Successfully retrieved all companies.")
	return companies, nil
}

func GetCompanyByID(id uint) (company []models.Company, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&company).Error
	if err != nil {
		logger.Error.Printf("[repository.GetCompanyByID]: Error retrieving company with ID %v. Error: %v\n", id, err)
		return company, err
	}
	logger.Info.Printf("[repository.GetCompanyByID]: Successfully retrieved company with ID %v.\n", id)
	return company, nil
}

func AddCompany(company models.Company) error {
	logger.Info.Printf("[repository.AddCompany]: Adding new company %v...\n", company.Name)
	err := db.GetDBConn().Create(&company).Error
	if err != nil {
		logger.Error.Printf("[repository.AddCompany]: Failed to add company. Error: %v\n", err)
		return err
	}
	logger.Info.Printf("[repository.AddCompany]: Successfully added company %v.\n", company.Name)
	return nil
}

func UpdateCompany(company models.Company) error {
	logger.Info.Printf("[repository.UpdateCompany]: Updating company with ID %v...\n", company.ID)
	err := db.GetDBConn().Save(&company).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateCompany]: Failed to update company with ID %v. Error: %v\n", company.ID, err)
		return err
	}
	logger.Info.Printf("[repository.UpdateCompany]: Successfully updated company with ID %v.\n", company.ID)
	return nil
}

func DeleteCompany(id uint) error {
	logger.Info.Printf("[repository.DeleteCompany]: Soft deleting company with ID %v...\n", id)
	err := db.GetDBConn().Model(&models.Company{}).Where("id = ?", id).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteCompany]: Failed to soft delete company with ID %v. Error: %v\n", id, err)
		return err
	}
	logger.Info.Printf("[repository.DeleteCompany]: Successfully soft deleted company with ID %v.\n", id)
	return nil
}
