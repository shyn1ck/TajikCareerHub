package service

import (
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
)

func GetAllApplications(userID uint) ([]models.Application, error) {
	err := checkUserBlocked(userID)
	if err != nil {
		return nil, err
	}
	applications, err := repository.GetAllApplications()
	if err != nil {
		return nil, err
	}

	return applications, nil
}

func GetApplicationByID(userID, id uint) (models.Application, error) {
	err := checkUserBlocked(userID)
	if err != nil {
		return models.Application{}, errs.ErrUserBlocked
	}
	application, err := repository.GetApplicationByID(id)
	if err != nil {
		return application, err
	}
	return application, nil
}

func AddApplication(application models.Application) error {
	err := checkUserBlocked(application.UserID)
	if err != nil {
		return errs.ErrUserBlocked
	}
	err = repository.AddApplication(application)
	if err != nil {
		return err
		logger.Error.Printf("[repository.AddApplication]: Error adding application to database", err)
	}
	return nil
}

func UpdateApplication(application models.Application) error {
	err := checkUserBlocked(application.UserID)
	if err != nil {
		return errs.ErrUserBlocked
	}
	err = repository.UpdateApplication(application.ID, application)
	if err != nil {
		return err
	}
	return nil
}

func DeleteApplication(id, userID uint) error {
	err := checkUserBlocked(userID)
	if err != nil {
		return errs.ErrUserBlocked
	}
	err = repository.DeleteApplication(id)
	if err != nil {
		return err
	}
	return nil
}

func GetSpecialistActivityReport(userID uint) ([]models.SpecialistActivityReport, error) {
	err := checkUserBlocked(userID)
	if err != nil {
		return nil, errs.ErrUserBlocked
	}
	reports, err := repository.GetSpecialistActivityReport()
	if err != nil {
		return nil, err
	}
	return reports, nil
}
