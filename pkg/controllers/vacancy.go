package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllVacancies
// @Summary Retrieve all vacancies with filters
// @Tags Vacancies
// @Description Get a list of all vacancies with optional filters such as search, salary range, location, category, and sort order.
// @ID get-all-vacancies
// @Accept json
// @Produce json
// @Param userID query integer true "User ID to check if the user is blocked"
// @Param search query string false "Search keyword for filtering vacancies"
// @Param minSalary query integer false "Minimum salary for filtering vacancies"
// @Param maxSalary query integer false "Maximum salary for filtering vacancies"
// @Param location query string false "Location for filtering vacancies"
// @Param category query string false "Category for filtering vacancies"
// @Param sort query string false "Sorting order for vacancies"
// @Success 200 {array}  models.Vacancy "Successfully retrieved list of vacancies"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Security     ApiKeyAuth
// @Router /vacancy [get]
func GetAllVacancies(c *gin.Context) {
	ip := c.ClientIP()
	search := c.Query("search")
	location := c.Query("location")
	category := c.Query("category")
	minSalaryStr := c.Query("min-salary")
	maxSalaryStr := c.Query("max-salary")
	sort := c.Query("sort")
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetAllVacancies] Client IP: %s - Request to get vacancies with keyword: %s, minSalary: %s, maxSalary: %s, location: %s, category: %s, sort: %s\n", ip, search, minSalaryStr, maxSalaryStr, location, category, sort)

	var minSalary, maxSalary int
	if minSalaryStr != "" {
		minSalary, err = strconv.Atoi(minSalaryStr)
		if err != nil {
			logger.Error.Printf("[controllers.GetAllVacancies] Error converting minSalary to int: %s", err.Error())
			handleError(c, err)
			return
		}
	}

	if maxSalaryStr != "" {
		maxSalary, err = strconv.Atoi(maxSalaryStr)
		if err != nil {
			logger.Error.Printf("[controllers.GetAllVacancies] Error converting maxSalary to int: %s", err.Error())
			handleError(c, err)
			return
		}
	}

	vacancies, err := service.GetAllVacancies(userID, search, minSalary, maxSalary, location, category, sort)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetAllVacancies] Client IP: %s - Successfully retrieved vacancies with keyword: %s, minSalary: %d, maxSalary: %d, location: %s, category: %s, sort: %s\n", ip, search, minSalary, maxSalary, location, category, sort)
	c.JSON(http.StatusOK, vacancies)
}

// GetVacancyByID
// @Summary Retrieve a specific vacancy by ID
// @Tags Vacancies
// @Description Get details of a single vacancy by its ID.
// @ID get-vacancy-by-id
// @Accept json
// @Produce json
// @Param userID query integer true "User ID to check if the user is blocked"
// @Param vacancyID path integer true "ID of the vacancy to retrieve"
// @Success 200 {object} models.Vacancy "Successfully retrieved vacancy"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Vacancy Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Security     ApiKeyAuth
// @Router /vacancy/{vacancyID} [get]
func GetVacancyByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("vacancyID")
	logger.Info.Printf("[controllers.GetVacancyByID] Client IP: %s - Request to get vacancy by ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.GetVacancyByID] Error converting id to int: %s", err.Error())
		handleError(c, err)
		return
	}
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}
	vacancy, err := service.GetVacancyByID(userID, uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetVacancyByID] Client IP: %s - Successfully retrieved vacancy with ID %v.\n", ip, id)
	c.JSON(http.StatusOK, vacancy)
}

// AddVacancy
// @Summary Create a new vacancy
// @Tags Vacancies
// @Description Add a new vacancy with the provided details.
// @ID add-vacancy
// @Accept json
// @Produce json
// @Param userID query integer true "User ID to check if the user is blocked"
// @Param vacancy body models.SwagVacancy true "Vacancy object to be added"
// @Success 201 {object} DefaultResponse "Vacancy created successfully"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 403 {object} ErrorResponse "ErrPermissionDenied"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Security     ApiKeyAuth
// @Router /vacancy [post]
func AddVacancy(c *gin.Context) {
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	var vacancy models.Vacancy
	if err := c.BindJSON(&vacancy); err != nil {
		logger.Error.Printf("[controllers.AddVacancy] Error parsing request: %s", err.Error())
		handleError(c, err)
		return
	}

	vacancy.UserID = userID

	err = service.AddVacancy(userID, vacancy)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, NewDefaultResponse("Vacancy added successfully"))
}

// UpdateVacancy
// @Summary Update an existing vacancy
// @Tags Vacancies
// @Description Update an existing vacancy by its ID.
// @ID update-vacancy
// @Accept json
// @Produce json
// @Param userID query integer true "User ID to check if the user is blocked"
// @Param vacancyID path integer true "ID of the vacancy to update"
// @Param vacancy body models.SwagVacancy true "Updated vacancy object"
// @Success 200 {object} DefaultResponse "Vacancy updated successfully"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Vacancy Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Security     ApiKeyAuth
// @Router /vacancy/{vacancyID} [put]
func UpdateVacancy(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("vacancyID")
	logger.Info.Printf("[controllers.UpdateVacancy] Client IP: %s - Request to update vacancy with ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateVacancy] Error converting id to int: %s", err.Error())
		handleError(c, err)
		return
	}

	var updatedVacancy models.Vacancy
	if err := c.BindJSON(&updatedVacancy); err != nil {
		logger.Error.Printf("[controllers.UpdateVacancy] Error parsing request: %s", err.Error())
		handleError(c, err)
		return
	}

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	err = service.UpdateVacancy(userID, uint(id), updatedVacancy)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateVacancy] Client IP: %s - Vacancy with ID %v updated successfully.\n", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Vacancy updated successfully"))
}

// DeleteVacancy
// @Summary Delete a vacancy
// @Tags Vacancies
// @Description Soft delete a specific vacancy by its ID.
// @ID delete-vacancy
// @Accept json
// @Produce json
// @Param userID query integer true "User ID to check if the user is blocked"
// @Param vacancyID path integer true "ID of the vacancy to delete"
// @Success 204 {object} DefaultResponse "Vacancy deleted successfully"
// @Failure 404 {object} ErrorResponse "Vacancy Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Security     ApiKeyAuth
// @Router /vacancy/{vacancyID} [delete]
func DeleteVacancy(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("vacancyID")
	logger.Info.Printf("[controllers.DeleteVacancy] Client IP: %s - Request to delete vacancy with ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteVacancy] Error converting id to int: %s", err.Error())
		handleError(c, err)
		return
	}
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	err = service.DeleteVacancy(userID, uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.DeleteVacancy] Client IP: %s - Vacancy with ID %v deleted successfully.\n", ip, id)
	c.JSON(http.StatusNoContent, NewDefaultResponse("Vacancy deleted successfully"))
}

func GetVacancyReport(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetVacancyReport]: Client with IP %s requested to get vacancy report.\n", ip)
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	reports, err := service.GetVacancyReport(userID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetVacancyReport] Client IP: %s - Successfully retrieved vacancy report.\n", ip)
	c.JSON(http.StatusOK, reports)
}
