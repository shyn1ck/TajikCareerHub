package errs

import "errors"

var (
	ErrValidationFailed            = errors.New("ErrValidationFailed")
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrIncorrectUserNameOrPassword = errors.New("ErrIncorrectUsernameorPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
	ErrUserBlocked                 = errors.New("ErrUserBlocked")
)
