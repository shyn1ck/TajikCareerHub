package controllers

import (
	"TajikCareerHub/errs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllResumes godoc
// @Summary      Get all resumes
// @Description  Retrieves a list of resumes with optional filters such as search term, location, category, and minimum experience years.
// @Tags         Resumes
// @Accept       json
// @Produce      json
// @Param        search                query   string  false  "Search term"
// @Param        location              query   string  false  "Location"
// @Param        category              query   string  false  "Category"
// @Param        min-experience-years  query   int     false  "Minimum years of experience"
// @Success      200  {array}   models.SwagResume   "Success"  "List of resumes"
// @Failure      400  {object}  ErrorResponse  "Invalid request"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /resume [get]
func GetAllResumes(c *gin.Context) {
	ip := c.ClientIP()
	search := c.Query("search")
	location := c.Query("location")
	category := c.Query("category")
	minExperienceYearsStr := c.Query("min-experience-years")

	logger.Info.Printf("[controllers.GetAllResumes] Client IP: %s - Request to get resumes with search: %s, minExperienceYears: %s, location: %s, category: %s", ip, search, minExperienceYearsStr, location, category)

	var minExperienceYears int
	var err error
	if minExperienceYearsStr != "" {
		minExperienceYears, err = strconv.Atoi(minExperienceYearsStr)
		if err != nil {
			logger.Error.Printf("[controllers.GetAllResumes] Client IP: %s - Error converting minExperienceYears to int: %v", ip, err)
			handleError(c, errs.ErrIDIsNotCorrect)
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

	logger.Info.Printf("[controllers.GetAllResumes] Client IP: %s - Successfully retrieved resumes with search: %s, minExperienceYears: %d, location: %s, category: %s", ip, search, minExperienceYears, location, category)
	c.JSON(http.StatusOK, resumes)
}

// GetResumeByID godoc
// @Summary      Get resume by ID
// @Description  Get a specific resume by its ID
// @Tags         Resumes
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "Resume ID"
// @Success      200  {object}  models.SwagResume  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /resume/{id} [get]
func GetResumeByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.GetResumeByID] Client IP: %s - Error parsing resume ID: %s, Error: %v", ip, idStr, err)
		handleError(c, errs.ErrIDIsNotCorrect)
		return
	}

	logger.Info.Printf("[controllers.GetResumeByID] Client IP: %s - Request to get resume with ID: %d", ip, id)

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

	logger.Info.Printf("[controllers.GetResumeByID] Client IP: %s - Successfully retrieved resume with ID: %d", ip, id)
	c.JSON(http.StatusOK, resume)
}

// AddResume godoc
// @Summary      Add a new resume
// @Description  Adds a new resume to the system for the authenticated user.
// @Tags         Resumes
// @Accept       json
// @Produce      json
// @Param        resume  body   models.SwagResume  true  "Resume object"
// @Success      201  {object}  DefaultResponse  "Resume created successfully"
// @Failure      400  {object}  ErrorResponse  "Invalid request"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /resume [post]
func AddResume(c *gin.Context) {
	ip := c.ClientIP()
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.AddResume] Client IP: %s - Request to add resume with UserID: %d", ip, userID)

	var resume models.Resume
	if err := c.BindJSON(&resume); err != nil {
		logger.Error.Printf("[controllers.AddResume] Client IP: %s - Error parsing resume JSON: %v", ip, err)
		handleError(c, errs.ErrShouldBindJson)
		return
	}
	resume.UserID = userID
	err = service.AddResume(resume, userID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.AddResume] Client IP: %s - Successfully added resume: %v", ip, resume)
	c.JSON(http.StatusCreated, NewDefaultResponse("Resume added successfully"))
}

// UpdateResume godoc
// @Summary      Update an existing resume
// @Description  Update the details of an existing resume by its ID
// @Tags         Resumes
// @Accept       json
// @Produce      json
// @Param        id      path    int             true    "Resume ID"
// @Param        resume  body   models.Resume  true  "Updated resume object"
// @Success      200  {object}  DefaultResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID or request"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /resume/{id} [put]
func UpdateResume(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateResume] Client IP: %s - Error parsing resume ID: %s, Error: %v", ip, idStr, err)
		handleError(c, errs.ErrIDIsNotCorrect)
		return
	}

	logger.Info.Printf("[controllers.UpdateResume] Client IP: %s - Request to update resume with ID: %d", ip, id)

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	var updatedResume models.Resume
	if err := c.BindJSON(&updatedResume); err != nil {
		logger.Error.Printf("[controllers.UpdateResume] Client IP: %s - Error parsing resume JSON: %v", ip, err)
		handleError(c, err)
		return
	}

	err = service.UpdateResume(uint(id), updatedResume, userID)
	if err != nil {
		handleError(c, errs.ErrShouldBindJson)
		return
	}

	logger.Info.Printf("[controllers.UpdateResume] Client IP: %s - Successfully updated resume with ID: %d", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Resume updated successfully"))
}

// DeleteResume godoc
// @Summary      Delete a resume
// @Description  Delete a specific resume by its ID
// @Tags         Resumes
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "Resume ID"
// @Success      200  {object}  DefaultResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /resume/{id} [delete]
func DeleteResume(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteResume] Client IP: %s - Error parsing resume ID: %s, Error: %v", ip, idStr, err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.DeleteResume] Client IP: %s - Request to delete resume with ID: %d", ip, id)

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	err = service.DeleteResume(uint(id), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.DeleteResume] Client IP: %s - Successfully deleted resume with ID: %d", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Resume deleted successfully"))
}

// BlockResume godoc
// @Summary      Block a resume
// @Description  Blocks a resume by its ID. Requires the user to be authenticated.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "Resume ID"
// @Success      200   {object}  DefaultResponse  "Resume blocked successfully"
// @Failure      400   {object}  ErrorResponse  "Invalid resume ID"
// @Failure      401   {object}  ErrorResponse  "Unauthorized"
// @Failure      500   {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /resume/block/{id} [patch]
func BlockResume(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.BlockResume] Client IP: %s - Request to block resume with ID: %d", ip, id)

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	err = service.BlockResume(uint(id), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.BlockResume] Client IP: %s - Successfully blocked resume with ID: %d", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Resume blocked successfully"))
}

// UnblockResume godoc
// @Summary      Unblock a resume
// @Description  Unblocks a resume by its ID. Requires the user to be authenticated.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "Resume ID"
// @Success      200   {object}  DefaultResponse  "Resume unblocked successfully"
// @Failure      400   {object}  ErrorResponse  "Invalid resume ID"
// @Failure      401   {object}  ErrorResponse  "Unauthorized"
// @Failure      500   {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /resume/unblock/{id} [patch]
func UnblockResume(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.UnblockResume] Client IP: %s - Error parsing resume ID: %s, Error: %v", ip, idStr, err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UnblockResume] Client IP: %s - Request to unblock resume with ID: %d", ip, id)

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	err = service.UnblockResume(uint(id), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UnblockResume] Client IP: %s - Resume with ID: %d unblocked successfully", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Resume unblocked successfully"))
}
