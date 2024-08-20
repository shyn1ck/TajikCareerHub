package controllers

import (
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllJobs(c *gin.Context) {
	jobs, err := service.GetAllJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve jobs"})
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func GetJobByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	job, err := service.GetJobByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve job"})
		}
		return
	}
	c.JSON(http.StatusOK, job)
}

func AddJob(c *gin.Context) {
	var job models.Job
	if err := c.BindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job data"})
		return
	}

	err := service.AddJob(job)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Job added successfully"})
}

func UpdateJob(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	var updatedJob models.Job
	if err := c.BindJSON(&updatedJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job data"})
		return
	}

	err = service.UpdateJob(uint(id), updatedJob)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job updated successfully"})
}

func DeleteJob(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	err = service.DeleteJob(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}

func FilterJobs(c *gin.Context) {
	location := c.Query("location")
	category := c.Query("category")

	jobs, err := service.FilterJobs(location, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter jobs"})
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func GetJobsBySalaryRange(c *gin.Context) {
	minSalary := c.Query("minSalary")
	maxSalary := c.Query("maxSalary")

	jobs, err := service.GetJobsBySalaryRange(minSalary, maxSalary)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func UpdateJobSalary(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	newSalary := c.Query("newSalary")
	err = service.UpdateJobSalary(uint(id), newSalary)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job salary updated successfully"})
}
