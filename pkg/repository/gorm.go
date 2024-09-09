package repository

import (
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"errors"
	"gorm.io/gorm"
)

func TranslateError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Warning.Printf("Record not found error: %v...", err)
		return errs.ErrRecordNotFound
	}
	logger.Error.Printf("Unhandled error: %v...", err)
	return err
}
