package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve a list of all users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.User  "Success"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /users [get]
func GetAllUsers(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetAllUsers] Client IP: %s - Request to get all users.\n", ip)
	users, err := service.GetAllUsers()
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetAllUsers] Client IP: %s - Successfully retrieved all users.\n", ip)
	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Retrieve a specific user by its ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "User ID"
// @Success      200  {object}  models.User  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID"
// @Failure      404  {object}  ErrorResponse  "User not found"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /users/{id} [get]
func GetUserByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.GetUserByID] Client IP: %s - Request to get user by ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	user, err := service.GetUserByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetUserByID] Client IP: %s - Successfully retrieved user with ID %v.\n", ip, id)
	c.JSON(http.StatusOK, user)
}

// GetUserByUsername godoc
// @Summary      Get user by username
// @Description  Retrieve a specific user by their username
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        username  path    string  true  "Username"
// @Success      200  {object}  models.User  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid username"
// @Failure      404  {object}  ErrorResponse  "User not found"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /users/username/{username} [get]
func GetUserByUsername(c *gin.Context) {
	ip := c.ClientIP()
	username := c.Param("username")
	logger.Info.Printf("[controllers.GetUserByUsername] Client IP: %s - Request to get user by username: %s\n", ip, username)

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	user, err := service.GetUserByUsername(username)
	if err != nil {
		handleError(c, err)
		return
	}

	if user == nil {
		logger.Info.Printf("[controllers.GetUserByUsername] Client IP: %s - No user found with username: %s\n", ip, username)
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	logger.Info.Printf("[controllers.GetUserByUsername] Client IP: %s - Successfully retrieved user with username: %s\n", ip, username)
	c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the provided details
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body    models.User  true  "User data"
// @Success      201  {object}  defaultResponse "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid input"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	ip := c.ClientIP()
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.CreateUser] Client IP: %s - Request to create user: %v\n", ip, user)
	if _, err := service.CreateUser(user); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.CreateUser] Client IP: %s - User %v created successfully.\n", ip, user.UserName)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// UpdateUser godoc
// @Summary      Update user details
// @Description  Update an existing user with the provided details
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "User ID"
// @Param        user  body    models.User  true  "Updated user data"
// @Success      200  {object}  defaultResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID or input"
// @Failure      404  {object}  ErrorResponse  "User not found"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /users/{id} [put]
func UpdateUser(c *gin.Context) {
	ip := c.ClientIP()
	id := c.Param("id")
	logger.Info.Printf("[controllers.UpdateUser] Client IP: %s - Request to update user with ID: %s\n", ip, id)

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserIDIsRequired"})
		return
	}
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	var userInput struct {
		FullName  *string `json:"full_name"`
		Username  *string `json:"username"`
		BirthDate *string `json:"birth_date"`
		Email     *string `json:"email"`
		Password  *string `json:"password"`
		Role      *string `json:"role"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		handleError(c, err)
		return
	}

	user := models.User{
		ID: uint(userID),
	}

	if userInput.FullName != nil {
		user.FullName = *userInput.FullName
	}
	if userInput.Username != nil {
		user.UserName = *userInput.Username
	}
	if userInput.BirthDate != nil {
		parsedDate, err := time.Parse(time.RFC3339, *userInput.BirthDate)
		if err != nil {
			handleError(c, err)
			return
		}
		user.BirthDate = parsedDate
	}
	if userInput.Email != nil {
		user.Email = *userInput.Email
	}

	if err := service.UpdateUser(user); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateUser] Client IP: %s - User with ID %v updated successfully.\n", ip, user.ID)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Soft delete a user by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "User ID"
// @Success      200  {object}  defaultResponse  "Success"
// @Failure      400  {object}  ErrorResponse  "Invalid ID"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.DeleteUser] Client IP: %s - Request to delete user with ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	if err := service.DeleteUser(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.DeleteUser] Client IP: %s - User with ID %v soft deleted successfully.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "User soft deleted successfully"})
}
