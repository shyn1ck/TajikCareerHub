package service

import (
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
)

func GetAllResumes(search string, minExperienceYears int, location string, category string, userID uint) ([]models.Resume, error) {
	if err := checkUserBlocked(userID); err != nil {
		return nil, err
	}
	resumes, err := repository.GetAllResumes(search, minExperienceYears, location, category)
	if err != nil {
		return nil, err
	}
	var filteredResumes []models.Resume
	for _, resume := range resumes {
		if err := checkResumeBlocked(resume.ID); err != nil {
			continue
		}
		filteredResumes = append(filteredResumes, resume)
	}

	return filteredResumes, nil
}

func GetResumeByID(id uint, userID uint) (models.Resume, error) {
	if err := checkUserBlocked(userID); err != nil {
		return models.Resume{}, err
	}
	resume, err := repository.GetResumeByID(id)
	if err != nil {
		return models.Resume{}, err
	}
	if err := checkResumeBlocked(resume.ID); err != nil {
		return models.Resume{}, err
	}

	return resume, nil
}

func AddResume(resume models.Resume, userID uint) error {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}

	if resume.UserID == 0 {
		return errs.ErrIDIsNotCorrect
	}
	if err := resume.ValidateResume(); err != nil {
		logger.Error.Printf("[service.AddResume] validation error: %v\n", err)
		return err
	}
	return repository.AddResume(resume)
}

func UpdateResume(resumeID uint, updatedResume models.Resume, userID uint) error {
	if err := checkUserBlocked(userID); err != nil {
		return err
	}
	resume, err := repository.GetResumeByID(resumeID)
	if err != nil {
		return err
	}
	if err := checkResumeBlocked(resume.ID); err != nil {
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
	if err := checkResumeBlocked(id); err != nil {
		return err
	}

	return repository.DeleteResume(id)
}
