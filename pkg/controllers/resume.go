package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllResumes(c *gin.Context) {
	ip := c.ClientIP()
	search := c.Query("search")
	location := c.Query("location")
	category := c.Query("category")
	minExperienceYearsStr := c.Query("min-experience-years")
	logger.Info.Printf("[controllers.GetAllResumes] Client IP: %s - Request to get resumes with search: %s, minExperienceYearsStr: %s, location: %s, category: %s\n", ip, search, minExperienceYearsStr, location, category)
	var minExperienceYears int
	var err error
	if minExperienceYearsStr != "" {
		minExperienceYears, err = strconv.Atoi(minExperienceYearsStr)
		if err != nil {
			handleError(c, err)
			return
		}
	}

	resumes, err := service.GetAllResume(search, minExperienceYears, location, category)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetAllResumes] Client IP: %s - Successufuly to get resumes with search: %s, minExperienceYearsStr: %s, location: %s, category: %s\n", ip, search, minExperienceYearsStr, location, category)
	c.JSON(http.StatusOK, resumes)
}

func GetResumeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	resume, err := service.GetResumeByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resume)
}

func AddResume(c *gin.Context) {
	var resume models.Resume
	if err := c.ShouldBindJSON(&resume); err != nil {
		handleError(c, err)
		return
	}
	if err := service.AddResume(resume); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Resume added successfully"})
}

func UpdateResume(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.UpdateResume] Client IP: %s - Request to update resume with ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	var updatedResume models.Resume
	if err := c.BindJSON(&updatedResume); err != nil {
		handleError(c, err)
		return
	}
	err = service.UpdateResume(uint(id), updatedResume)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateResume] Client IP: %s - Resume with ID %v updated successfully.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Resume updated successfully"})
}

func DeleteResume(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := service.DeleteResume(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete resume"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Resume deleted successfully"})
}
