package service

import (
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
	"TajikCareerHub/utils"
	"errors"
	"strings"
)

func validateUserCredentials(username, email, password string) error {
	if strings.TrimSpace(username) == "" {
		return errors.New("username cannot be empty")
	}
	if strings.TrimSpace(email) == "" {
		return errors.New("email cannot be empty")
	}
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id uint) (models.User, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		logger.Error.Printf("[service.GetUserByID] error retrieving user by ID: %v\n", err)
		return models.User{}, err
	}
	return user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		logger.Error.Printf("[service.GetUserByUsername] error retrieving user by username: %v\n", err)
		return nil, err
	}
	return user, nil
}

func CheckUserExists(username, email string) (bool, bool, error) {
	usernameExists, emailExists, err := repository.UserExists(username, email)
	if err != nil {
		logger.Error.Printf("[service.CheckUserExists] error checking user existence: %v\n", err)
		return false, false, err
	}
	return usernameExists, emailExists, nil
}

func CreateUser(user models.User) (uint, error) {
	if err := validateUserCredentials(user.UserName, user.Email, user.Password); err != nil {
		logger.Error.Printf("[service.CreateUser] validation error: %v\n", err)
		return 0, err
	}
	user.Password = utils.GenerateHash(user.Password)
	id, err := repository.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateUser(user models.User) error {
	existingUser, err := repository.GetUserByID(user.ID)
	if err != nil {
		logger.Error.Printf("[service.UpdateUser] Failed to get existing user with ID %v: %v\n", user.ID, err)
		return err
	}

	if user.FullName != "" {
		existingUser.FullName = user.FullName
	}
	if user.UserName != "" {
		existingUser.UserName = user.UserName
	}
	if !user.BirthDate.IsZero() {
		existingUser.BirthDate = user.BirthDate
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}

	err = repository.UpdateUser(existingUser)
	if err != nil {
		logger.Error.Printf("[service.UpdateUser] Failed to update user with ID %v: %v\n", user.ID, err)
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("invalid user ID")
	}
	return repository.DeleteUser(id)
}

func UpdateUserPassword(id uint, newPassword string) error {
	logger.Info.Printf("[service.UpdateUserPassword] Updating password for user ID: %d\n", id)
	if len(newPassword) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	hashedPassword := utils.GenerateHash(newPassword)
	return repository.UpdateUserPassword(id, hashedPassword)
}

func BlockUser(id uint) error {
	if id == 0 {
		logger.Error.Printf("[service.BlockUser] Invalid ID: %v", id)
		return errs.ErrIDIsNotCorrect
	}
	err := repository.BlockUser(id)
	if err != nil {
		logger.Error.Printf("[service.BlockUser] Failed to block user with ID %v: %v", id, err)
	}
	return err
}

func UnblockUser(id uint) error {
	if id == 0 {
		logger.Error.Printf("[service.UnblockUser] Invalid ID: %v", id)
		return errs.ErrIDIsNotCorrect
	}
	err := repository.UnBlockUser(id)
	if err != nil {
		logger.Error.Printf("[service.UnblockUser] Failed to unblock user with ID %v: %v", id, err)
	}
	return err
}
