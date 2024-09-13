package controllers

import (
	"TajikCareerHub/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}

func handleError(c *gin.Context, err error) {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrRecordNotFound),
		errors.Is(err, errs.ErrUsersNotFound),
		errors.Is(err, errs.ErrUserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrPermissionDenied),
		errors.Is(err, errs.ErrAccessDenied),
		errors.Is(err, errs.ErrUserBlocked),
		errors.Is(err, errs.ErrResumeBlocked),
		errors.Is(err, errs.ErrVacancyBlocked),
		errors.Is(err, errs.ErrRoleCannotBeAdmin),
		errors.Is(err, errs.ErrRoleExist):
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrResumeCreationFailed),
		errors.Is(err, errs.ErrJobCreationFailed),
		errors.Is(err, errs.ErrApplicationFailed),
		errors.Is(err, errs.ErrReviewSubmissionFailed),
		errors.Is(err, errs.ErrReportGenerationFailed):
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrForeignKeyViolation),
		errors.Is(err, errs.ErrNotNullViolation),
		errors.Is(err, errs.ErrStringTooLong),
		errors.Is(err, errs.ErrCheckConstraintViolation),
		errors.Is(err, errs.ErrUniqueViolation),
		errors.Is(err, errs.ErrDeadlockDetected):
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
	}
}

type defaultResponse struct {
	Message string `json:"message"`
}

func newDefaultResponse(message string) defaultResponse {
	return defaultResponse{
		Message: message,
	}
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}
