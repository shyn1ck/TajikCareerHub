package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %s\n", err.Error())
		return nil, errs.TranslateError(err)
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		return user, errs.TranslateError(err)
	}
	return user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.GetDBConn().Where("user_name = ?", username).First(&user).Error

	if err != nil {
		if err == errs.ErrRecordNotFound {
			return nil, nil
		}
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return nil, errs.TranslateError(err)
	}
	return &user, nil
}

func UserExists(username, email string) (bool, bool, error) {
	users, err := GetAllUsers()
	if err != nil {
		return false, false, err
	}

	var usernameExists, emailExists bool
	for _, user := range users {
		if user.UserName == username {
			usernameExists = true
		}
		if user.Email == email {
			emailExists = true
		}
	}
	return usernameExists, emailExists, nil
}

func CreateUser(user models.User) (id uint, err error) {
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return 0, errs.TranslateError(err)
	}
	return user.ID, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("user_name = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, errs.TranslateError(err)
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
	for k, v := range updateData {
		if v == "" {
			delete(updateData, k)
		}
	}
	if len(updateData) == 0 {
		return nil
	}
	err := db.GetDBConn().Model(&user).Where("id = ?", user.ID).Updates(updateData).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUser] Failed to update user with ID %v: %v\n", user.ID, err)
		return errs.TranslateError(err)
	}
	return nil
}

func DeleteUser(id uint) error {
	err := db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Println("[repository.DeleteUser] Failed to soft delete user with ID %v: %v\n", id, err)
		return errs.TranslateError(err)
	}
	return nil
}

func UpdateUserPassword(id uint, newPassword string) error {
	err := db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("password", newPassword).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserPassword] Failed to update password for user with ID %v: %v\n", id, err)
		return errs.TranslateError(err)
	}
	return nil
}
