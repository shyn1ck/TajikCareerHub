package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllJobs(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetAllJobs] Client IP: %s - Request to get all jobs.\n", ip)
	jobs, err := service.GetAllJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve jobs"})
		return
	}
	logger.Info.Printf("[controllers.GetAllJobs] Client IP: %s - Successfully retrieved all jobs.\n", ip)
	c.JSON(http.StatusOK, jobs)
}

func GetJobByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.GetJobByID] Client IP: %s - Request to get job by ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
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
	logger.Info.Printf("[controllers.GetJobByID] Client IP: %s - Successfully retrieved job with ID %v.\n", ip, id)
	c.JSON(http.StatusOK, job)
}

func AddJob(c *gin.Context) {
	ip := c.ClientIP()
	var job models.Job
	if err := c.BindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job data"})
		return
	}

	logger.Info.Printf("[controllers.AddJob] Client IP: %s - Request to add job: %v\n", ip, job)
	err := service.AddJob(job)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("[controllers.AddJob] Client IP: %s - Job added successfully: %v\n", ip, job)
	c.JSON(http.StatusCreated, gin.H{"message": "Job added successfully"})
}

func UpdateJob(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.UpdateJob] Client IP: %s - Request to update job with ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
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
	logger.Info.Printf("[controllers.UpdateJob] Client IP: %s - Job with ID %v updated successfully.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Job updated successfully"})
}

func DeleteJob(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.DeleteJob] Client IP: %s - Request to delete job with ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	err = service.DeleteJob(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job"})
		return
	}
	logger.Info.Printf("[controllers.DeleteJob] Client IP: %s - Job with ID %v deleted successfully.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}

func FilterJobs(c *gin.Context) {
	ip := c.ClientIP()
	location := c.Query("location")
	category := c.Query("category")
	logger.Info.Printf("[controllers.FilterJobs] Client IP: %s - Request to filter jobs by location: %s and category: %s\n", ip, location, category)

	jobs, err := service.FilterJobs(location, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter jobs"})
		return
	}
	logger.Info.Printf("[controllers.FilterJobs] Client IP: %s - Jobs filtered by location %s and category %s successfully.\n", ip, location, category)
	c.JSON(http.StatusOK, jobs)
}

func GetJobsBySalaryRange(c *gin.Context) {
	ip := c.ClientIP()
	minSalary := c.Query("minSalary")
	maxSalary := c.Query("maxSalary")
	logger.Info.Printf("[controllers.GetJobsBySalaryRange] Client IP: %s - Request to get jobs by salary range: %s to %s\n", ip, minSalary, maxSalary)

	jobs, err := service.GetJobsBySalaryRange(minSalary, maxSalary)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("[controllers.GetJobsBySalaryRange] Client IP: %s - Successfully retrieved jobs by salary range %s to %s.\n", ip, minSalary, maxSalary)
	c.JSON(http.StatusOK, jobs)
}

func UpdateJobSalary(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	newSalary := c.Query("newSalary")
	logger.Info.Printf("[controllers.UpdateJobSalary] Client IP: %s - Request to update salary for job with ID: %s to new salary: %s\n", ip, idStr, newSalary)

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	err = service.UpdateJobSalary(uint(id), newSalary)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("[controllers.UpdateJobSalary] Client IP: %s - Job salary for ID %v updated to %s successfully.\n", ip, id, newSalary)
	c.JSON(http.StatusOK, gin.H{"message": "Job salary updated successfully"})
}
