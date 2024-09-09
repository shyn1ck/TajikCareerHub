package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCategoryByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	category, err := service.GetCategoryByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetCategoryByID] Client IP: %s - Successfully retrieved vacancy category with ID %v\n", ip, id)
	c.JSON(http.StatusOK, category)
}

func GetAllCategories(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetAllVacancyCategories] Client IP: %s - Client requested all vacancy categories\n", ip)

	categories, err := service.GetAllCategories()
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetAllCategories] Client IP: %s - Successfully retrieved all vacancy categories\n", ip)
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	ip := c.ClientIP()
	var category models.VacancyCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		handleError(c, err)
		return
	}

	if err := service.AddCategory(category); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.CreateVacancyCategory] Client IP: %s - Successfully created vacancy category with data %v\n", ip, category)
	c.JSON(http.StatusCreated, gin.H{"message": "Vacancy category created successfully"})
}

func UpdateCategory(c *gin.Context) {
	ip := c.ClientIP()
	var category models.VacancyCategory
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	category.ID = uint(id)

	if err := c.ShouldBindJSON(&category); err != nil {
		handleError(c, err)
		return
	}

	if err := service.UpdateCategory(category); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateVacancyCategory] Client IP: %s - Successfully updated vacancy category with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Vacancy category updated successfully"})
}

func DeleteCategory(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	if err := service.DeleteCategory(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.DeleteVacancyCategory] Client IP: %s - Successfully soft deleted vacancy category with ID %v\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Vacancy category deleted successfully"})
}
