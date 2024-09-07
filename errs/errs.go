package errs

import "errors"

var (
	ErrValidationFailed            = errors.New("ErrValidationFailed")
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrRoutesNotFound              = errors.New("ErrRoutesNotFound")
	ErrIncorrectUsernameorPassword = errors.New("ErrIncorrectUsernameorPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
	ErrNotImplemented              = errors.New("ErrNotImplemented")
	ErrUnsupportedDriver           = errors.New("ErrUnsupportedDriver")
	ErrInvalidData                 = errors.New("ErrInvalidData")
	ErrInvalidField                = errors.New("ErrInvalidField")
	ErrDuplicateEntry              = errors.New("ErrDuplicateEntry")
)
