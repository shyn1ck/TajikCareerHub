package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
