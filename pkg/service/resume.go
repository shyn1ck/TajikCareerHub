package service

import (
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
)

func GetAllResume(keyword, location, category string, minExperienceYears, maxExperienceYears uint) ([]models.Resume, error) {
	resumes, err := repository.GetAllResumes(keyword, location, category, minExperienceYears, maxExperienceYears)
	if err != nil {
		return nil, err
	}
	return resumes, nil
}

func GetResumeByID(id uint) (models.Resume, error) {
	return repository.GetResumeByID(id)
}

func AddResume(resume models.Resume) error {
	return repository.AddResume(resume)
}

func UpdateResume(resume models.Resume) error {
	existingResume, err := repository.GetResumeByID(resume.ID)
	if err != nil {
		logger.Error.Printf("[service.UpdateResume] Failed to get existing resume with ID %v: %v\n", resume.ID, err)
		return err
	}

	if resume.FullName != "" {
		existingResume.FullName = resume.FullName
	}
	if resume.Summary != "" {
		existingResume.Summary = resume.Summary
	}
	if resume.Skills != "" {
		existingResume.Skills = resume.Skills
	}
	if resume.ExperienceYears > 0 {
		existingResume.ExperienceYears = resume.ExperienceYears
	}
	if resume.Education != "" {
		existingResume.Education = resume.Education
	}
	if resume.Certifications != "" {
		existingResume.Certifications = resume.Certifications
	}
	if resume.Location != "" {
		existingResume.Location = resume.Location
	}
	if resume.JobCategoryID > 0 {
		existingResume.JobCategoryID = resume.JobCategoryID
	}

	err = repository.UpdateResume(existingResume)
	if err != nil {
		logger.Error.Printf("[service.UpdateResume] Failed to update resume with ID %v: %v\n", resume.ID, err)
		return errs.TranslateError(err)
	}
	return nil
}

func DeleteResume(id uint) error {
	return repository.DeleteResume(id)
}
