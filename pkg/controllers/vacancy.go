package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllVacancies(c *gin.Context) {
	ip := c.ClientIP()
	search := c.Query("search")
	location := c.Query("location")
	category := c.Query("category")
	minSalaryStr := c.Query("min-salary")
	maxSalaryStr := c.Query("max-salary")
	sort := c.Query("sort")

	logger.Info.Printf("[controllers.GetAllVacancies] Client IP: %s - Request to get vacancies with keyword: %s, minSalary: %s, maxSalary: %s, location: %s, category: %s, sort: %s\n", ip, search, minSalaryStr, maxSalaryStr, location, category, sort)

	var minSalary, maxSalary int
	var err error
	if minSalaryStr != "" {
		minSalary, err = strconv.Atoi(minSalaryStr)
		if err != nil {
			handleError(c, err)
			return
		}
	}

	if maxSalaryStr != "" {
		maxSalary, err = strconv.Atoi(maxSalaryStr)
		if err != nil {
			handleError(c, err)
			return
		}
	}

	vacancies, err := service.GetAllVacancies(search, minSalary, maxSalary, location, category, sort)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetAllVacancies] Client IP: %s - Successfully retrieved vacancies with keyword: %s, minSalary: %d, maxSalary: %d, location: %s, category: %s, sort: %s\n", ip, search, minSalary, maxSalary, location, category, sort)
	c.JSON(http.StatusOK, gin.H{"vacancies": vacancies})
}

func GetVacancyByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.GetVacancyByID] Client IP: %s - Request to get vacancy by ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	vacancy, err := service.GetVacancyByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetVacancyByID] Client IP: %s - Successfully retrieved vacancy with ID %v.\n", ip, id)
	c.JSON(http.StatusOK, vacancy)
}

func AddVacancy(c *gin.Context) {
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	var vacancy models.Vacancy
	if err := c.BindJSON(&vacancy); err != nil {
		handleError(c, err)
		return
	}

	vacancy.UserID = userID

	err = service.AddVacancy(vacancy)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vacancy added successfully"})
}

func UpdateVacancy(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.UpdateVacancy] Client IP: %s - Request to update vacancy with ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	var updatedVacancy models.Vacancy
	if err := c.BindJSON(&updatedVacancy); err != nil {
		handleError(c, err)
		return
	}

	err = service.UpdateVacancy(uint(id), updatedVacancy)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateVacancy] Client IP: %s - Vacancy with ID %v updated successfully.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Vacancy updated successfully"})
}

func DeleteVacancy(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.DeleteVacancy] Client IP: %s - Request to delete vacancy with ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	err = service.DeleteVacancy(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.DeleteVacancy] Client IP: %s - Vacancy with ID %v deleted successfully.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Vacancy deleted successfully"})
}
