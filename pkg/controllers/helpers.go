package controllers

import (
	"TajikCareerHub/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUsernameUniquenessFailed),
		errors.Is(err, errs.ErrDuplicateEntry),
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword),
		errors.Is(err, errs.ErrInvalidID),
		errors.Is(err, errs.ErrFailedToBindJSON),
		errors.Is(err, errs.ErrPasswordUpdateFailed),
		errors.Is(err, errs.ErrUserCreationFailed),
		errors.Is(err, errs.ErrUserUpdateFailed),
		errors.Is(err, errs.ErrUserDeletionFailed),
		errors.Is(err, errs.ErrUserExistsCheckFailed),
		errors.Is(err, errs.ErrEmailAlreadyExists),
		errors.Is(err, errs.ErrInvalidEmailFormat),
		errors.Is(err, errs.ErrWeakPassword):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrRecordNotFound),
		errors.Is(err, errs.ErrOperationNotFound),
		errors.Is(err, errs.ErrEmailNotFound),
		errors.Is(err, errs.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
	}
}
