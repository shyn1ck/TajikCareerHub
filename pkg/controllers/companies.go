package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllCompanies(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetAllCompanies] Client IP: %s - Client requested all companies\n", ip)
	companies, err := service.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve companies"})
		return
	}
	logger.Info.Printf("[controllers.GetAllCompanies] Client IP: %s - Successfully retrieved all companies\n", ip)
	c.JSON(http.StatusOK, companies)
}

func GetCompanyByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.GetCompanyByID] Client IP: %s - Client requested company with ID %s. Error: Invalid company ID\n", ip, idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}
	company, err := service.GetCompanyByID(uint(id))
	if err != nil {
		logger.Info.Printf("[controllers.GetCompanyByID] Client IP: %s - Client requested company with ID %v. Error retrieving company\n", ip, id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	logger.Info.Printf("[controllers.GetCompanyByID] Client IP: %s - Successfully retrieved company with ID %v\n", ip, id)
	c.JSON(http.StatusOK, company)
}

func AddCompany(c *gin.Context) {
	ip := c.ClientIP()
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		logger.Info.Printf("[controllers.AddCompany] Client IP: %s - Client attempted to add company with data %v. Error: Invalid input\n", ip, company)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := service.AddCompany(company); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to add company"})
		return
	}
	logger.Info.Printf("[controllers.AddCompany] Client IP: %s - Successfully added company with data %v\n", ip, company)
	c.JSON(http.StatusCreated, gin.H{"message": "Company added successfully"})
}

func UpdateCompany(c *gin.Context) {
	ip := c.ClientIP()
	var company models.Company
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.UpdateCompany] Client IP: %s - Client attempted to update company with ID %s. Error: Invalid company ID\n", ip, idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}
	company.ID = uint(id)
	if err := c.ShouldBindJSON(&company); err != nil {
		logger.Info.Printf("[controllers.UpdateCompany] Client IP: %s - Client attempted to update company with ID %v using data %v. Error: Invalid input\n", ip, id, company)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := service.UpdateCompany(company); err != nil {
		logger.Info.Printf("[controllers.UpdateCompany] Client IP: %s - Client attempted to update company with ID %v using data %v. Error updating company\n", ip, id, company)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
		return
	}
	logger.Info.Printf("[controllers.UpdateCompany] Client IP: %s - Successfully updated company with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully"})
}

func DeleteCompany(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Info.Printf("[controllers.DeleteCompany] Client IP: %s - Client attempted to delete company with ID %s. Error: Invalid company ID\n", ip, idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	if err := service.DeleteCompany(uint(id)); err != nil {
		logger.Info.Printf("[controllers.DeleteCompany] Client IP: %s - Client attempted to delete company with ID %v. Error soft deleting company\n", ip, id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}
	logger.Info.Printf("[controllers.DeleteCompany] Client IP: %s - Successfully soft deleted company with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
