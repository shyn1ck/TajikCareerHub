package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllCompanies godoc
// @Summary Get all companies
// @Description Retrieve a list of all companies. No authentication required.
// @Tags Companies
// @Accept json
// @Produce json
// @Success 200 {array} models.Company
// @Failure 500 {object} ErrorResponse
// @Router /company [get]
func GetAllCompanies(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetAllCompanies] Client IP: %s - Request to get all companies\n", ip)

	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	companies, err := service.GetAllCompanies(userID)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetAllCompanies] Client IP: %s - Successfully retrieved all companies\n", ip)
	c.JSON(http.StatusOK, companies)
}

// GetCompanyByID godoc
// @Summary Get company by ID
// @Description Retrieve a single company by its ID. No authentication required.
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path integer true "Company ID"
// @Success 200 {object} models.Company
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /company/{id} [get]
func GetCompanyByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.GetCompanyByID] Client IP: %s - Request to get company by id: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.GetCompanyByID] Client IP: %s - Invalid company ID %s: %v\n", ip, idStr, err)
		handleError(c, err)
		return
	}
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	company, err := service.GetCompanyByID(uint(id), userID)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetCompanyByID] Client IP: %s - Successfully retrieved company with ID %v\n", ip, id)
	c.JSON(http.StatusOK, company)
}

// AddCompany godoc
// @Summary Add a new company
// @Description Add a new company to the database. Requires authentication.
// @Tags Companies
// @Accept json
// @Produce json
// @Param company body models.Company true "Company data"
// @Success 201 {object} DefaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /company [post]
// @Security ApiKeyAuth
func AddCompany(c *gin.Context) {
	ip := c.ClientIP()
	var company models.Company
	logger.Info.Printf("[controllers.AddCompany] Client IP: %s - Request to add company\n", ip)
	if err := c.ShouldBindJSON(&company); err != nil {
		logger.Error.Printf("[controllers.AddCompany] Client IP: %s - Error parsing company data %v: %v\n", ip, company, err)
		handleError(c, err)
		return
	}
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

	if err := service.AddCompany(userID, company, role); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.AddCompany] Client IP: %s - Successfully added company %v\n", ip, company)
	c.JSON(http.StatusCreated, NewDefaultResponse("Company added successfully"))
}

// UpdateCompany godoc
// @Summary Update an existing company
// @Description Update a company by its ID. Requires authentication.
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path integer true "Company ID"
// @Param company body models.Company true "Updated company data"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /company/{id} [put]
// @Security ApiKeyAuth
func UpdateCompany(c *gin.Context) {
	ip := c.ClientIP()
	var company models.Company
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateCompany] Client IP: %s - Invalid company ID %s: %v\n", ip, idStr, err)
		handleError(c, err)
		return
	}
	company.ID = uint(id)
	if err := c.ShouldBindJSON(&company); err != nil {
		logger.Error.Printf("[controllers.UpdateCompany] Client IP: %s - Error parsing updated company data %v: %v\n", ip, company, err)
		handleError(c, err)
		return
	}
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

	if err := service.UpdateCompany(userID, company, role); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateCompany] Client IP: %s - Successfully updated company with ID %v\n", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Company updated successfully"))
}

// DeleteCompany godoc
// @Summary Delete a company
// @Description Soft delete a company by its ID. Requires authentication.
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path integer true "Company ID"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /company/{id} [delete]
// @Security ApiKeyAuth
func DeleteCompany(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteCompany] Client IP: %s - Invalid company ID %s: %v\n", ip, idStr, err)
		handleError(c, err)
		return
	}

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
	if err := service.DeleteCompany(uint(id), userID, role); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.DeleteCompany] Client IP: %s - Successfully soft deleted company with ID %v\n", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Company deleted successfully"))
}
