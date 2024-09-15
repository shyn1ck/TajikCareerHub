package controllers

import (
	"TajikCareerHub/errs"
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
// @Param        username  path    string     false    "username"
// @Success      200  {array}   models.User  "Success"
// @Success      200  {object}  models.User  "Success"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /users [get]
func GetAllUsers(c *gin.Context) {
	ip := c.ClientIP()
	username := c.Param("username")
	logger.Info.Printf("[controllers.GetAllUsers] Client IP: %s - Request to get users:\n", ip)
	if username != "" {
		user, err := service.GetUserByUsername(username)
		if err != nil {
			handleError(c, err)
			return
		}
		logger.Info.Printf("[controllers.GetAllUsers] Client IP: %s - Successfully retrieved user: %s.\n", ip, username)
		c.JSON(http.StatusOK, user)
	} else {
		users, err := service.GetAllUsers()
		if err != nil {
			handleError(c, err)
			return
		}
		logger.Info.Printf("[controllers.GetAllUsers] Client IP: %s - Successfully retrieved all users.\n", ip)
		c.JSON(http.StatusOK, users)
	}
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
		handleError(c, errs.ErrIDIsNotCorrect)
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

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the provided details
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body     models.User  true  "User data"
// @Success      201  {object}  DefaultResponse "Success"
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
	c.JSON(http.StatusCreated, NewDefaultResponse("User created successfully."))
}

// UpdateUserPassword godoc
// @Summary Update user password
// @Description Update the password for the current user. Requires authentication.
// @Tags Users
// @Accept json
// @Produce json
// @Param passwordRequest body PasswordRequest true "Password request data"
// @Success 200 {object} DefaultResponse "user password update successfully"
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /users/password [patch]
func UpdateUserPassword(c *gin.Context) {
	ip := c.ClientIP()
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateUserPassword] Client IP: %s - Request to update user password: %v\n", ip, userID)
	username, err := service.GetUsernameFromToken(c)
	if err != nil {
		handleError(c, err)
		return
	}

	var passwordRequest PasswordRequest
	if err := c.ShouldBindJSON(&passwordRequest); err != nil {
		logger.Error.Printf("[controllers.UpdateUserPassword] Client IP: %s - Error parsing request body: %v\n", ip, err)
		handleError(c, err)
		return
	}

	err = service.UpdateUserPassword(userID, username, passwordRequest.OldPassword, passwordRequest.NewPassword)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateUserPassword] Client IP: %s - Successfully updated password for user ID: %d.\n", ip, userID)
	c.JSON(http.StatusOK, NewDefaultResponse("User password updated successfully."))
}

// UpdateUser godoc
// @Summary      Update user details
// @Description  Update an existing user with the provided details
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "User ID"
// @Param        user  body     models.User  true  "Updated user data"
// @Success      200  {object}  DefaultResponse  "Success"
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
		c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid ID"})
		return
	}
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateUser] Client IP: %s - Error parsing user ID: %v", ip, err)
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
		logger.Error.Printf("[controllers.UpdateUser] Client IP: %s - Error parsing request body: %v\n", ip, err)
		handleError(c, errs.ErrShouldBindJson)
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
	c.JSON(http.StatusOK, NewDefaultResponse("User updated successfully."))
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Soft delete a user by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path    int     true    "User ID"
// @Success      200  {object}  DefaultResponse  "User deleted successfully"
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
		logger.Error.Printf("[controllers.DeleteUser] Client IP: %s - Error parsing user ID: %v", ip, err)
		return
	}
	if err := service.DeleteUser(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.DeleteUser] Client IP: %s - User with ID %v soft deleted successfully.\n", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("User deleted successfully."))
}

// BlockUser godoc
// @Summary      Block user
// @Description  Block a user by their ID
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id  path    integer  true  "User ID"  example(1)
// @Success      200  {object}  DefaultResponse  "User blocked successfully"
// @Failure      400  {object}  ErrorResponse    "Invalid ID"
// @Failure      500  {object}  ErrorResponse    "Internal server error"
// @Security     ApiKeyAuth
// @Router       /users/block/{id} [put]
func BlockUser(c *gin.Context) {
	ip := c.ClientIP()
	idParam := c.Param("id")
	logger.Info.Printf("[controllers.BlockUserController] Client IP: %s - Request to block user with ID %s.\n", ip, idParam)
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id == 0 {
		logger.Error.Printf("[controllers.BlockUserController] Client IP: %s - Invalid user ID: %s.\n", ip, idParam)
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: idParam})
		return
	}
	err = service.BlockUser(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.BlockUserController] Client IP: %s - Successfully blocked user with ID %d.\n", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("User blocked successfully."))
}

// UnblockUser godoc
// @Summary      Unblock user
// @Description  Unlock a user by ID
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id  path    integer  true  "User ID"  example(1)
// @Success      200  {object}  DefaultResponse  "User unblocked successfully"
// @Failure      400  {object}  ErrorResponse    "Invalid ID"
// @Failure      500  {object}  ErrorResponse    "Internal server error"
// @Security     ApiKeyAuth
// @Router       /users/unblock/{id} [put]
func UnblockUser(c *gin.Context) {
	ip := c.ClientIP()
	idParam := c.Param("id")
	logger.Info.Printf("[controllers.UnblockUserController] Client IP: %s - Request to unblock user with ID %s.\n", ip, idParam)
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id == 0 {
		logger.Error.Printf("[controllers.UnblockUserController] Client IP: %s - Invalid user ID: %s.\n", ip, idParam)
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: idParam})
		return
	}
	err = service.UnblockUser(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UnblockUserController] Client IP: %s - Successfully unblocked user with ID %d.\n", ip, id)
	c.JSON(http.StatusOK, NewDefaultResponse("User unblocked successfully."))
}
