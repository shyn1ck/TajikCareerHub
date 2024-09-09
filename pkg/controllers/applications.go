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

func GetApplicationsByJobID(c *gin.Context) {
	ip := c.ClientIP()
	jobIDStr := c.Param("jobID")
	jobID, err := strconv.ParseUint(jobIDStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByJobID] Client IP: %s - Client requested applications for job ID %s. Error: Invalid job ID\n", ip, jobIDStr)
		handleError(c, err)
		return
	}

	applications, err := service.GetApplicationsByJobID(uint(jobID))
	if err != nil {
		logger.Info.Printf("[controllers.GetApplicationsByJobID] Client IP: %s - Client requested applications for job ID %v. Error retrieving applications\n", ip, jobID)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetApplicationsByJobID] Client IP: %s - Successfully retrieved applications for job ID %v\n", ip, jobID)
	c.JSON(http.StatusOK, applications)
}

func ApplyForJob(c *gin.Context) {
	ip := c.ClientIP()
	userIDStr := c.Param("user_id")
	jobIDStr := c.Param("job_id")
	resumeIDStr := c.Param("resume_id")
	logger.Info.Printf("[controllers.ApplyForJob] Client IP: %s - Request to apply for job with user ID: %s, job ID: %s, resume ID: %s\n", ip, userIDStr, jobIDStr, resumeIDStr)
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	jobID, err := strconv.ParseUint(jobIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	resumeID, err := strconv.ParseUint(resumeIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	err = service.ApplyForVacancy(uint(userID), uint(jobID), uint(resumeID))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.ApplyForJob] Client IP: %s - Successfully applied for job %v by user %v.\n", ip, jobID, userID)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully applied for job"})
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

}
