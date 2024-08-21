package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Where("deleted_at = ?", false).Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] Error retrieving all users: %v\n", err)
		return nil, err
	}
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] Failed to fetch user by ID %v: %v\n", id, err)
		return user, err
	}
	return user, nil
}

func GetUserByUserName(username string) (user models.User, err error) {
	err = db.GetDBConn().Where("user_name = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUserName] Failed to fetch user by username %v: %v\n", username, err)
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) (err error) {
	err = db.GetDBConn().Create(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.CreateUser] Failed to create user: %v\n", err)
		return err
	}
	return nil
}

func UpdateUser(user models.User) (err error) {
	err = db.GetDBConn().Save(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUser] Failed to update user with ID %v: %v\n", user.ID, err)
		return err
	}
	return nil
}

func DeleteUser(id uint) (err error) {
	err = db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteUser] Failed to soft delete user with ID %v: %v\n", id, err)
		return err
	}
	return nil
}

func UpdateUserPassword(id uint, newPassword string) (err error) {
	err = db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("password", newPassword).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserPassword] Failed to update password for user with ID %v: %v\n", id, err)
		return err
	}
	return nil
}

func CheckUserExists(username string, email string) (exists bool, err error) {
	var count int64
	err = db.GetDBConn().Model(&models.User{}).
		Where("user_name = ? OR email = ?", username, email).
		Count(&count).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckUserExists] Error checking user existence: %v\n", err)
		return false, err
	}
	exists = count > 0
	return exists, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, err
	}
	return user, nil
}
