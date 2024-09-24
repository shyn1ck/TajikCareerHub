package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"TajikCareerHub/utils/errs"
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
// @Param search query string false "Search keyword for filtering vacancies"
// @Param minSalary query integer false "Minimum salary for filtering vacancies"
// @Param maxSalary query integer false "Maximum salary for filtering vacancies"
// @Param location query string false "Location for filtering vacancies"
// @Param category query string false "Category for filtering vacancies"
// @Param sort query string false "Sorting order for vacancies"
// @Success 200 {array}    models.Vacancy "Successfully retrieved list of vacancies"
// @Failure 400 {object}   ErrorResponse "Bad Request"
// @Failure 403  {object}  ErrorResponse 	 "Access Denied"
// @Failure 500 {object}   ErrorResponse "Internal Server Error"
// @Security     ApiKeyAuth
// @Router /vacancies [get]
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
			handleError(c, errs.ErrIDIsNotCorrect)
			return
		}
	}

	if maxSalaryStr != "" {
		maxSalary, err = strconv.Atoi(maxSalaryStr)
		if err != nil {
			logger.Error.Printf("[controllers.GetAllVacancies] Error converting maxSalary to int: %s", err.Error())
			handleError(c, errs.ErrIncorrectInput)
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
// @Param vacancyID path integer true "ID of the vacancy to retrieve"
// @Success 200 {object} models.SwagVacancy "Successfully retrieved vacancy"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 403  {object}  ErrorResponse 	 "Access Denied"
// @Failure 404 {object} ErrorResponse "Vacancy Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Security     ApiKeyAuth
// @Router /vacancies/{vacancyID} [get]
func GetVacancyByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("vacancyID")
	logger.Info.Printf("[controllers.GetVacancyByID] Client IP: %s - Request to get vacancy by ID: %s\n", ip, idStr)

	if idStr == "" {
		logger.Error.Printf("[controllers.GetVacancyByID] Vacancy ID is missing in the request.")
		handleError(c, errs.ErrIDIsNotProvided)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.GetVacancyByID] Error converting id to int: %s", err.Error())
		handleError(c, errs.ErrIDIsNotCorrect)
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
// @Router /vacancies [post]
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
// @Failure 403 {object} ErrorResponse 	 "Access Denied"
// @Failure 404 {object} ErrorResponse "Vacancy Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Security     ApiKeyAuth
// @Router /vacancies/{vacancyID} [put]
func UpdateVacancy(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("vacancyID")
	logger.Info.Printf("[controllers.UpdateVacancy] Client IP: %s - Request to update vacancy with ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateVacancy] Error converting id to int: %s", err.Error())
		handleError(c, errs.ErrIDIsNotCorrect)
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
// @Failure 403 {object} ErrorResponse 	 "Access Denied"
// @Failure 404 {object} ErrorResponse "Vacancy Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Security     ApiKeyAuth
// @Router /vacancies/{vacancyID} [delete]
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

// GetVacancyReportByID godoc
// @Summary Get report for a specific vacancy
// @Tags Reports
// @Description Get a report of how many people viewed or applied to a specific vacancy
// @ID get-vacancy-report-by-id
// @Accept json
// @Produce json
// @Param id path uint true "Vacancy ID"
// @Success 200 {object} models.VacancyReport
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Forbidden access"
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /activities/vacancy/{id} [get]
func GetVacancyReportByID(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetVacancyReportByID] Client IP: %s - Request to get report for vacancy ID %s\n", ip, c.Param("id"))

	vacancyIDStr := c.Param("id")
	vacancyID, err := strconv.ParseUint(vacancyIDStr, 10, 32)
	if err != nil {
		handleError(c, errs.ErrIDIsNotCorrect)
		logger.Error.Printf("[controllers.GetVacancyReportByID] Error converting id to int: %s", err.Error())
		return
	}

	report, err := service.GetVacancyReportByID(uint(vacancyID))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetVacancyReportByID] Client IP: %s - Successfully retrieved report for vacancy ID %d\n", ip, vacancyID)
	c.JSON(http.StatusOK, report)
}

// BlockVacancy godoc
// @Summary      Block a vacancy
// @Description  Blocks a vacancy by its ID. Requires the user to be authenticated.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "Vacancy ID"
// @Success      200   {object}  DefaultResponse  "Vacancy blocked successfully"
// @Failure      400   {object}  ErrorResponse  "Invalid vacancy ID"
// @Failure      401   {object}  ErrorResponse  "Unauthorized"
// @Failure      403   {object}  ErrorResponse  "Access Denied"
// @Failure      500   {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /vacancies/block/{id} [patch]
func BlockVacancy(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.BlockVacancy] Client IP: %s - Error parsing vacancy ID: %s, Error: %v", ip, idStr, err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.BlockVacancy] Client IP: %s - Request to block vacancy with ID: %d", ip, id)

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}
	role, err := service.GetRoleFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	err = service.BlockVacancy(uint(id), userID, role)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.BlockVacancy] Client IP: %s - Successfully blocked vacancy with ID: %d", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Vacancy blocked successfully"))
}

// UnblockVacancy godoc
// @Summary      Unblock a vacancy
// @Description  Unblocks a vacancy by its ID. Requires the user to be authenticated.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "Vacancy ID"
// @Success      200   {object}  DefaultResponse  "Vacancy unblocked successfully"
// @Failure      400   {object}  ErrorResponse  "Invalid vacancy ID"
// @Failure      401   {object}  ErrorResponse  "Unauthorized"
// @Failure      403   {object}  ErrorResponse  "Access Denied"
// @Failure      500   {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /vacancies/unblock/{id} [patch]
func UnblockVacancy(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.UnblockVacancy] Client IP: %s - Error parsing vacancy ID: %s, Error: %v", ip, idStr, err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UnblockVacancy] Client IP: %s - Request to unblock vacancy with ID: %d", ip, id)

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}
	role, err := service.GetRoleFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	err = service.UnblockVacancy(uint(id), userID, role)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UnblockVacancy] Client IP: %s - Vacancy with ID: %d unblocked successfully", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Vacancy unblocked successfully"))
}
