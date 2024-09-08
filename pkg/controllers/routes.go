package controllers

import (
	"TajikCareerHub/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	r.GET("/ping", PingPong)

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	userGroup := r.Group("/users").Use(checkUserAuthentication)
	{
		userGroup.GET("/", GetAllUsers)
		userGroup.POST("/", CreateUser)
		userGroup.PUT("/:id", UpdateUser)
		userGroup.DELETE("/:id", DeleteUser)
	}

	// Specific routes for username and id should be placed before general routes to avoid conflicts
	r.GET("/users/username/:username", GetUserByUsername)
	r.GET("/users/:id", GetUserByID)

	passwordGroup := r.Group("/users/:id/password").Use(checkUserAuthentication)
	{
		passwordGroup.PATCH("/", UpdateUserPassword)
	}

	existenceGroup := r.Group("/users/existence")
	{
		existenceGroup.GET("/", CheckUserExists)
	}

	jobGroup := r.Group("/jobs").Use(checkUserAuthentication)
	{
		jobGroup.GET("/", GetAllJobs)
		jobGroup.GET("/:id", GetJobByID)
		jobGroup.POST("/", adminOnly, AddJob)                      // Employer only
		jobGroup.PUT("/:id", employerOnly, UpdateJob)              // Employer only
		jobGroup.DELETE("/:id", employerOnly, DeleteJob)           // Employer only
		jobGroup.PUT("/:id/salary", employerOnly, UpdateJobSalary) // Employer only
	}

	applicationGroup := r.Group("/applications").Use(checkUserAuthentication)
	{
		applicationGroup.GET("/", adminOnly, GetAllApplications) // Admin only
		applicationGroup.GET("/:id", GetApplicationByID)
		applicationGroup.POST("/", specialistOnly, AddApplication)         // Specialist only
		applicationGroup.PUT("/:id", specialistOnly, UpdateApplication)    // Specialist only
		applicationGroup.DELETE("/:id", specialistOnly, DeleteApplication) // Specialist only
		applicationGroup.GET("/user/:userID", GetApplicationsByUserID)
		applicationGroup.GET("/job/:jobID", employerOnly, GetApplicationsByJobID) // Employer only
	}

	companyGroup := r.Group("/companies").Use(checkUserAuthentication)
	{
		companyGroup.GET("/", adminOnly, GetAllCompanies) // Admin only
		companyGroup.GET("/:id", GetCompanyByID)
		companyGroup.POST("/", employerOnly, AddCompany)      // Employer only
		companyGroup.PUT("/:id", employerOnly, UpdateCompany) // Employer only
		companyGroup.DELETE("/:id", adminOnly, DeleteCompany) // Admin only
	}

	favoriteGroup := r.Group("/favorites").Use(checkUserAuthentication)
	{
		favoriteGroup.GET("/user/:userID", GetFavoritesByUserID)
		favoriteGroup.GET("/user/:userID/job/:jobID", GetFavoriteByUserIDAndJobID)
		favoriteGroup.POST("/", specialistOnly, AddFavorite)      // Specialist only
		favoriteGroup.DELETE("/", specialistOnly, RemoveFavorite) // Specialist only
		favoriteGroup.GET("/exists/user/:userID/job/:jobID", CheckFavoriteExists)
	}

	jobCategoryGroup := r.Group("/job-categories").Use(checkUserAuthentication)
	{
		jobCategoryGroup.GET("/", GetAllJobCategories)
		jobCategoryGroup.GET("/:id", GetJobCategoryByID)
		jobCategoryGroup.POST("/", adminOnly, CreateJobCategory)      // Admin only
		jobCategoryGroup.PUT("/:id", adminOnly, UpdateJobCategory)    // Admin only
		jobCategoryGroup.DELETE("/:id", adminOnly, DeleteJobCategory) // Admin only
	}

	err := r.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun))
	if err != nil {
		return r
	}

	return nil
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func checkUserRole(c *gin.Context, role string) bool {
	// Retrieve user role from context or JWT token
	// This is a placeholder, replace with your actual role check logic
	userRole, exists := c.Get("role")
	return exists && userRole == role
}
