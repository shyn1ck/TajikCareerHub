package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
)

func UpdateApplicationStatus(statusUpdate models.ApplicationStatusUpdate) error {

	application, err := repository.GetApplicationByID(statusUpdate.ApplicationID)
	if err != nil {
		return err
	}

	if err := checkUserBlocked(application.UserID); err != nil {
		return err
	}

	if err := checkResumeBlocked(application.ResumeID); err != nil {
		return err
	}

	if err := checkVacancyBlocked(application.VacancyID); err != nil {
		return err
	}

	application.Status = statusUpdate.Status
	return repository.UpdateApplication(application)
}

func GetAllApplications() ([]models.Application, error) {
	applications, err := repository.GetAllApplications()
	if err != nil {
		return nil, err
	}
	for _, app := range applications {
		if err := checkUserBlocked(app.UserID); err != nil {
			return nil, err
		}
		if err := checkResumeBlocked(app.ResumeID); err != nil {
			return nil, err
		}
		if err := checkVacancyBlocked(app.VacancyID); err != nil {
			return nil, err
		}
	}
	return applications, nil
}

func GetApplicationByID(applicationID uint) (models.Application, error) {
	application, err := repository.GetApplicationByID(applicationID)
	if err != nil {
		return models.Application{}, err
	}

	if err := checkUserBlocked(application.UserID); err != nil {
		return models.Application{}, err
	}

	if err := checkResumeBlocked(application.ResumeID); err != nil {
		return models.Application{}, err
	}

	if err := checkVacancyBlocked(application.VacancyID); err != nil {
		return models.Application{}, err
	}

	return application, nil
}

func GetApplicationsByUserID(userID uint) ([]models.Application, error) {
	if err := checkUserBlocked(userID); err != nil {
		return nil, err
	}

	applications, err := repository.GetApplicationsByUserID(userID)
	if err != nil {
		return nil, err
	}

	for _, app := range applications {
		if err := checkResumeBlocked(app.ResumeID); err != nil {
			return nil, err
		}
		if err := checkVacancyBlocked(app.VacancyID); err != nil {
			return nil, err
		}
	}

	return applications, nil
}

func GetApplicationsByVacancyID(vacancyID uint) ([]models.Application, error) {
	if err := checkVacancyBlocked(vacancyID); err != nil {
		return nil, err
	}

	applications, err := repository.GetApplicationsByVacancyID(vacancyID)
	if err != nil {
		return nil, err
	}

	for _, app := range applications {
		if err := checkUserBlocked(app.UserID); err != nil {
			return nil, err
		}
		if err := checkResumeBlocked(app.ResumeID); err != nil {
			return nil, err
		}
	}

	return applications, nil
}

func ApplyForVacancy(userID, vacancyID, resumeID uint) error {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	if err := checkResumeBlocked(resumeID); err != nil {
		return err
	}
	if err := checkVacancyBlocked(vacancyID); err != nil {
		return err
	}

	application := models.Application{
		UserID:    userID,
		VacancyID: vacancyID,
		ResumeID:  resumeID,
		Status:    "pending",
	}
	return repository.AddApplication(application)
}

func UpdateApplication(application models.Application) error {
	if err := checkUserBlocked(application.UserID); err != nil {
		return err
	}
	if err := checkResumeBlocked(application.ResumeID); err != nil {
		return err
	}
	if err := checkVacancyBlocked(application.VacancyID); err != nil {
		return err
	}

	return repository.UpdateApplication(application)
}

func DeleteApplication(applicationID uint) error {
	return repository.DeleteApplication(applicationID)
}
