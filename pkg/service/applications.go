package service

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"TajikCareerHub/utils/errs"
)

func GetAllApplications(userID uint) (applications []models.Application, err error) {
	err = checkUserBlocked(userID)
	if err != nil {
		logger.Error.Printf("[service.GetAllApplications] Error checking user blocked]")
		return nil, err
	}
	applications, err = repository.GetAllApplications()
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func GetApplicationByID(userID, id uint) (application models.Application, err error) {
	err = checkUserBlocked(userID)
	if err != nil {
		return models.Application{}, err
	}
	application, err = repository.GetApplicationByID(id)
	if err != nil {
		return application, err
	}
	return application, nil
}

func AddApplication(application models.Application) (err error) {
	err = checkUserBlocked(application.UserID)
	if err != nil {
		return err
	}
	err = repository.AddApplication(application)
	if err != nil {
		return err
	}
	return nil
}

func UpdateApplication(application models.Application) (err error) {
	err = checkUserBlocked(application.UserID)
	if err != nil {
		return err
	}

	err = repository.UpdateApplication(application.ID, application)
	if err != nil {
		return err
	}
	return nil
}

func DeleteApplication(id, userID uint) (err error) {
	err = checkUserBlocked(userID)
	if err != nil {
		return err
	}

	err = repository.DeleteApplication(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateApplicationStatus(applicationID uint, statusID uint, userID uint) (err error) {
	err = checkUserBlocked(userID)
	if err != nil {
		return err
	}

	validStatusIDs := map[uint]bool{
		1: true, // Applied
		2: true, // Under Review
		3: true, // Rejected
		4: true, // Interview
	}
	if !validStatusIDs[statusID] {
		return errs.ErrIDIsNotCorrect
	}

	err = repository.UpdateApplicationStatus(applicationID, statusID)
	if err != nil {
		return err
	}
	return nil
}
