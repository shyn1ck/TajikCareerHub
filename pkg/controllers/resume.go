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
	minExperienceYearsStr := c.Query("min_experience_years")
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	resume, err := service.GetResumeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume not found"})
		return
	}
	c.JSON(http.StatusOK, resume)
}

func AddResume(c *gin.Context) {
	var resume models.Resume
	if err := c.ShouldBindJSON(&resume); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := service.AddResume(resume); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add resume"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Resume added successfully"})
}

func UpdateResume(c *gin.Context) {
	var resume models.Resume
	if err := c.ShouldBindJSON(&resume); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := service.UpdateResume(resume); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update resume"})
		return
	}
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
