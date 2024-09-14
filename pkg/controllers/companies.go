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
// @Success 200 {array} defaultResponse
// @Failure 500 {object} ErrorResponse
// @Router /companies [get]
func GetAllCompanies(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetAllCompanies] Client IP: %s - Client requested all companies\n", ip)
	companies, err := service.GetAllCompanies()
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
// @Success 200 {object} defaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /companies/{id} [get]
func GetCompanyByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.GetCompanyByID] Client IP: %s - Client requested company with ID %s. Error: Invalid company ID\n", ip, idStr)
		handleError(c, err)
		return
	}
	company, err := service.GetCompanyByID(uint(id))
	if err != nil {
		logger.Info.Printf("[controllers.GetCompanyByID] Client IP: %s - Client requested company with ID %v. Error retrieving company\n", ip, id)
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
// @Success 201 {object} defaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /companies [post]
// @Security ApiKeyAuth
func AddCompany(c *gin.Context) {
	ip := c.ClientIP()
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		logger.Info.Printf("[controllers.AddCompany] Client IP: %s - Client attempted to add company with data %v. Error: Invalid input\n", ip, company)
		handleError(c, err)
		return
	}
	if err := service.AddCompany(company); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.AddCompany] Client IP: %s - Successfully added company with data %v\n", ip, company)
	c.JSON(http.StatusCreated, gin.H{"message": "Company added successfully"})
}

// UpdateCompany godoc
// @Summary Update an existing company
// @Description Update a company by its ID. Requires authentication.
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path integer true "Company ID"
// @Param company body models.Company true "Updated company data"
// @Success 200 {object} defaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /companies/{id} [put]
// @Security ApiKeyAuth
func UpdateCompany(c *gin.Context) {
	ip := c.ClientIP()
	var company models.Company
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.UpdateCompany] Client IP: %s - Client attempted to update company with ID %s. Error: Invalid company ID\n", ip, idStr)
		handleError(c, err)
		return
	}
	company.ID = uint(id)
	if err := c.ShouldBindJSON(&company); err != nil {
		logger.Info.Printf("[controllers.UpdateCompany] Client IP: %s - Client attempted to update company with ID %v using data %v. Error: Invalid input\n", ip, id, company)
		handleError(c, err)
		return
	}
	if err := service.UpdateCompany(company); err != nil {
		logger.Info.Printf("[controllers.UpdateCompany] Client IP: %s - Client attempted to update company with ID %v using data %v. Error updating company\n", ip, id, company)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateCompany] Client IP: %s - Successfully updated company with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully"})
}

// DeleteCompany godoc
// @Summary Delete a company
// @Description Soft delete a company by its ID. Requires authentication.
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path integer true "Company ID"
// @Success 200 {object} defaultResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /companies/{id} [delete]
// @Security ApiKeyAuth
func DeleteCompany(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.DeleteCompany] Client IP: %s - Client attempted to delete company with ID %s. Error: Invalid company ID\n", ip, idStr)
		handleError(c, err)
		return
	}

	if err := service.DeleteCompany(uint(id)); err != nil {
		logger.Info.Printf("[controllers.DeleteCompany] Client IP: %s - Client attempted to delete company with ID %v. Error soft deleting company\n", ip, id)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.DeleteCompany] Client IP: %s - Successfully soft deleted company with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
