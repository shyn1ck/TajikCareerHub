package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Where("deleted_at = false").Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %s\n", err.Error())
		return nil, TranslateError(err)
	}
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ? AND deleted_at = false", id).First(&user).Error
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
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return nil, TranslateError(err)
	}
	return &user, nil
}

//unused function
//func UserExists(username, email string) (bool, bool, error) {
//	var usernameExists, emailExists bool
//	var user models.User
//	err := db.GetDBConn().Where("user_name = ? AND deleted_at = false", username).First(&user).Error
//	if err == nil {
//		usernameExists = true
//	} else if !errors.Is(err, errs.ErrRecordNotFound) {
//		return false, false, TranslateError(err)
//	}
//	err = db.GetDBConn().Where("email = ? AND deleted_at = false", email).First(&user).Error
//	if err == nil {
//		emailExists = true
//	} else if !errors.Is(err, errs.ErrRecordNotFound) {
//		return false, false, TranslateError(err)
//	}
//
//	return usernameExists, emailExists, nil
//}

func CreateUser(user models.User) (uint, error) {
	if err := db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return 0, TranslateError(err)
	}
	return user.ID, nil
}

func GetUserByUsernameAndPassword(username, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("user_name = ? AND password = ? AND deleted_at = false", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return models.User{}, TranslateError(err)
	}
	return user, nil
}

func UpdateUser(user models.User) (err error) {
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
	err = db.GetDBConn().
		Model(&models.User{}).
		Where("id = ? AND deleted_at = false", user.ID).
		Updates(updateData).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUser] Failed to update user with ID %v: %v\n", user.ID, err)
		return TranslateError(err)
	}
	return nil
}

func DeleteUser(id uint) (err error) {
	err = db.GetDBConn().
		Model(&models.User{}).
		Where("id = ? AND deleted_at = false", id).
		Update("deleted_at", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteUser] Failed to soft delete user with ID %v: %v\n", id, err)
		return TranslateError(err)
	}
	return nil
}

func UpdateUserPassword(id uint, newPassword string) (err error) {
	err = db.GetDBConn().
		Model(&models.User{}).
		Where("id = ? AND deleted_at = false", id).
		Update("password", newPassword).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserPassword] Failed to update password for user with ID %v: %v\n", id, err)
		return TranslateError(err)
	}
	return nil
}

func updateBlockStatus(id uint, isBlocked bool) (err error) {
	err = db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("is_blocked", isBlocked).Error
	if err != nil {
		action := "block"
		if !isBlocked {
			action = "unblock"
		}
		logger.Error.Printf("[repository.updateBlockStatus] Failed to %s user with ID %v: %v\n", action, id, err)
		return TranslateError(err)
	}
	return nil
}

func GetSpecialistActivityReportByUser(userID uint) (reports []models.SpecialistActivityReport, err error) {
	err = db.GetDBConn().
		Table("users").
		Select("users.id as user_id, users.full_name as user_name, COUNT(applications.id) as application_count").
		Joins("left join applications on applications.user_id = users.id").
		Where("applications.deleted_at = false AND users.id = ?", userID).
		Group("users.id").
		Scan(&reports).Error

	if err != nil {
		logger.Error.Printf("[repository.GetSpecialistActivityReportByUser] Error retrieving specialist activity report for user with ID %d: %v", userID, err)
		return nil, TranslateError(err)
	}

	if len(reports) == 0 {
		logger.Info.Printf("[repository.GetSpecialistActivityReportByUser] No activity reports found for user with ID %d", userID)
		return nil, TranslateError(err)
	}
	logger.Info.Printf("[repository.GetSpecialistActivityReportByUser] Successfully retrieved activity report for user with ID %d", userID)
	return reports, nil
}

func BlockUser(id uint) error {
	return updateBlockStatus(id, true)
}

func UnBlockUser(id uint) error {
	return updateBlockStatus(id, false)
}
