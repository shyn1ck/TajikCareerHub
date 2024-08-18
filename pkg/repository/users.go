package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/models"
	"log"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Where("deleted_at = ?", false).Find(&users).Error
	if err != nil {
		log.Printf("repository.GetAllUsers: Error retrieving all users. Error: %v\n", err)
		return nil, err
	}
	log.Println("repository.GetAllUsers: Successfully retrieved all users.")
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	log.Printf("repository.GetUserByID: Fetching user by ID %v from the database...\n", id)
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Printf("repository.GetUserByID: Failed to fetch user by ID %v. Error: %v\n", id, err)
		return user, err
	}
	log.Printf("repository.GetUserByID: Successfully fetched user by ID %v.\n", id)
	return user, nil
}

func GetUserByUserName(username string) (user models.User, err error) {
	log.Printf("repository.GetUserByUserName: Fetching user by username %v from the database...\n", username)
	err = db.GetDBConn().Where("user_name = ?", username).First(&user).Error
	if err != nil {
		log.Printf("repository.GetUserByUserName: Failed to fetch user by user_name %v. Error: %v\n", username, err)
		return user, err
	}
	log.Printf("repository.GetUserByUserName: Successfully fetched user by user_name %v.\n", username)
	return user, nil
}

func CreateUser(user models.User) (err error) {
	log.Println("repository.CreateUser: Creating a new user in the database...")
	err = db.GetDBConn().Create(&user).Error
	if err != nil {
		log.Printf("repository.CreateUser: Failed to create user. Error: %v\n", err)
		return err
	}
	log.Println("repository.CreateUser: User created successfully.")
	return nil
}

func UpdateUser(user models.User) (err error) {
	log.Printf("repository.UpdateUser: Updating user with ID %v...\n", user.ID)
	err = db.GetDBConn().Save(&user).Error
	if err != nil {
		log.Printf("repository.UpdateUser: Failed to update user with ID %v. Error: %v\n", user.ID, err)
		return err
	}
	log.Printf("repository.UpdateUser: User with ID %v updated successfully.\n", user.ID)
	return nil
}

func DeleteUser(id uint) (err error) {
	log.Printf("repository.DeleteUser: Soft deleting user with ID %v...\n", id)
	err = db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("is_deleted", true).Error
	if err != nil {
		log.Printf("repository.DeleteUser: Failed to soft delete user with ID %v. Error: %v\n", id, err)
		return err
	}
	log.Printf("repository.DeleteUser: User with ID %v successfully soft deleted.\n", id)
	return nil
}

func UpdateUserPassword(id uint, newPassword string) (err error) {
	log.Printf("repository.UpdateUserPassword: Updating password for user with ID %v...\n", id)
	err = db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("password", newPassword).Error
	if err != nil {
		log.Printf("repository.UpdateUserPassword: Failed to update password for user with ID %v. Error: %v\n", id, err)
		return err
	}
	log.Printf("repository.UpdateUserPassword: Password for user with ID %v updated successfully.\n", id)
	return nil
}

func CheckUserExists(username string, email string) (exists bool, err error) {
	log.Printf("repository.CheckUserExists: Checking if user with username %v or email %v exists...\n", username, email)
	var count int64
	err = db.GetDBConn().Model(&models.User{}).
		Where("user_name = ? OR email = ?", username, email).
		Count(&count).Error
	if err != nil {
		log.Printf("repository.CheckUserExists: Error checking user existence. Error: %v\n", err)
		return false, err
	}
	exists = count > 0
	log.Printf("repository.CheckUserExists: User existence check complete. Exists: %v\n", exists)
	return exists, nil
}
