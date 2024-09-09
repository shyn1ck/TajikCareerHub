package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"errors"
)

func GetAllResume(search string, minExperienceYears int, location string, category string) (resumes []models.Resume, err error) {
	resumes, err = repository.GetAllResumes(search, minExperienceYears, location, category)
	if err != nil {
		return nil, err
	}
	return resumes, nil
}

func GetResumeByID(id uint) (models.Resume, error) {
	return repository.GetResumeByID(id)
}

func AddResume(resume models.Resume) error {
	if resume.UserID == 0 {
		return errors.New("user_id must be provided")
	}
	err := repository.AddResume(resume)
	if err != nil {
		return err
	}
	return nil
}

func UpdateResume(resumeID uint, updatedResume models.Resume) error {
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

func DeleteResume(id uint) error {
	return repository.DeleteResume(id)
}
