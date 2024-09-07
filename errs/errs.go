package errs

import "errors"

var (
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrOperationNotFound           = errors.New("ErrOperationNotFound")
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
	ErrInvalidID                   = errors.New("ErrInvalidID")
	ErrFailedToBindJSON            = errors.New("ErrFailedToBindJSON")
	ErrPasswordUpdateFailed        = errors.New("ErrPasswordUpdateFailed")
	ErrUserCreationFailed          = errors.New("ErrUserCreationFailed")
	ErrUserUpdateFailed            = errors.New("ErrUserUpdateFailed")
	ErrUserDeletionFailed          = errors.New("ErrUserDeletionFailed")
	ErrUserExistsCheckFailed       = errors.New("ErrUserExistsCheckFailed")
	ErrEmailAlreadyExists          = errors.New("ErrEmailAlreadyExists")
	ErrInvalidEmailFormat          = errors.New("ErrInvalidEmailFormat")
	ErrWeakPassword                = errors.New("ErrWeakPassword")
	ErrEmailNotFound               = errors.New("ErrEmailNotFound")
	ErrDuplicateEntry              = errors.New("ErrDuplicateEntry")
	ErrInvalidData                 = errors.New("ErrInvalidData")
	ErrConnectionFailed            = errors.New("ErrConnectionFailed")
	ErrNotFound                    = errors.New(`ErrNotFound`)
)
