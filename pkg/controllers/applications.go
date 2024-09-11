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
		logger.Info.Printf("[controllers.GetAllApplications] Client IP: %s - Client requested all applications.\n", ip)
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

	application, err := service.GetApplicationByID(uint(id))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationByID] Client IP: %s - Client requested application with ID %v. Error retrieving application\n", ip, id)
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
		logger.Info.Printf("[controllers.GetApplicationsByUserID] Client IP: %s - Client requested applications for user ID %s. Error: Invalid user ID\n", ip, userIDStr)
		handleError(c, err)
		return
	}

	applications, err := service.GetApplicationsByUserID(uint(userID))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByUserID] Client IP: %s - Client requested applications for user ID %v. Error retrieving applications\n", ip, userID)
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
		logger.Info.Printf("[controllers.GetApplicationsByVacancyID] Client IP: %s - Client requested applications for vacancy ID %s. Error: Invalid vacancy ID\n", ip, vacancyIDStr)
		handleError(c, err)
		return
	}

	applications, err := service.GetApplicationsByVacancyID(uint(vacancyID))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByVacancyID] Client IP: %s - Client requested applications for vacancy ID %v. Error retrieving applications\n", ip, vacancyID)
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
	logger.Info.Printf("[controllers.ApplyForVacancy] Client IP: %s - Request to apply for vacancy with user ID: %s, vacancy ID: %s, resume ID: %s\n", ip, userIDStr, vacancyIDStr, resumeIDStr)
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	vacancyID, err := strconv.ParseUint(vacancyIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	resumeID, err := strconv.ParseUint(resumeIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	err = service.ApplyForVacancy(uint(userID), uint(vacancyID), uint(resumeID))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.ApplyForVacancy] Client IP: %s - Successfully applied for vacancy %v by user %v.\n", ip, vacancyID, userID)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully applied for vacancy"})
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
		logger.Info.Printf("[controllers.UpdateApplication] Client IP: %s - Client attempted to update application with ID %v using data %v. Error updating application\n", ip, id, application)
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

	if err := service.DeleteApplication(uint(id)); err != nil {
		logger.Info.Printf("[controllers.DeleteApplication] Client IP: %s - Client attempted to delete application with ID %v. Error soft deleting application\n", ip, id)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.DeleteApplication] Client IP: %s - Successfully soft deleted application with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}

func GetUserApplicationActivity(c *gin.Context) {

}

func GetJobApplications(c *gin.Context) {

}

func UpdateApplicationStatus(c *gin.Context) {

}

func GetJobReport(c *gin.Context) {
	// TODO: Implement this function
}
