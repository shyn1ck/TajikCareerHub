package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
)

func GetAllApplications() ([]models.Application, error) {
	applications, err := repository.GetAllApplications()
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func GetApplicationByID(id uint) (models.Application, error) {
	application, err := repository.GetApplicationByID(id)
	if err != nil {
		return application, err
	}
	return application, nil
}

func GetApplicationsByUserID(userID uint) ([]models.Application, error) {
	applications, err := repository.GetApplicationsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func GetApplicationsByJobID(jobID uint) ([]models.Application, error) {
	applications, err := repository.GetApplicationsByJobID(jobID)
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func ApplyForVacancy(userID uint, vacancyID uint, resumeID uint) error {
	_, err := repository.GetVacancyByID(vacancyID)
	if err != nil {
		return err
	}
	resume, err := repository.GetResumeByID(resumeID)
	if err != nil {
		return err
	}
	application := models.Application{
		UserID:    userID,
		VacancyID: vacancyID,
		Resume:    resume,
		Status:    "pending",
	}
	err = repository.AddApplication(application)
	if err != nil {
		return err
	}
	return nil
}

func UpdateApplication(application models.Application) error {
	err := repository.UpdateApplication(application)
	if err != nil {
		return err
	}
	return nil
}

func DeleteApplication(id uint) error {
	err := repository.DeleteApplication(id)
	if err != nil {
		return err
	}
	return nil
}
