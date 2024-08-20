package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	ip := c.ClientIP()
	users, err := service.GetAllUsers()
	if err != nil {
		logger.Error.Printf("[controllers.GetAllUsers] Client IP: %s - Error retrieving all users: %v\n", ip, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	logger.Info.Printf("[controllers.GetAllUsers] Client IP: %s - Successfully retrieved all users.\n", ip)
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] Client IP: %s - Invalid user ID: %v\n", ip, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := service.GetUserByID(uint(id))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] Client IP: %s - Error retrieving user with ID %v: %v\n", ip, id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	logger.Info.Printf("[controllers.GetUserByID] Client IP: %s - Successfully retrieved user with ID %v.\n", ip, id)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	ip := c.ClientIP()
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Error.Printf("[controllers.CreateUser] Client IP: %s - Invalid input: %v\n", ip, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := service.CreateUser(user); err != nil {
		logger.Error.Printf("[controllers.CreateUser] Client IP: %s - Error creating user %v: %v\n", ip, user.UserName, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("[controllers.CreateUser] Client IP: %s - User %v created successfully.\n", ip, user.UserName)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func UpdateUser(c *gin.Context) {
	ip := c.ClientIP()
	var user models.User
	id := c.Param("id")
	if id == "" {
		logger.Error.Printf("[controllers.UpdateUser] Client IP: %s - Missing user ID\n", ip)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateUser] Client IP: %s - Invalid user ID format: %v\n", ip, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	user.ID = uint(userID)
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Error.Printf("[controllers.UpdateUser] Client IP: %s - Invalid input: %v\n", ip, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := service.UpdateUser(user); err != nil {
		logger.Error.Printf("[controllers.UpdateUser] Client IP: %s - Error updating user with ID %v: %v\n", ip, user.ID, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Info.Printf("[controllers.UpdateUser] Client IP: %s - User with ID %v updated successfully.\n", ip, user.ID)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteUser] Client IP: %s - Invalid user ID: %v\n", ip, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := service.DeleteUser(uint(id)); err != nil {
		logger.Error.Printf("[controllers.DeleteUser] Client IP: %s - Error soft deleting user with ID %v: %v\n", ip, id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	logger.Info.Printf("[controllers.DeleteUser] Client IP: %s - User with ID %v soft deleted successfully.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func UpdateUserPassword(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateUserPassword] Client IP: %s - Invalid user ID: %v\n", ip, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var request struct {
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error.Printf("[controllers.UpdateUserPassword] Client IP: %s - Invalid input: %v\n", ip, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := service.UpdateUserPassword(uint(id), request.NewPassword); err != nil {
		logger.Error.Printf("[controllers.UpdateUserPassword] Client IP: %s - Error updating password for user with ID %v: %v\n", ip, id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("[controllers.UpdateUserPassword] Client IP: %s - Password for user with ID %v updated successfully.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func CheckUserExists(c *gin.Context) {
	ip := c.ClientIP()
	username := c.Query("username")
	email := c.Query("email")
	exists, err := service.CheckUserExists(username, email)
	if err != nil {
		logger.Error.Printf("[controllers.CheckUserExists] Client IP: %s - Error checking user existence with username %v or email %v: %v\n", ip, username, email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user existence"})
		return
	}
	logger.Info.Printf("[controllers.CheckUserExists] Client IP: %s - User existence check complete for username %v and email %v. Exists: %v\n", ip, username, email, exists)
	c.JSON(http.StatusOK, gin.H{"exists": exists})
}
