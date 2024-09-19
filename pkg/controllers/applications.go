package controllers

import (
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllApplications godoc
// @Summary Get all applications
// @Description Get a list of all applications. Requires authentication.
// @Tags Applications
// @Accept json
// @Produce json
// @Success 200 {array} models.SwaggerApplication
// @Failure 401 {object} ErrorResponse
// @Router /application [get]
// @Security ApiKeyAuth
func GetAllApplications(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetAllApplications] Client IP: %s - Client requested all applications\n", ip)
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}
	applications, err := service.GetAllApplications(userID)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetAllApplications] Client IP: %s - Successfully retrieved all applications\n", ip)
	c.JSON(http.StatusOK, applications)
}

// GetApplicationByID godoc
// @Summary Get application by ID
// @Description Get a single application by its ID. Requires authentication.
// @Tags Applications
// @Accept json
// @Produce json
// @Param id path integer true "Application ID"
// @Success 200 {object} models.SwaggerApplication
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /application/{id} [get]
// @Security ApiKeyAuth
func GetApplicationByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationByID] Client IP: %s - Client requested application with ID %s. Error: Invalid application ID\n", ip, idStr)
		handleError(c, errs.ErrIDIsNotCorrect)
		return
	}
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}
	application, err := service.GetApplicationByID(userID, uint(id))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationByID] Client IP: %s - Error retrieving application with ID %v\n", ip, id)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetApplicationByID] Client IP: %s - Successfully retrieved application with ID %v\n", ip, id)
	c.JSON(http.StatusOK, application)
}

// AddApplication godoc
// @Summary Add a new application
// @Description Add a new application. Requires authentication.
// @Tags Applications
// @Accept json
// @Produce json
// @Param application body models.SwaggerApplication true "Application data"
// @Success 201 {object} DefaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /application [post]
// @Security ApiKeyAuth
func AddApplication(c *gin.Context) {
	ip := c.ClientIP()
	var application models.SwaggerApplication
	if err := c.ShouldBindJSON(&application); err != nil {
		logger.Info.Printf("[controllers.AddApplication] Client IP: %s - Client attempted to add application with data %v. Error: Invalid input\n", ip, application)
		handleError(c, err)
		return
	}
	app := models.Application{
		UserID:    application.UserID,
		VacancyID: application.VacancyID,
		ResumeID:  application.ResumeID,
		StatusID:  application.StatusID,
	}
	if err := service.AddApplication(app); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.AddApplication] Client IP: %s - Successfully added application with data %v\n", ip, application)
	c.JSON(http.StatusCreated, NewDefaultResponse("Application added successfully"))
}

// UpdateApplication godoc
// @Summary Update an existing application
// @Description Update an application by its ID. Requires authentication.
// @Tags Applications
// @Accept json
// @Produce json
// @Param id path integer true "Application ID"
// @Param application body models.SwaggerApplication true "Updated application data"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /application/{id} [put]
// @Security ApiKeyAuth
func UpdateApplication(c *gin.Context) {
	ip := c.ClientIP()
	var application models.SwaggerApplication
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Client attempted to update application with ID %s. Error: Invalid application ID\n", ip, idStr)
		handleError(c, err)
		return
	}
	app := models.Application{
		ID:        uint(id),
		UserID:    application.UserID,
		VacancyID: application.VacancyID,
		ResumeID:  application.ResumeID,
		StatusID:  application.StatusID,
	}
	if err := c.ShouldBindJSON(&application); err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Client attempted to update application with ID %v using data %v. Error: Invalid input\n", ip, id, application)
		handleError(c, err)
		return
	}
	if err := service.UpdateApplication(app); err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Error updating application with ID %v\n", ip, id)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Successfully updated application with ID %v\n", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Application updated successfully"))
}

// DeleteApplication godoc
// @Summary Delete an application
// @Description Soft delete an application by its ID. Requires authentication.
// @Tags Applications
// @Accept json
// @Produce json
// @Param id path integer true "Application ID"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /application/{id} [delete]
// @Security ApiKeyAuth
func DeleteApplication(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.DeleteApplication] Client IP: %s - Client attempted to delete application with ID %s. Error: Invalid application ID\n", ip, idStr)
		handleError(c, err)
		return
	}
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	if err := service.DeleteApplication(userID, uint(id)); err != nil {
		logger.Info.Printf("[controllers.DeleteApplication] Client IP: %s - Error deleting application with ID %v\n", ip, id)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.DeleteApplication] Client IP: %s - Successfully soft deleted application with ID %v\n", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Application deleted successfully"))
}

// UpdateApplicationStatus godoc
// @Summary Update the status of an application
// @Description Update the status of a specific application by its ID
// @Tags Applications
// @Accept json
// @Produce json
// @Param application_id path uint true "Application ID" example(123)
// @Param status_id path uint true "Status ID" example(2)
// @Success 200 {object} DefaultResponse "Status updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Forbidden access"
// @Failure 404 {object} ErrorResponse "Application not found"
// @Security ApiKeyAuth
// @Router /application/{application_id}/status/{status_id} [put]
func UpdateApplicationStatus(c *gin.Context) {
	applicationIDStr := c.Param("application_id")
	statusIDStr := c.Param("status_id")
	applicationID, err := strconv.ParseUint(applicationIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	statusID, err := strconv.ParseUint(statusIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	err = service.UpdateApplicationStatus(uint(applicationID), uint(statusID))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, NewDefaultResponse("Application status updated successfully"))
}

// GetSpecialistActivityReportByUser godoc
// @Summary Get specialist activity report for a specific user
// @Tags Reports
// @Description Get a report of how many vacancies a specific specialist has applied for
// @ID get-specialist-activity-report-by-user
// @Accept json
// @Produce json
// @Param user_id path uint true "User ID"
// @Success 200 {array} models.SpecialistActivityReport
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Forbidden access"
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /activity/{user_id} [get]
func GetSpecialistActivityReportByUser(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetSpecialistActivityReportByUser] Client IP: %s - Request to get specialist activity report for user\n", ip)
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	reports, err := service.GetSpecialistActivityReportByUser(uint(userID))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetSpecialistActivityReportByUser] Client IP: %s - Successfully retrieved specialist activity report for user\n", ip)
	c.JSON(http.StatusOK, reports)
}
