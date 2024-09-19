package service

import (
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/pkg/repository"
	"TajikCareerHub/utils"
)

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", err
	}
	if err := checkUserBlocked(user.ID); err != nil {
		logger.Error.Printf("[service.SignIn]: Error user blocked")
		return "", errs.ErrUserBlocked
	}
	accessToken, err = GenerateToken(user.ID, user.UserName, user.Role)
	if err != nil {
		logger.Error.Printf("[service.SignIn]: Error generating access token")
		return "", err
	}
	return accessToken, nil
}

func checkUserBlocked(userID uint) (err error) {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		logger.Error.Printf("[service.checkUserBlocked] Error retrieving user with ID %d: %v\n", userID, err)
		return err
	}
	if user.IsBlocked {
		logger.Info.Printf("[service.checkUserBlocked] User with ID %d is blocked.\n", userID)
		return errs.ErrUserBlocked
	}
	return nil
}

func checkVacancyBlocked(vacancyID uint) (err error) {
	vacancy, err := repository.GetVacancyByID(vacancyID)
	if err != nil {
		logger.Error.Printf("[service.checkVacancyBlocked] Error retrieving vacancy with ID %d: %v\n", vacancyID, err)
		return err
	}
	if vacancy.IsBlocked {
		logger.Info.Printf("[service.checkVacancyBlocked] Vacancy with ID %d is blocked.\n", vacancyID)
		return errs.ErrVacancyBlocked
	}
	return nil
}

func checkResumeBlocked(resumeID uint) (err error) {
	resume, err := repository.GetResumeByID(resumeID)
	if err != nil {
		logger.Error.Printf("[service.checkResumeBlocked] Error retrieving resume with ID %d: %v\n", resumeID, err)
		return err
	}
	if resume.IsBlocked {
		logger.Info.Printf("[service.checkResumeBlocked] Resume with ID %d is blocked.\n", resumeID)
		return errs.ErrResumeBlocked
	}
	return nil
}
