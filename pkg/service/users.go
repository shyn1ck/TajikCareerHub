package service

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"TajikCareerHub/utils"
	"errors"
	"gorm.io/gorm"
	"strings"
)

func validateUserName(username string) error {
	if strings.TrimSpace(username) == "" {
		return errors.New("username cannot be empty")
	}
	return nil
}

func validateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("email cannot be empty")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

func GetUserByID(id uint) (models.User, error) {
	return repository.GetUserByID(id)
}

func CreateUser(user models.User) error {
	if err := validateUserName(user.UserName); err != nil {
		return err
	}
	if err := validateEmail(user.Email); err != nil {
		return err
	}
	_, err := repository.GetUserByUserName(user.UserName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err := validatePassword(user.Password); err != nil {
		return err
	}
	user.Password = utils.GenerateHash(user.Password)
	return repository.CreateUser(user)
}

func UpdateUser(user models.User) error {
	_, err := repository.GetUserByID(user.ID)
	if err != nil {
		return err
	}
	if err := validateUserName(user.UserName); err != nil {
		return err
	}
	if err := validateEmail(user.Email); err != nil {
		return err
	}
	return repository.UpdateUser(user)
}

func DeleteUser(id uint) error {
	return repository.DeleteUser(id)
}

func UpdateUserPassword(id uint, newPassword string) error {
	logger.Info.Printf("[service.UpdateUserPassword] Received password: %s\n", newPassword)
	if len(newPassword) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	hashedPassword := utils.GenerateHash(newPassword)
	return repository.UpdateUserPassword(id, hashedPassword)
}

func CheckUserExists(username string, email string) (bool, error) {
	if err := validateUserName(username); err != nil {
		return false, err
	}
	if err := validateEmail(email); err != nil {
		return false, err
	}
	return repository.CheckUserExists(username, email)
}
