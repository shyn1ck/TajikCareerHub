package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"TajikCareerHub/utils"
	"errors"
	"gorm.io/gorm"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) error {
	_, err := repository.GetUserByUserName(user.UserName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	user.Password = utils.GenerateHash(user.Password)
	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user models.User) error {
	err := repository.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	err := repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserPassword(id uint, newPassword string) error {
	hashedPassword := utils.GenerateHash(newPassword)
	err := repository.UpdateUserPassword(id, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func CheckUserExists(username, email string) (bool, error) {
	exists, err := repository.CheckUserExists(username, email)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func GetUserByUsername(username string) (models.User, error) {
	user, err := repository.GetUserByUserName(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	user, err := repository.GetUserByUserName(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, err
	}
	return user, nil
}
