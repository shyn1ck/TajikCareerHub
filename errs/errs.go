package errs

import "errors"

var (
	ErrValidationFailed            = errors.New("ValidationFailed")
	ErrPermissionDenied            = errors.New("PermissionDenied")
	ErrUsernameUniquenessFailed    = errors.New("UsernameUniquenessFailed")
	ErrIncorrectUsernameOrPassword = errors.New("IncorrectUsernameOrPassword")
	ErrRecordNotFound              = errors.New("RecordNotFound")
	ErrSomethingWentWrong          = errors.New("SomethingWentWrong")
	ErrUserBlocked                 = errors.New("UserBlocked")
	ErrUsersNotFound               = errors.New("UsersNotFound")
	ErrUsernameExists              = errors.New("UsernameAlreadyExists")
	ErrEmailExists                 = errors.New("EmailAlreadyExists")
	ErrIncorrectPassword           = errors.New("IncorrectPassword")
	ErrUserNotFound                = errors.New("UserNotFound")
	ErrAccessDenied                = errors.New("AccessDenied")
	ErrResumeCreationFailed        = errors.New("ResumeCreationFailed")
	ErrJobCreationFailed           = errors.New("JobCreationFailed")
	ErrApplicationFailed           = errors.New("ApplicationFailed")
	ErrReviewSubmissionFailed      = errors.New("ReviewSubmissionFailed")
	ErrReportGenerationFailed      = errors.New("ReportGenerationFailed")
)
