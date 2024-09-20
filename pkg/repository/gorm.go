package repository

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/utils/errs"
	"errors"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"strings"
)

func TranslateError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Warning.Printf("Record not found error: %v...", err)
		return errs.ErrRecordNotFound
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		logger.Warning.Printf("Duplicated key error: %v...", err)
		return errs.ErrUniqueViolation
	}

	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		logger.Warning.Printf("Foreign key violation: %v...", err)
		return errs.ErrForeignKeyViolation
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505":
			if strings.Contains(pqErr.Message, "uni_users_user_name") {
				logger.Warning.Printf("Username already exists error: %v...", err)
				return errs.ErrUsernameExists
			}
			if strings.Contains(pqErr.Message, "uni_users_email") {
				logger.Warning.Printf("Email already exists error: %v...", err)
				return errs.ErrEmailExists
			}
			logger.Warning.Printf("Unique constraint violation: %v...", err)
			return errs.ErrUniqueViolation

		case "23503":
			logger.Warning.Printf("Foreign key violation: %v...", err)
			return errs.ErrForeignKeyViolation

		case "23502":
			logger.Warning.Printf("Not null constraint violation: %v...", err)
			return errs.ErrNotNullViolation

		case "22001":
			logger.Warning.Printf("String too long error: %v...", err)
			return errs.ErrStringTooLong

		case "23514":
			logger.Warning.Printf("Check constraint violation: %v...", err)
			return errs.ErrCheckConstraintViolation

		case "40P01":
			logger.Warning.Printf("Deadlock detected: %v...", err)
			return errs.ErrDeadlockDetected
		}
	}
	logger.Error.Printf("Standard error occurred: %v...", err)
	return err
}
