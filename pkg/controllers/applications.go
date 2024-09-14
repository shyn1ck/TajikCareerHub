package controllers

import (
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
// @Success 200 {array} models.Application
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
// @Success 200 {object} models.Application
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
		handleError(c, err)
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
// @Param application body models.Application true "Application data"
// @Success 201 {object} defaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /applications [post]
// @Security ApiKeyAuth
func AddApplication(c *gin.Context) {
	ip := c.ClientIP()
	var application models.Application
	if err := c.ShouldBindJSON(&application); err != nil {
		logger.Info.Printf("[controllers.AddApplication] Client IP: %s - Client attempted to add application with data %v. Error: Invalid input\n", ip, application)
		handleError(c, err)
		return
	}
	if err := service.AddApplication(application); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.AddApplication] Client IP: %s - Successfully added application with data %v\n", ip, application)
	c.JSON(http.StatusCreated, gin.H{"message": "Application added successfully"})
}

// UpdateApplication godoc
// @Summary Update an existing application
// @Description Update an application by its ID. Requires authentication.
// @Tags Applications
// @Accept json
// @Produce json
// @Param id path integer true "Application ID"
// @Param application body models.Application true "Updated application data"
// @Success 200 {object} defaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /applications/{id} [put]
// @Security ApiKeyAuth
func UpdateApplication(c *gin.Context) {
	ip := c.ClientIP()
	var application models.Application
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Client attempted to update application with ID %s. Error: Invalid application ID\n", ip, idStr)
		handleError(c, err)
		return
	}
	application.ID = uint(id)
	if err := c.ShouldBindJSON(&application); err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Client attempted to update application with ID %v using data %v. Error: Invalid input\n", ip, id, application)
		handleError(c, err)
		return
	}
	if err := service.UpdateApplication(application); err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Error updating application with ID %v\n", ip, id)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Successfully updated application with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Application updated successfully"})
}

// DeleteApplication godoc
// @Summary Delete an application
// @Description Soft delete an application by its ID. Requires authentication.
// @Tags Applications
// @Accept json
// @Produce json
// @Param id path integer true "Application ID"
// @Success 200 {object} defaultResponse
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
	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}

// GetSpecialistActivityReport
// @Summary      Get specialist activity report
// @Tags         Reports
// @Description  Get a report of how many vacancies a specialist has applied for
// @ID get-specialist-activity-report
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.SpecialistActivityReport  "Success"
// @Failure      500  {object} ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /reports/specialist-activity [get]
func GetSpecialistActivityReport(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetSpecialistActivityReport] Client IP: %s - Request to get specialist activity report\n", ip)
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	reports, err := service.GetSpecialistActivityReport(userID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetSpecialistActivityReport] Client IP: %s - Successfully retrieved specialist activity report\n", ip)
	c.JSON(http.StatusOK, reports)
}
