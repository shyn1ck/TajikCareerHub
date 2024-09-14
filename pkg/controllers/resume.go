package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllResumes godoc
// @Summary      Get all resumes
// @Description  Get resumes with optional filters: search term, location, category, and minimum experience years
// @Tags         resumes
// @Accept       json
// @Produce      json
// @Param        search       query   string  false  "Search term"
// @Param        location     query   string  false  "Location"
// @Param        category     query   string  false  "Category"
// @Param        min-experience-years  query   int     false  "Minimum years of experience"
// @Success      200  {array}   models.Resume  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid request"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Router       /resumes [get]
func GetAllResumes(c *gin.Context) {
	ip := c.ClientIP()
	search := c.Query("search")
	location := c.Query("location")
	category := c.Query("category")
	minExperienceYearsStr := c.Query("min-experience-years")

	logger.Info.Printf("[controllers.GetAllResumes] Client IP: %s - Request to get resumes with search: %s, minExperienceYears: %s, location: %s, category: %s\n", ip, search, minExperienceYearsStr, location, category)

	var minExperienceYears int
	var err error
	if minExperienceYearsStr != "" {
		minExperienceYears, err = strconv.Atoi(minExperienceYearsStr)
		if err != nil {
			handleError(c, err)
			return
		}
	}

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	resumes, err := service.GetAllResumes(search, minExperienceYears, location, category, userID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetAllResumes] Client IP: %s - Successfully retrieved resumes with search: %s, minExperienceYears: %d, location: %s, category: %s\n", ip, search, minExperienceYears, location, category)
	c.JSON(http.StatusOK, resumes)
}

// GetResumeByID godoc
// @Summary      Get resume by ID
// @Description  Get a specific resume by its ID
// @Tags         resumes
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "Resume ID"
// @Success      200  {object}  models.Resume  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Router       /resumes/{id} [get]
func GetResumeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	resume, err := service.GetResumeByID(uint(id), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resume)
}

// AddResume godoc
// @Summary      Add a new resume
// @Description  Add a new resume to the system
// @Tags         resumes
// @Accept       json
// @Produce      json
// @Param        resume  body    models.Resume  true  "Resume object"
// @Success      201  {object}  defaultResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid request"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Router       /resumes [post]
func AddResume(c *gin.Context) {
	ip := c.ClientIP()
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	var resume models.Resume
	if err := c.BindJSON(&resume); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.AddResume] Client IP: %s - Request to add resume: %v\n", ip, resume)
	resume.UserID = userID

	err = service.AddResume(resume, userID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.AddResume] Client IP: %s - Resume added successfully: %v\n", ip, resume)
	c.JSON(http.StatusCreated, newDefaultResponse("Resume added successfully"))
}

// UpdateResume godoc
// @Summary      Update an existing resume
// @Description  Update the details of an existing resume by its ID
// @Tags         resumes
// @Accept       json
// @Produce      json
// @Param        id      path    int             true    "Resume ID"
// @Param        resume  body    models.Resume  true  "Updated resume object"
// @Success      200  {object}  defaultResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID or request"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Router       /resumes/{id} [put]
func UpdateResume(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.UpdateResume] Client IP: %s - Request to update resume with ID: %s\n", ip, idStr)

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	var updatedResume models.Resume
	if err := c.BindJSON(&updatedResume); err != nil {
		handleError(c, err)
		return
	}

	err = service.UpdateResume(uint(id), updatedResume, userID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateResume] Client IP: %s - Resume with ID %v updated successfully.\n", ip, id)
	c.JSON(http.StatusOK, newDefaultResponse("Resume updated successfully"))
}

// DeleteResume godoc
// @Summary      Delete a resume
// @Description  Delete a specific resume by its ID
// @Tags         resumes
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "Resume ID"
// @Success      200  {object}  defaultResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Router       /resumes/{id} [delete]
func DeleteResume(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse("Invalid ID"))
		return
	}
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	if err := service.DeleteResume(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, newErrorResponse("Failed to delete resume"))
		return
	}

	c.JSON(http.StatusOK, newDefaultResponse("Resume deleted successfully"))
}
