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
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword),
		errors.Is(err, errs.ErrIncorrectPassword),
		errors.Is(err, errs.ErrEmailExists),
		errors.Is(err, errs.ErrUsernameExists),
		errors.Is(err, errs.ErrValidationFailed):
		// Ошибка уникальности, неверного пароля или валидации
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrRecordNotFound),
		errors.Is(err, errs.ErrUsersNotFound),
		errors.Is(err, errs.ErrUserNotFound):
		// Ошибка "Запись не найдена" или "Пользователь не найден"
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrPermissionDenied),
		errors.Is(err, errs.ErrAccessDenied),
		errors.Is(err, errs.ErrUserBlocked):
		// Ошибка "Доступ запрещен" или "Пользователь заблокирован"
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrResumeCreationFailed),
		errors.Is(err, errs.ErrJobCreationFailed),
		errors.Is(err, errs.ErrApplicationFailed),
		errors.Is(err, errs.ErrReviewSubmissionFailed),
		errors.Is(err, errs.ErrReportGenerationFailed):
		// Ошибка при создании резюме, вакансии, подачи заявки, отправке отзыва или генерации отчета
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	default:
		// Здесь просто возвращаем внутреннюю ошибку клиенту
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
	}
}
