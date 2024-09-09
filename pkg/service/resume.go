package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"errors"
)

func GetAllResumes(search string, minExperienceYears int, location string, category string, userID uint) (resumes []models.Resume, err error) {
	if err := checkUserBlocked(userID); err != nil {
		return nil, err
	}

	resumes, err = repository.GetAllResumes(search, minExperienceYears, location, category)
	if err != nil {
		return nil, err
	}
	return resumes, nil
}

func GetResumeByID(id uint, userID uint) (models.Resume, error) {
	if err := checkUserBlocked(userID); err != nil {
		return models.Resume{}, err
	}

	return repository.GetResumeByID(id)
}

func AddResume(resume models.Resume, userID uint) error {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}

	if resume.UserID == 0 {
		return errors.New("user_id must be provided")
	}
	err := repository.AddResume(resume)
	if err != nil {
		return err
	}
	return nil
}

func UpdateResume(resumeID uint, updatedResume models.Resume, userID uint) error {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	resume, err := repository.GetResumeByID(resumeID)
	if err != nil {
		return err
	}
	if updatedResume.FullName != "" {
		resume.FullName = updatedResume.FullName
	}
	if updatedResume.Summary != "" {
		resume.Summary = updatedResume.Summary
	}
	if updatedResume.Skills != "" {
		resume.Skills = updatedResume.Skills
	}
	if updatedResume.ExperienceYears != 0 {
		resume.ExperienceYears = updatedResume.ExperienceYears
	}
	if updatedResume.Education != "" {
		resume.Education = updatedResume.Education
	}
	if updatedResume.Certifications != "" {
		resume.Certifications = updatedResume.Certifications
	}
	if updatedResume.Location != "" {
		resume.Location = updatedResume.Location
	}
	if updatedResume.VacancyCategoryID != 0 {
		resume.VacancyCategoryID = updatedResume.VacancyCategoryID
	}

	return repository.UpdateResume(resumeID, resume)
}

func DeleteResume(id uint, userID uint) error {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	return repository.DeleteResume(id)
}
