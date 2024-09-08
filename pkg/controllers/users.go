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
	logger.Info.Printf("[controllers.GetAllUsers] Client IP: %s - Request to get all users.\n", ip)
	users, err := service.GetAllUsers()
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetAllUsers] Client IP: %s - Successfully retrieved all users.\n", ip)
	c.JSON(http.StatusOK, users)
}

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

func GetUserByUsername(c *gin.Context) {
	ip := c.ClientIP()
	username := c.Query("username")
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

func UpdateUser(c *gin.Context) {
	ip := c.ClientIP()
	var user models.User
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
	user.ID = uint(userID)
	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, err)
		return
	}

	if err := service.UpdateUser(user); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateUser] Client IP: %s - User with ID %v updated successfully.\n", ip, user.ID)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

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
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func UpdateUserPassword(c *gin.Context) {
	ip := c.ClientIP()
	idStr := c.Param("id")
	logger.Info.Printf("[controllers.UpdateUserPassword] Client IP: %s - Request to update password for user with ID: %s\n", ip, idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	var request struct {
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		handleError(c, err)
		return
	}
	if err := service.UpdateUserPassword(uint(id), request.NewPassword); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateUserPassword] Client IP: %s - Password for user with ID %v updated successfully.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func CheckUserExists(c *gin.Context) {
	ip := c.ClientIP()
	username := c.Query("username")
	email := c.Query("email")
	logger.Info.Printf("[controllers.CheckUserExists] Client IP: %s - Request to check user existence with username: %s and email: %s\n", ip, username, email)
	usernameExists, emailExists, err := service.CheckUserExists(username, email)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.CheckUserExists] Client IP: %s - User existence check complete for username %v and email %v. Username Exists: %v, Email Exists: %v\n", ip, username, email, usernameExists, emailExists)
	c.JSON(http.StatusOK, gin.H{"username_exists": usernameExists, "email_exists": emailExists})
}
