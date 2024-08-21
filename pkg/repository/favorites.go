package repository

import (
	"TajikCareerHub/db"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
)

func GetFavoritesByUserID(userID uint) (favorites []models.Favorite, err error) {
	err = db.GetDBConn().Where("user_id = ?", userID).Find(&favorites).Error
	if err != nil {
		logger.Error.Printf("[repository.GetFavoritesByUserID]: Error retrieving favorites for user ID %v. Error: %v\n", userID, err)
		return nil, err
	}
	return favorites, nil
}

func GetFavoriteByUserIDAndJobID(userID uint, jobID uint) (favorite models.Favorite, err error) {
	err = db.GetDBConn().Where("user_id = ? AND job_id = ?", userID, jobID).First(&favorite).Error
	if err != nil {
		logger.Error.Printf("[repository.GetFavoriteByUserIDAndJobID]: Error retrieving favorite for user ID %v and job ID %v. Error: %v\n", userID, jobID, err)
		return favorite, err
	}
	return favorite, nil
}

func AddFavorite(favorite models.Favorite) error {
	err := db.GetDBConn().Create(&favorite).Error
	if err != nil {
		logger.Error.Printf("[repository.AddFavorite]: Failed to add favorite. Error: %v\n", err)
		return err
	}
	return nil
}

func RemoveFavorite(userID uint, jobID uint) error {
	err := db.GetDBConn().Where("user_id = ? AND job_id = ?", userID, jobID).Delete(&models.Favorite{}).Error
	if err != nil {
		logger.Error.Printf("[repository.RemoveFavorite]: Failed to remove favorite. Error: %v\n", err)
		return err
	}
	return nil
}

func CheckFavoriteExists(count int64, userID uint, jobID uint) (bool, error) {
	err := db.GetDBConn().Model(&models.Favorite{}).Where("user_id = ? AND job_id = ?", userID, jobID).Count(&count).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckFavoriteExists]: Error checking if job ID %v is in favorites for user ID %v. Error: %v\n", jobID, userID, err)
		return false, err
	}
	exists := count > 0
	return exists, nil
}
