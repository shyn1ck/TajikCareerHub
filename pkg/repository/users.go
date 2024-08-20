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
	logger.Info.Println("[repository.GetAllUsers] Successfully retrieved all users.")
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	logger.Info.Printf("[repository.GetUserByID] Fetching user by ID %v...\n", id)
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] Failed to fetch user by ID %v: %v\n", id, err)
		return user, err
	}
	logger.Info.Printf("[repository.GetUserByID] Successfully fetched user by ID %v.\n", id)
	return user, nil
}

func GetUserByUserName(username string) (user models.User, err error) {
	logger.Info.Printf("[repository.GetUserByUserName] Fetching user by username %v...\n", username)
	err = db.GetDBConn().Where("user_name = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUserName] Failed to fetch user by username %v: %v\n", username, err)
		return user, err
	}
	logger.Info.Printf("[repository.GetUserByUserName] Successfully fetched user by username %v.\n", username)
	return user, nil
}

func CreateUser(user models.User) (err error) {
	logger.Info.Println("[repository.CreateUser] Creating a new user...")
	err = db.GetDBConn().Create(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.CreateUser] Failed to create user: %v\n", err)
		return err
	}
	logger.Info.Println("[repository.CreateUser] User created successfully.")
	return nil
}

func UpdateUser(user models.User) (err error) {
	logger.Info.Printf("[repository.UpdateUser] Updating user with ID %v...\n", user.ID)
	err = db.GetDBConn().Save(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUser] Failed to update user with ID %v: %v\n", user.ID, err)
		return err
	}
	logger.Info.Printf("[repository.UpdateUser] User with ID %v updated successfully.\n", user.ID)
	return nil
}

func DeleteUser(id uint) (err error) {
	logger.Info.Printf("[repository.DeleteUser] Soft deleting user with ID %v...\n", id)
	err = db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteUser] Failed to soft delete user with ID %v: %v\n", id, err)
		return err
	}
	logger.Info.Printf("[repository.DeleteUser] User with ID %v successfully soft deleted.\n", id)
	return nil
}

func UpdateUserPassword(id uint, newPassword string) (err error) {
	logger.Info.Printf("[repository.UpdateUserPassword] Updating password for user with ID %v...\n", id)
	err = db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("password", newPassword).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserPassword] Failed to update password for user with ID %v: %v\n", id, err)
		return err
	}
	logger.Info.Printf("[repository.UpdateUserPassword] Password for user with ID %v updated successfully.\n", id)
	return nil
}

func CheckUserExists(username string, email string) (exists bool, err error) {
	logger.Info.Printf("[repository.CheckUserExists] Checking if user with username %v or email %v exists...\n", username, email)
	var count int64
	err = db.GetDBConn().Model(&models.User{}).
		Where("user_name = ? OR email = ?", username, email).
		Count(&count).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckUserExists] Error checking user existence: %v\n", err)
		return false, err
	}
	exists = count > 0
	logger.Info.Printf("[repository.CheckUserExists] User existence check complete. Exists: %v\n", exists)
	return exists, nil
}
