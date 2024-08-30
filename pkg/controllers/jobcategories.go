package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetJobCategoryByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	category, err := service.GetJobCategoryByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetJobCategoryByID] Client IP: %s - Successfully retrieved job category with ID %v\n", ip, id)
	c.JSON(http.StatusOK, category)
}

func GetAllJobCategories(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetAllJobCategories] Client IP: %s - Client requested all job categories\n", ip)

	categories, err := service.GetAllJobCategories()
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetAllJobCategories] Client IP: %s - Successfully retrieved all job categories\n", ip)
	c.JSON(http.StatusOK, categories)
}

func CreateJobCategory(c *gin.Context) {
	ip := c.ClientIP()
	var category models.JobCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		handleError(c, err)
		return
	}

	if err := service.CreateJobCategory(category); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.CreateJobCategory] Client IP: %s - Successfully created job category with data %v\n", ip, category)
	c.JSON(http.StatusCreated, gin.H{"message": "Job category created successfully"})
}

func UpdateJobCategory(c *gin.Context) {
	ip := c.ClientIP()
	var category models.JobCategory
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	category.ID = uint(id)

	if err := c.ShouldBindJSON(&category); err != nil {
		handleError(c, err)
		return
	}

	if err := service.UpdateJobCategory(category); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateJobCategory] Client IP: %s - Successfully updated job category with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Job category updated successfully"})
}

func DeleteJobCategory(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	if err := service.DeleteJobCategory(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.DeleteJobCategory] Client IP: %s - Successfully soft deleted job category with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Job category deleted successfully"})
}
