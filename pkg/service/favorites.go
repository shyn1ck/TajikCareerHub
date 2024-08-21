package service

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/repository"
)

func GetFavoritesByUserID(userID uint) ([]models.Favorite, error) {
	favorites, err := repository.GetFavoritesByUserID(userID)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func GetFavoriteByUserIDAndJobID(userID uint, jobID uint) (models.Favorite, error) {
	favorite, err := repository.GetFavoriteByUserIDAndJobID(userID, jobID)
	if err != nil {
		return favorite, err
	}
	return favorite, nil
}

func AddFavorite(favorite models.Favorite) error {
	exists, err := CheckFavoriteExists(favorite.UserID, favorite.JobID)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return repository.AddFavorite(favorite)
}

func RemoveFavorite(userID uint, jobID uint) error {
	return repository.RemoveFavorite(userID, jobID)
}

func CheckFavoriteExists(userID uint, jobID uint) (bool, error) {
	var count int64
	exists, err := repository.CheckFavoriteExists(count, userID, jobID)
	if err != nil {
		return false, err
	}
	return exists, nil
}
