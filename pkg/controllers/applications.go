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
	applications, err := service.GetAllApplications()
	if err != nil {
		logger.Info.Printf("[controllers.GetAllApplications] Client IP: %s - Client requested all applications. Successfully retrieved all applications.\n", ip)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve applications"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}

	application, err := service.GetApplicationByID(uint(id))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationByID] Client IP: %s - Client requested application with ID %v. Error retrieving application\n", ip, id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	logger.Info.Printf("[controllers.GetApplicationByID] Client IP: %s - Successfully retrieved application with ID %v\n", ip, id)
	c.JSON(http.StatusOK, application)
}

func GetApplicationsByUserID(c *gin.Context) {
	ip := c.ClientIP()
	userIDStr := c.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByUserID] Client IP: %s - Client requested applications for user ID %s. Error: Invalid user ID\n", ip, userIDStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	applications, err := service.GetApplicationsByUserID(uint(userID))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByUserID] Client IP: %s - Client requested applications for user ID %v. Error retrieving applications\n", ip, userID)
		c.JSON(http.StatusNotFound, gin.H{"error": "No applications found for user"})
		return
	}

	logger.Info.Printf("[controllers.GetApplicationsByUserID] Client IP: %s - Successfully retrieved applications for user ID %v\n", ip, userID)
	c.JSON(http.StatusOK, applications)
}

func GetApplicationsByJobID(c *gin.Context) {
	ip := c.ClientIP()
	jobIDStr := c.Param("jobID")
	jobID, err := strconv.ParseUint(jobIDStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByJobID] Client IP: %s - Client requested applications for job ID %s. Error: Invalid job ID\n", ip, jobIDStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	applications, err := service.GetApplicationsByJobID(uint(jobID))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByJobID] Client IP: %s - Client requested applications for job ID %v. Error retrieving applications\n", ip, jobID)
		c.JSON(http.StatusNotFound, gin.H{"error": "No applications found for job"})
		return
	}

	logger.Info.Printf("[controllers.GetApplicationsByJobID] Client IP: %s - Successfully retrieved applications for job ID %v\n", ip, jobID)
	c.JSON(http.StatusOK, applications)
}

func AddApplication(c *gin.Context) {
	ip := c.ClientIP()
	var application models.Application
	if err := c.ShouldBindJSON(&application); err != nil {
		logger.Info.Printf("[controllers.AddApplication] Client IP: %s - Client attempted to add application with data %v. Error: Invalid input\n", ip, application)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := service.AddApplication(application); err != nil {
		logger.Info.Printf("[controllers.AddApplication] Client IP: %s - Client attempted to add application with data %v. Error adding application\n", ip, application)
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to add application"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}
	application.ID = uint(id)

	if err := c.ShouldBindJSON(&application); err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Client attempted to update application with ID %v using data %v. Error: Invalid input\n", ip, id, application)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := service.UpdateApplication(application); err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Client attempted to update application with ID %v using data %v. Error updating application\n", ip, id, application)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update application"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}

	if err := service.DeleteApplication(uint(id)); err != nil {
		logger.Info.Printf("[controllers.DeleteApplication] Client IP: %s - Client attempted to delete application with ID %v. Error soft deleting application\n", ip, id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete application"})
		return
	}

	logger.Info.Printf("[controllers.DeleteApplication] Client IP: %s - Successfully soft deleted application with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}
