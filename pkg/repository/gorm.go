package repository

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/utils/errs"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func TranslateError(err error) error {

	// Проверка на отсутствие записи
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Warning.Printf("Record not found error: %v", err)
		return errs.ErrRecordNotFound
	}

	// Проверка ошибок PostgreSQL (pgconn.PgError)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // Нарушение уникальности
			logger.Warning.Printf("Uniqueness violation: %v", err)
			return errs.ErrUniquenessViolation

		case "23503": // Нарушение внешнего ключа
			logger.Warning.Printf("Foreign key violation: %v", err)
			return errs.ErrForeignKeyViolation

		case "23502": // Нарушение not-null constraint
			logger.Warning.Printf("Not null constraint violation: %v", err)
			return errs.ErrNotNullViolation

		case "22001": // Превышение длины строки
			logger.Warning.Printf("String too long error: %v", err)
			return errs.ErrStringTooLong

		case "23514": // Нарушение check constraint
			logger.Warning.Printf("Check constraint violation: %v", err)
			return errs.ErrCheckConstraintViolation

		case "40P01": // Обнаружен дедлок
			logger.Warning.Printf("Deadlock detected: %v", err)
			return errs.ErrDeadlockDetected
		case "42702":
			logger.Warning.Printf("Foreign key violation: %v", err)
			return errs.ErrRecordNotFound

		default: // Необработанная ошибка PostgreSQL
			logger.Error.Printf("Unhandled PostgreSQL error: %v", err)
			return errs.ErrSomethingWentWrong
		}
	}

	// Если ошибка не попадает под перечисленные категории
	logger.Error.Printf("Unhandled error: %v", err)
	return errs.ErrSomethingWentWrong
}
