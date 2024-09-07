package controllers

import (
	"TajikCareerHub/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(c *gin.Context, err error) {
	if errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrIncorrectUsernameorPassword) ||
		errors.Is(err, errs.ErrValidationFailed) ||
		errors.Is(err, errs.ErrDuplicateEntry) ||
		errors.Is(err, errs.ErrInvalidField) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrPermissionDenied) {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrRecordNotFound) ||
		errors.Is(err, errs.ErrRoutesNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
	}
}
