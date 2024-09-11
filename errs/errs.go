package errs

import "errors"

var (
	ErrValidationFailed            = errors.New("ErrValidationFailed")
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
	ErrUserBlocked                 = errors.New("ErrUserBlocked")
	ErrUsersNotFound               = errors.New("ErrUsersNotFound")
	ErrUsernameExists              = errors.New("ErrUsernameAlreadyExists")
	ErrEmailExists                 = errors.New("ErrEmailAlreadyExists")
	ErrIncorrectPassword           = errors.New("ErrIncorrectPassword")
	ErrUserNotFound                = errors.New("ErrUserNotFound")
	ErrAccessDenied                = errors.New("ErrAccessDenied")
	ErrResumeCreationFailed        = errors.New("ErrResumeCreationFailed")
	ErrJobCreationFailed           = errors.New("ErrVacancyCreationFailed")
	ErrApplicationFailed           = errors.New("ErrApplicationFailed")
	ErrReviewSubmissionFailed      = errors.New("ErrReviewSubmissionFailed")
	ErrReportGenerationFailed      = errors.New("ErrReportGenerationFailed")
	ErrResumeBlocked               = errors.New("ErrResumeBlocked")
	ErrVacancyBlocked              = errors.New("ErrVacancyBlocked")
	ErrIDIsNotCorrect              = errors.New("ErrIDIsNotCorrect")
)
