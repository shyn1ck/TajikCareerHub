package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"errors"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := db.GetDBConn().Where("deleted_at = false").Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %s\n", err.Error())
		return nil, TranslateError(err)
	}

	return users, nil
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := db.GetDBConn().Where("id = ? AND deleted_at = false", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		return models.User{}, TranslateError(err)
	}
	return user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.GetDBConn().Where("user_name = ? AND deleted_at = false", username).First(&user).Error

	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return nil, TranslateError(err)
	}
	return &user, nil
}

func UserExists(username, email string) (bool, bool, error) {
	var usernameExists, emailExists bool

	var user models.User

	// Check username existence
	err := db.GetDBConn().Where("user_name = ? AND deleted_at = false", username).First(&user).Error
	if err == nil {
		usernameExists = true
	} else if !errors.Is(err, errs.ErrRecordNotFound) {
		return false, false, TranslateError(err)
	}

	// Check email existence
	err = db.GetDBConn().Where("email = ? AND deleted_at = false", email).First(&user).Error
	if err == nil {
		emailExists = true
	} else if !errors.Is(err, errs.ErrRecordNotFound) {
		return false, false, TranslateError(err)
	}

	return usernameExists, emailExists, nil
}

func CreateUser(user models.User) (uint, error) {
	if err := db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return 0, TranslateError(err)
	}
	return user.ID, nil
}

func GetUserByUsernameAndPassword(username, password string) (models.User, error) {
	var user models.User
	err := db.GetDBConn().Where("user_name = ? AND password = ? AND deleted_at = false", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return models.User{}, TranslateError(err)
	}

	return user, nil
}

func UpdateUser(user models.User) error {
	updateData := map[string]interface{}{
		"full_name":  user.FullName,
		"user_name":  user.UserName,
		"birth_date": user.BirthDate,
		"email":      user.Email,
		"password":   user.Password,
	}

	// Remove empty fields
	for k, v := range updateData {
		if v == "" {
			delete(updateData, k)
		}
	}

	if len(updateData) == 0 {
		return nil
	}

	err := db.GetDBConn().
		Model(&models.User{}).
		Where("id = ? AND deleted_at = false", user.ID).
		Updates(updateData).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUser] Failed to update user with ID %v: %v\n", user.ID, err)
		return TranslateError(err)
	}
	return nil
}

func DeleteUser(id uint) error {
	err := db.GetDBConn().
		Model(&models.User{}).
		Where("id = ? AND deleted_at = false", id).
		Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteUser] Failed to soft delete user with ID %v: %v\n", id, err)
		return TranslateError(err)
	}
	return nil
}

func UpdateUserPassword(id uint, newPassword string) error {
	logger.Info.Printf("[repository.UpdateUserPassword] Updating password for user ID: %d with new hashed password\n", id)
	err := db.GetDBConn().
		Model(&models.User{}).
		Where("id = ? AND deleted_at = false", id).
		Update("password", newPassword).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserPassword] Failed to update password for user with ID %v: %v\n", id, err)
		return TranslateError(err)
	}
	return nil
}
