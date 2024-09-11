package service

import (
	"TajikCareerHub/errs"
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
		return "", errs.ErrUserBlocked
	}
	accessToken, err = GenerateToken(user.ID, user.UserName, user.Role)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func checkUserBlocked(userID uint) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user.IsBlocked {
		return errs.ErrUserBlocked
	}
	return nil
}

func checkVacancyBlocked(vacancyID uint) error {
	vacancy, err := repository.GetVacancyByID(vacancyID)
	if err != nil {
		return err
	}
	if vacancy.IsBlocked {
		return errs.ErrVacancyBlocked
	}
	return nil
}

func checkResumeBlocked(resumeID uint) error {
	resume, err := repository.GetResumeByID(resumeID)
	if err != nil {
		return err
	}
	if resume.IsBlocked {
		return errs.ErrResumeBlocked
	}
	return nil
}
