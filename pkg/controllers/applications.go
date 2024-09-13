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
		logger.Info.Printf("[controllers.GetAllApplications] Client IP: %s - Error retrieving all applications. Error: %v\n", ip, err)
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
		logger.Info.Printf("[controllers.GetApplicationByID] Client IP: %s - Invalid application ID %s. Error: %v\n", ip, idStr, err)
		handleError(c, err)
		return
	}

	application, err := service.GetApplicationByID(uint(id))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationByID] Client IP: %s - Error retrieving application with ID %v. Error: %v\n", ip, id, err)
		handleError(c, err)
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
		logger.Info.Printf("[controllers.GetApplicationsByUserID] Client IP: %s - Invalid user ID %s. Error: %v\n", ip, userIDStr, err)
		handleError(c, err)
		return
	}

	applications, err := service.GetApplicationsByUserID(uint(userID))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByUserID] Client IP: %s - Error retrieving applications for user ID %v. Error: %v\n", ip, userID, err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetApplicationsByUserID] Client IP: %s - Successfully retrieved applications for user ID %v\n", ip, userID)
	c.JSON(http.StatusOK, applications)
}

func GetApplicationsByVacancyID(c *gin.Context) {
	ip := c.ClientIP()
	vacancyIDStr := c.Param("vacancyID")
	vacancyID, err := strconv.ParseUint(vacancyIDStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByVacancyID] Client IP: %s - Invalid vacancy ID %s. Error: %v\n", ip, vacancyIDStr, err)
		handleError(c, err)
		return
	}

	applications, err := service.GetApplicationsByVacancyID(uint(vacancyID))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByVacancyID] Client IP: %s - Error retrieving applications for vacancy ID %v. Error: %v\n", ip, vacancyID, err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetApplicationsByVacancyID] Client IP: %s - Successfully retrieved applications for vacancy ID %v\n", ip, vacancyID)
	c.JSON(http.StatusOK, applications)
}

func ApplyForVacancy(c *gin.Context) {
	ip := c.ClientIP()
	userIDStr := c.Param("user_id")
	vacancyIDStr := c.Param("vacancy_id")
	resumeIDStr := c.Param("resume_id")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.ApplyForVacancy] Client IP: %s - Invalid user ID %s. Error: %v\n", ip, userIDStr, err)
		handleError(c, err)
		return
	}
	vacancyID, err := strconv.ParseUint(vacancyIDStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.ApplyForVacancy] Client IP: %s - Invalid vacancy ID %s. Error: %v\n", ip, vacancyIDStr, err)
		handleError(c, err)
		return
	}
	resumeID, err := strconv.ParseUint(resumeIDStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.ApplyForVacancy] Client IP: %s - Invalid resume ID %s. Error: %v\n", ip, resumeIDStr, err)
		handleError(c, err)
		return
	}

	err = service.ApplyForVacancy(uint(userID), uint(vacancyID), uint(resumeID))
	if err != nil {
		logger.Info.Printf("[controllers.ApplyForVacancy] Client IP: %s - Error applying for vacancy %v by user %v. Error: %v\n", ip, vacancyID, userID, err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.ApplyForVacancy] Client IP: %s - Successfully applied for vacancy %v by user %v\n", ip, vacancyID, userID)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully applied for vacancy"})
}

func UpdateApplication(c *gin.Context) {
	ip := c.ClientIP()
	var application models.Application
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Invalid application ID %s. Error: %v\n", ip, idStr, err)
		handleError(c, err)
		return
	}
	application.ID = uint(id)

	if err := c.ShouldBindJSON(&application); err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Failed to bind JSON for application ID %v with data %v. Error: %v\n", ip, id, application, err)
		handleError(c, err)
		return
	}

	if err := service.UpdateApplication(application); err != nil {
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Error updating application ID %v. Error: %v\n", ip, id, err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Successfully updated application ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Application updated successfully"})
}

func DeleteApplication(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.DeleteApplication] Client IP: %s - Invalid application ID %s. Error: %v\n", ip, idStr, err)
		handleError(c, err)
		return
	}

	if err := service.DeleteApplication(uint(id)); err != nil {
		logger.Info.Printf("[controllers.DeleteApplication] Client IP: %s - Error soft deleting application ID %v. Error: %v\n", ip, id, err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.DeleteApplication] Client IP: %s - Successfully soft deleted application ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}

func UpdateApplicationStatus(c *gin.Context) {
	ip := c.ClientIP()
	var statusUpdate models.ApplicationStatusUpdate
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Invalid application ID %s. Error: %v\n", ip, idStr, err)
		handleError(c, err)
		return
	}
	statusUpdate.ApplicationID = uint(id)

	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Failed to bind JSON for application ID %v with data %v. Error: %v\n", ip, id, statusUpdate, err)
		handleError(c, err)
		return
	}

	if err := service.UpdateApplicationStatus(statusUpdate); err != nil {
		logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Error updating status for application ID %v. Error: %v\n", ip, id, err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateApplicationStatus] Client IP: %s - Successfully updated status for application ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Application status updated successfully"})
}
