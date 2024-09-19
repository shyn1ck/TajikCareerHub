package service

import (
	"TajikCareerHub/errs"
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

func UpdateApplicationStatus(applicationID uint, statusID uint) error {
	validStatusIDs := map[uint]bool{
		1: true, // Applied
		2: true, // Under Review
		3: true, // Rejected
		4: true, // Interview
	}
	if !validStatusIDs[statusID] {
		return errs.ErrIDIsNotCorrect
	}
	err := repository.UpdateApplicationStatus(applicationID, statusID)
	if err != nil {
		return err
	}
	return nil
}
