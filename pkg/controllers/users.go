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

func BlockUser(c *gin.Context) {
	ip := c.ClientIP()
	idParam := c.Param("id")
	logger.Info.Printf("[controllers.BlockUserController] Client IP: %s - Request to block user with ID %s.\n", ip, idParam)
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id == 0 {
		logger.Error.Printf("[controllers.BlockUserController] Client IP: %s - Invalid user ID: %s.\n", ip, idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = service.BlockUser(uint(id))
	if err != nil {
		logger.Error.Printf("[controllers.BlockUserController] Client IP: %s - Failed to block user with ID %d: %v.\n", ip, id, err)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.BlockUserController] Client IP: %s - Successfully blocked user with ID %d.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "User blocked successfully"})
}

func UnblockUser(c *gin.Context) {
	ip := c.ClientIP()
	idParam := c.Param("id")
	logger.Info.Printf("[controllers.UnblockUserController] Client IP: %s - Request to unblock user with ID %s.\n", ip, idParam)
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id == 0 {
		logger.Error.Printf("[controllers.UnblockUserController] Client IP: %s - Invalid user ID: %s.\n", ip, idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = service.UnblockUser(uint(id))
	if err != nil {
		logger.Error.Printf("[controllers.UnblockUserController] Client IP: %s - Failed to unblock user with ID %d: %v.\n", ip, id, err)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UnblockUserController] Client IP: %s - Successfully unblocked user with ID %d.\n", ip, id)
	c.JSON(http.StatusOK, gin.H{"message": "User unblocked successfully"})
}
