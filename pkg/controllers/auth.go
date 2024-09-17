package controllers

import (
	_ "TajikCareerHub/docs"
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUp
// @Summary Register a new user
// @Tags Authorization
// @Description Create a new user account with the provided details
// @ID create-account
// @Accept json
// @Produce json
// @Param input body models.User true "User registration information"
// @Success 201 {object} DefaultResponse "User created successfully"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("Client with IP: %s requested to create a new user", ip)
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("Client with IP: %s failed to create new user: Error parsing request body: %v", ip, err)
		handleError(c, err)
		return
	}

	id, err := service.CreateUser(user)
	if err != nil {
		handleError(c, err)
		return
	}
	response := NewDefaultResponse("User created successfully")
	logger.Info.Printf("Client with IP: %s successfully created a new user with ID: %d", ip, id)
	c.JSON(http.StatusCreated, gin.H{
		"message": response.Message,
		"user_id": id,
	})
}

// SignIn
// @Summary Sign in to an existing account
// @Tags Authorization
// @Description Authenticate a user and return an access token
// @ID sign-in-to-account
// @Accept json
// @Produce json
// @Param input body models.User true "User sign-in information"
// @Success 200 {object} AccessTokenResponse
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("Client with IP: %s requested to sign in", ip)

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("Client with IP: %s failed to sign in: Error parsing request body: %v", ip, err)
		handleError(c, err)
		return
	}
	accessToken, err := service.SignIn(user.UserName, user.Password)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("Client with IP: %s successfully signed in", ip)
	c.JSON(http.StatusOK, AccessTokenResponse{accessToken})
}

// Sign in account
