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

// GetCategoryByID godoc
// @Summary      Get category by ID
// @Description  Retrieve a specific category by its ID
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "Category ID"
// @Success      200  {object}  models.VacancyCategory   "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID"
// @Failure      404  {object}  ErrorResponse  "Category not found"
// @Failure      403  {object}  ErrorResponse  "Access Denied"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.GetCategory] Error converting id to int: %s", err.Error())
		handleError(c, errs.ErrIDIsNotCorrect)
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

// GetAllCategories godoc
// @Summary      Get all categories
// @Description  Retrieve a list of all categories
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.VacancyCategory  "Success"
// @Failure      403  {object}  ErrorResponse  "Access Denied"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /categories [get]
func GetAllCategories(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetAllCategories] Client IP: %s - Client requested all categories\n", ip)

	categories, err := service.GetAllCategories()
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetAllCategories] Client IP: %s - Successfully retrieved all categories\n", ip)
	c.JSON(http.StatusOK, categories)
}

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Create a new category with the provided details
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        category  body models.VacancyCategory  true  "Category data"
// @Success      201  {object}  DefaultResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid input"
// @Failure      403  {object}  ErrorResponse  "Access Denied"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /categories [post]
func CreateCategory(c *gin.Context) {
	ip := c.ClientIP()
	var category models.VacancyCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		handleError(c, errs.ErrShouldBindJson)
		return
	}
	role, err := service.GetRoleFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	if err := service.AddCategory(category, role); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.CreateCategory] Client IP: %s - Successfully created category with data %v\n", ip, category)
	c.JSON(http.StatusCreated, NewDefaultResponse("Category created successfully"))
}

// UpdateCategory godoc
// @Summary      Update category
// @Description  Update an existing category with the provided details
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "Category ID"
// @Param        category  body models.VacancyCategory  true  "Updated category data"
// @Success      200  {object}  DefaultResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID or input"
// @Failure      404  {object}  ErrorResponse  "Category not found"
// @Failure      403   {object} ErrorResponse  "Access Denied"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	ip := c.ClientIP()
	var category models.VacancyCategory
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, errs.ErrIDIsNotCorrect)
		return
	}
	category.ID = uint(id)

	if err := c.ShouldBindJSON(&category); err != nil {
		handleError(c, errs.ErrShouldBindJson)
		return
	}
	role, err := service.GetRoleFromToken(c)
	if err != nil {
		handleError(c, err)
	}

	if err := service.UpdateCategory(category, role); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateCategory] Client IP: %s - Successfully updated category with ID %v\n", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Category updated successfully"))
}

// DeleteCategory godoc
// @Summary      Delete category
// @Description  Soft delete a category by ID
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "Category ID"
// @Success      200  {object}  DefaultResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID"
// @Failure      403  {object}  ErrorResponse  "Access Denied"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, errs.ErrIDIsNotCorrect)
		return
	}
	role, err := service.GetRoleFromToken(c)
	if err != nil {
		handleError(c, err)
	}

	if err := service.DeleteCategory(uint(id), role); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.DeleteCategory] Client IP: %s - Successfully soft deleted category with ID %v\n", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("Category deleted successfully"))
}
