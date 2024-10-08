package errs

import "errors"

var (
	ErrValidationFailed                            = errors.New("ErrValidationFailed")
	ErrPermissionDenied                            = errors.New("ErrPermissionDenied")
	ErrUsernameUniquenessFailed                    = errors.New("ErrUsernameUniquenessFailed")
	ErrIncorrectUsernameOrPassword                 = errors.New("ErrIncorrectUsernameOrPassword")
	ErrRecordNotFound                              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong                          = errors.New("ErrSomethingWentWrong")
	ErrUserBlocked                                 = errors.New("ErrUserBlocked")
	ErrUsersNotFound                               = errors.New("ErrUsersNotFound")
	ErrUsernameExists                              = errors.New("ErrUsernameAlreadyExists")
	ErrEmailExists                                 = errors.New("ErrEmailAlreadyExists")
	ErrIncorrectPassword                           = errors.New("ErrIncorrectPassword")
	ErrUserNotFound                                = errors.New("ErrUserNotFound")
	ErrAccessDenied                                = errors.New("ErrAccessDenied")
	ErrResumeCreationFailed                        = errors.New("ErrResumeCreationFailed")
	ErrJobCreationFailed                           = errors.New("ErrVacancyCreationFailed")
	ErrApplicationFailed                           = errors.New("ErrApplicationFailed")
	ErrReviewSubmissionFailed                      = errors.New("ErrReviewSubmissionFailed")
	ErrReportGenerationFailed                      = errors.New("ErrReportGenerationFailed")
	ErrResumeBlocked                               = errors.New("ErrResumeBlocked")
	ErrVacancyBlocked                              = errors.New("ErrVacancyBlocked")
	ErrIDIsNotCorrect                              = errors.New("ErrIDIsNotCorrect")
	ErrRoleCannotBeAdmin                           = errors.New("ErrRoleCannotBeAdmin")
	ErrRoleExist                                   = errors.New("ErrRoleExist ")
	ErrIncorrectPasswordLength                     = errors.New("ErrIncorrectPasswordLength ")
	ErrForeignKeyViolation                         = errors.New("ErrForeignKeyViolation ")
	ErrNotNullViolation                            = errors.New("ErrNotNullViolation ")
	ErrStringTooLong                               = errors.New("ErrStringTooLong")
	ErrCheckConstraintViolation                    = errors.New("ErrCheckConstraintViolation")
	ErrUniqueViolation                             = errors.New("ErrUniqueViolation")
	ErrDeadlockDetected                            = errors.New("ErrDeadlockDetected")
	ErrInvalidRole                                 = errors.New("ErrInvalidRole")
	ErrFullNameIsRequired                          = errors.New("ErrFullNameIsRequired")
	ErrVacancyCategoryIsRequired                   = errors.New("ErrVacancyCategoryIsRequired")
	ExperienceYearsCannotBeNegative                = errors.New("ExperienceYearsCannotBeNegative")
	SummaryCannotExceedDefiniteCharacters          = errors.New("SummaryCannotExceed1000Characters")
	ErrTitleIsRequired                             = errors.New("ErrTitleIsRequired")
	ErrTitleMustBeLessThanDefiniteCharacters       = errors.New("ErrTitleMustBeLessThan100Characters")
	ErrDescriptionIsRequired                       = errors.New("ErrDescriptionIsRequired")
	ErrDescriptionMustBeLessThanDefiniteCharacters = errors.New("ErrDescriptionMustBeLessThan100Characters")
	ErrSalaryMustBeANonNegativeNumber              = errors.New("ErrSalaryMustBeANonNegativeNumber")
	ErrCompanyIDIsRequired                         = errors.New("ErrCompanyIDIsRequired")
	ErrUserIdDoesNotMatchTheProvidedUsername       = errors.New("ErrUserIdDoesNotMatchTheProvidedUsername")
	ErrShouldBindJson                              = errors.New("ErrShouldBindJson")
	ErrCategoryAlreadyExist                        = errors.New("ErrCategoryAlreadyExist")
	ErrNoReportsFound                              = errors.New("ErrNoReportsFound")
	ErrIDIsNotProvided                             = errors.New("ErrIDIsNotProvided")
	ErrInvalidToken                                = errors.New("ErrInvalidToken")
	ErrUnexpectedSigningMethod                     = errors.New("ErrUnexpectedSigningMethod")
	ErrTokenParseError                             = errors.New("ErrTokenParseError")
	ErrAuthorizationHeaderMissing                  = errors.New("ErrAuthorizationHeaderMissing")
	ErrCompanyNotFound                             = errors.New("ErrCompanyNotFound")
	ErrIncorrectInput                              = errors.New("ErrIncorrectInput")
	ErrUniquenessViolation                         = errors.New("ErrUniquenessViolation")
	ErrResumeNotFound                              = errors.New("ErrResumeNotFound")
	ErrVacancyNotFound                             = errors.New("ErrVacancyNotFound")
)
