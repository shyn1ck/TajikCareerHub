package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"TajikCareerHub/utils/errs"
	"fmt"
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
// @failure 403 {object} ErrorResponse "Access Denied"
// @Failure 401 {object} ErrorResponse
// @Router /applications [get]
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
// @Router /applications/{id} [get]
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
// @Router /applications [post]
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
// @Router /applications/{id} [put]
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
// @Router /applications/{id} [delete]
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
// @Router /applications/{application_id}/status/{status_id} [put]
func UpdateApplicationStatus(c *gin.Context) {
	ip := c.ClientIP()
	applicationIDStr := c.Param("application_id")
	statusIDStr := c.Param("status_id")
	logger.Info.Printf("[controllers.UpdateApplicationStatus] Raw application ID: %s, status ID: %s, Client IP: %s\n", applicationIDStr, statusIDStr, ip)

	if applicationIDStr == "" {
		logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Missing application ID\n", ip)
		handleError(c, fmt.Errorf("missing application ID"))
		return
	}
	if statusIDStr == "" {
		logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Missing status ID\n", ip)
		handleError(c, fmt.Errorf("missing status ID"))
		return
	}

	applicationID, err := strconv.ParseUint(applicationIDStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Invalid application ID: %s. Error: %v\n", ip, applicationIDStr, err)
		handleError(c, fmt.Errorf("invalid application ID: %s", applicationIDStr))
		return
	}

	statusID, err := strconv.ParseUint(statusIDStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Invalid status ID: %s. Error: %v\n", ip, statusIDStr, err)
		handleError(c, fmt.Errorf("invalid status ID: %s", statusIDStr))
		return
	}

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Failed to extract user ID from token. Error: %v\n", ip, err)
		handleError(c, fmt.Errorf("unable to extract user ID from token"))
		return
	}

	logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Attempting to update status for application ID: %d to status ID: %d by user ID: %d\n", ip, applicationID, statusID, userID)

	err = service.UpdateApplicationStatus(uint(applicationID), uint(statusID), userID)
	if err != nil {
		logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Failed to update status for application ID: %d. Error: %v\n", ip, applicationID, err)
		handleError(c, fmt.Errorf("failed to update application status"))
		return
	}

	logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Successfully updated application ID: %d to status ID: %d\n", ip, applicationID, statusID)
	c.JSON(http.StatusOK, DefaultResponse{Message: "Application status updated successfully"})
}
