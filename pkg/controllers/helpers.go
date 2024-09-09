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
		errors.Is(err, errs.ErrIncorrectUserNameOrPassword),
		errors.Is(err, errs.ErrValidationFailed),
		errors.Is(err, errs.ErrUserBlocked):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, errs.ErrPermissionDenied):
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	case errors.Is(err, errs.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
	}
}
