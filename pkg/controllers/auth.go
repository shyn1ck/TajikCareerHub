package controllers

import (
	_ "TajikCareerHub/docs"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUp
// @Summary Register a new user
// @Tags auth
// @Description Create a new user account with the provided details
// @ID create-account
// @Accept json
// @Produce json
// @Param input body models.User true "User registration information"
// @Success 201 {object} defaultResponse
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err)
		return
	}
	id, err := service.CreateUser(user)
	if err != nil {
		handleError(c, err)
		return
	}

	response := newDefaultResponse("user created successfully")
	c.JSON(http.StatusCreated, gin.H{
		"message": response.Message,
		"user_id": id,
	})
}

// SignIn
// @Summary Sign in to an existing account
// @Tags auth
// @Description Authenticate a user and return an access token
// @ID sign-in-to-account
// @Accept json
// @Produce json
// @Param input body models.User true "User sign-in information"
// @Success 200 {object} accessTokenResponse
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err)
		return
	}
	accessToken, err := service.SignIn(user.UserName, user.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, accessTokenResponse{accessToken})
}
