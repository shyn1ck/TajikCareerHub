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
	logger.Info.Printf("[repository.GetFavoritesByUserID]: Successfully retrieved favorites for user ID %v.\n", userID)
	return favorites, nil
}

func GetFavoriteByUserIDAndJobID(userID uint, jobID uint) (favorite models.Favorite, err error) {
	err = db.GetDBConn().Where("user_id = ? AND job_id = ?", userID, jobID).First(&favorite).Error
	if err != nil {
		logger.Error.Printf("[repository.GetFavoriteByUserIDAndJobID]: Error retrieving favorite for user ID %v and job ID %v. Error: %v\n", userID, jobID, err)
		return favorite, err
	}
	logger.Info.Printf("[repository.GetFavoriteByUserIDAndJobID]: Successfully retrieved favorite for user ID %v and job ID %v.\n", userID, jobID)
	return favorite, nil
}

func AddFavorite(favorite models.Favorite) error {
	logger.Info.Printf("[repository.AddFavorite]: Adding job ID %v to favorites for user ID %v...\n", favorite.JobID, favorite.UserID)
	err := db.GetDBConn().Create(&favorite).Error
	if err != nil {
		logger.Error.Printf("[repository.AddFavorite]: Failed to add favorite. Error: %v\n", err)
		return err
	}
	logger.Info.Printf("[repository.AddFavorite]: Job ID %v successfully added to favorites for user ID %v.\n", favorite.JobID, favorite.UserID)
	return nil
}

func RemoveFavorite(userID uint, jobID uint) error {
	logger.Info.Printf("[repository.RemoveFavorite]: Removing job ID %v from favorites for user ID %v...\n", jobID, userID)
	err := db.GetDBConn().Where("user_id = ? AND job_id = ?", userID, jobID).Delete(&models.Favorite{}).Error
	if err != nil {
		logger.Error.Printf("[repository.RemoveFavorite]: Failed to remove favorite. Error: %v\n", err)
		return err
	}
	logger.Info.Printf("[repository.RemoveFavorite]: Job ID %v successfully removed from favorites for user ID %v.\n", jobID, userID)
	return nil
}

func CheckFavoriteExists(count int64, userID uint, jobID uint) (bool, error) {
	err := db.GetDBConn().Model(&models.Favorite{}).Where("user_id = ? AND job_id = ?", userID, jobID).Count(&count).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckFavoriteExists]: Error checking if job ID %v is in favorites for user ID %v. Error: %v\n", jobID, userID, err)
		return false, err
	}
	exists := count > 0
	logger.Info.Printf("[repository.CheckFavoriteExists]: Job ID %v in favorites for user ID %v: %v\n", jobID, userID, exists)
	return exists, nil
}
