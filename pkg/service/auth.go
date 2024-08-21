package service

import (
	"TajikCareerHub/pkg/repository"
	"TajikCareerHub/utils"
)

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", err
	}

	accessToken, err = GenerateToken(user.ID, user.UserName)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
