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
