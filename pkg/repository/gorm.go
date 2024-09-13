package repository

import (
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"errors"
	"gorm.io/gorm"
)

import (
	"strings"
)

func TranslateError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Warning.Printf("Record not found error: %v...", err)
		return errs.ErrRecordNotFound
	}

	// Обработка ошибки нарушения уникальности (SQLSTATE 23505)
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		if strings.Contains(err.Error(), "uni_users_user_name") {
			logger.Warning.Printf("Username already exists error: %v...", err)
			return errs.ErrUsernameExists
		}
		if strings.Contains(err.Error(), "uni_users_email") {
			logger.Warning.Printf("Email already exists error: %v...", err)
			return errs.ErrEmailExists
		}
	}

	// Обработка ошибки нарушения внешнего ключа (SQLSTATE 23503)
	if strings.Contains(err.Error(), "violates foreign key constraint") {
		logger.Warning.Printf("Foreign key violation: %v...", err)
		return errs.ErrForeignKeyViolation
	}

	// Обработка ошибки нарушения NOT NULL (SQLSTATE 23502)
	if strings.Contains(err.Error(), "null value in column") {
		logger.Warning.Printf("Not null constraint violation: %v...", err)
		return errs.ErrNotNullViolation
	}

	// Обработка ошибки превышения длины строки (SQLSTATE 22001)
	if strings.Contains(err.Error(), "value too long for type") {
		logger.Warning.Printf("String too long error: %v...", err)
		return errs.ErrStringTooLong
	}

	// Обработка ошибки нарушения CHECK constraint (SQLSTATE 23514)
	if strings.Contains(err.Error(), "violates check constraint") {
		logger.Warning.Printf("Check constraint violation: %v...", err)
		return errs.ErrCheckConstraintViolation
	}

	// Обработка ошибки нарушения уникальности (SQLSTATE 23505)
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		logger.Warning.Printf("Unique constraint violation: %v...", err)
		return errs.ErrUniqueViolation
	}

	// Обработка других общих ошибок
	if strings.Contains(err.Error(), "deadlock detected") {
		logger.Warning.Printf("Deadlock detected: %v...", err)
		return errs.ErrDeadlockDetected
	}

	logger.Error.Printf("Unhandled error: %v...", err)
	return errs.ErrSomethingWentWrong
}
