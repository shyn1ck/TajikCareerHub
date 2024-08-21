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
	return companies, nil
}

func GetCompanyByID(id uint) (company []models.Company, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&company).Error
	if err != nil {
		logger.Error.Printf("[repository.GetCompanyByID]: Error retrieving company with ID %v. Error: %v\n", id, err)
		return company, err
	}
	return company, nil
}

func AddCompany(company models.Company) error {
	err := db.GetDBConn().Create(&company).Error
	if err != nil {
		logger.Error.Printf("[repository.AddCompany]: Failed to add company. Error: %v\n", err)
		return err
	}
	return nil
}

func UpdateCompany(company models.Company) error {
	err := db.GetDBConn().Save(&company).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateCompany]: Failed to update company with ID %v. Error: %v\n", company.ID, err)
		return err
	}
	return nil
}

func DeleteCompany(id uint) error {
	err := db.GetDBConn().Model(&models.Company{}).Where("id = ?", id).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteCompany]: Failed to soft delete company with ID %v. Error: %v\n", id, err)
		return err
	}
	return nil
}
