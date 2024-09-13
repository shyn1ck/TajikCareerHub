package controllers

import (
	"TajikCareerHub/errs"
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
		errors.Is(err, errs.ErrIncorrectPasswordLength):
		statusCode = http.StatusBadRequest
		errorResponse = newErrorResponse(err.Error())

	case errors.Is(err, errs.ErrRecordNotFound),
		errors.Is(err, errs.ErrUsersNotFound),
		errors.Is(err, errs.ErrUserNotFound):
		statusCode = http.StatusNotFound
		errorResponse = newErrorResponse(err.Error())

	case errors.Is(err, errs.ErrPermissionDenied),
		errors.Is(err, errs.ErrAccessDenied),
		errors.Is(err, errs.ErrUserBlocked),
		errors.Is(err, errs.ErrResumeBlocked),
		errors.Is(err, errs.ErrVacancyBlocked),
		errors.Is(err, errs.ErrRoleCannotBeAdmin),
		errors.Is(err, errs.ErrRoleExist):
		statusCode = http.StatusForbidden
		errorResponse = newErrorResponse(err.Error())

	case errors.Is(err, errs.ErrResumeCreationFailed),
		errors.Is(err, errs.ErrJobCreationFailed),
		errors.Is(err, errs.ErrApplicationFailed),
		errors.Is(err, errs.ErrReviewSubmissionFailed),
		errors.Is(err, errs.ErrReportGenerationFailed):
		statusCode = http.StatusInternalServerError
		errorResponse = newErrorResponse(err.Error())

	case errors.Is(err, errs.ErrForeignKeyViolation),
		errors.Is(err, errs.ErrNotNullViolation),
		errors.Is(err, errs.ErrStringTooLong),
		errors.Is(err, errs.ErrCheckConstraintViolation),
		errors.Is(err, errs.ErrUniqueViolation),
		errors.Is(err, errs.ErrDeadlockDetected):
		statusCode = http.StatusInternalServerError
		errorResponse = newErrorResponse(err.Error())

	default:
		statusCode = http.StatusInternalServerError
		errorResponse = newErrorResponse(errs.ErrSomethingWentWrong.Error())
	}

	c.JSON(statusCode, errorResponse)
}
