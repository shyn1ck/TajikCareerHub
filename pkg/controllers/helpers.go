package controllers

import (
	"TajikCareerHub/utils/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(c *gin.Context, err error) {
	var statusCode int
	var errorResponse ErrorResponse

	switch {
	case errors.Is(err, errs.ErrUsernameUniquenessFailed),
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword),
		errors.Is(err, errs.ErrIncorrectPassword),
		errors.Is(err, errs.ErrEmailExists),
		errors.Is(err, errs.ErrUsernameExists),
		errors.Is(err, errs.ErrValidationFailed),
		errors.Is(err, errs.ErrIDIsNotCorrect),
		errors.Is(err, errs.ErrInvalidRole),
		errors.Is(err, errs.ErrIDIsNotProvided),
		errors.Is(err, errs.ErrIncorrectPasswordLength),
		errors.Is(err, errs.ErrFullNameIsRequired),
		errors.Is(err, errs.ErrVacancyCategoryIsRequired),
		errors.Is(err, errs.ExperienceYearsCannotBeNegative),
		errors.Is(err, errs.SummaryCannotExceedDefiniteCharacters),
		errors.Is(err, errs.ErrTitleIsRequired),
		errors.Is(err, errs.ErrTitleMustBeLessThanDefiniteCharacters),
		errors.Is(err, errs.ErrDescriptionIsRequired),
		errors.Is(err, errs.ErrDescriptionMustBeLessThanDefiniteCharacters),
		errors.Is(err, errs.ErrSalaryMustBeANonNegativeNumber),
		errors.Is(err, errs.ErrCompanyIDIsRequired),
		errors.Is(err, errs.ErrUserIdDoesNotMatchTheProvidedUsername),
		errors.Is(err, errs.ErrShouldBindJson),
		errors.Is(err, errs.ErrIncorrectInput),
		errors.Is(err, errs.ErrCategoryAlreadyExist):
		statusCode = http.StatusBadRequest
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrRecordNotFound),
		errors.Is(err, errs.ErrUsersNotFound),
		errors.Is(err, errs.ErrUserNotFound),
		errors.Is(err, errs.ErrCompanyNotFound):
		statusCode = http.StatusNotFound
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrPermissionDenied),
		errors.Is(err, errs.ErrAccessDenied),
		errors.Is(err, errs.ErrUserBlocked),
		errors.Is(err, errs.ErrResumeBlocked),
		errors.Is(err, errs.ErrVacancyBlocked),
		errors.Is(err, errs.ErrRoleCannotBeAdmin),
		errors.Is(err, errs.ErrRoleExist),
		errors.Is(err, errs.ErrInvalidToken),
		errors.Is(err, errs.ErrUnexpectedSigningMethod),
		errors.Is(err, errs.ErrAuthorizationHeaderMissing):
		statusCode = http.StatusForbidden
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrResumeCreationFailed),
		errors.Is(err, errs.ErrJobCreationFailed),
		errors.Is(err, errs.ErrApplicationFailed),
		errors.Is(err, errs.ErrReviewSubmissionFailed),
		errors.Is(err, errs.ErrReportGenerationFailed),
		errors.Is(err, errs.ErrTokenParseError):
		statusCode = http.StatusInternalServerError
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrForeignKeyViolation),
		errors.Is(err, errs.ErrNotNullViolation),
		errors.Is(err, errs.ErrStringTooLong),
		errors.Is(err, errs.ErrCheckConstraintViolation),
		errors.Is(err, errs.ErrUniqueViolation),
		errors.Is(err, errs.ErrDeadlockDetected):
		statusCode = http.StatusInternalServerError
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrNoReportsFound):
		statusCode = http.StatusNotFound
		errorResponse = NewErrorResponse(err.Error())

	default:
		statusCode = http.StatusInternalServerError
		errorResponse = NewErrorResponse(errs.ErrSomethingWentWrong.Error())
	}

	c.JSON(statusCode, errorResponse)
}
