package repository

import (
	"TajikCareerHub/errs"
	"errors"
	"gorm.io/gorm"
)

func translateError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}

	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return errs.ErrDuplicateEntry
	}

	if err != nil && errors.Is(err, gorm.ErrInvalidData) {
		return errs.ErrInvalidData
	}

	return err
}
