package controllers

import (
	"TajikCareerHub/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunRoutes() error {
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
		userGroup.GET("/", adminOnly, GetAllUsers)
		userGroup.GET("/:id", GetUserByID)
		userGroup.POST("/", adminOnly, CreateUser)
		userGroup.PUT("/:id", UpdateUser)
		userGroup.DELETE("/:id", adminOnly, DeleteUser)
	}

	passwordGroup := r.Group("/users/:id/password").Use(checkUserAuthentication)
	{
		passwordGroup.PUT("/", UpdateUserPassword)
	}

	existenceGroup := r.Group("/users/existence")
	{
		existenceGroup.GET("/", CheckUserExists)
	}

	jobGroup := r.Group("/jobs").Use(checkUserAuthentication)
	{
		jobGroup.GET("/", GetAllJobs)
		jobGroup.GET("/:id", GetJobByID)
		jobGroup.POST("/", employerOnly, AddJob)
		jobGroup.PUT("/:id", employerOnly, UpdateJob)
		jobGroup.DELETE("/:id", employerOnly, DeleteJob)
		jobGroup.GET("/salary-range", GetJobsBySalaryRange)
		jobGroup.PUT("/:id/salary", employerOnly, UpdateJobSalary)
	}

	applicationGroup := r.Group("/applications").Use(checkUserAuthentication)
	{
		applicationGroup.GET("/", adminOnly, GetAllApplications)
		applicationGroup.GET("/:id", GetApplicationByID)
		applicationGroup.POST("/", specialistOnly, AddApplication)
		applicationGroup.PUT("/:id", specialistOnly, UpdateApplication)
		applicationGroup.DELETE("/:id", specialistOnly, DeleteApplication)
		applicationGroup.GET("/user/:userID", GetApplicationsByUserID)
		applicationGroup.GET("/job/:jobID", employerOnly, GetApplicationsByJobID)
	}

	companyGroup := r.Group("/companies").Use(checkUserAuthentication)
	{
		companyGroup.GET("/", adminOnly, GetAllCompanies)
		companyGroup.GET("/:id", GetCompanyByID)
		companyGroup.POST("/", employerOnly, AddCompany)
		companyGroup.PUT("/:id", employerOnly, UpdateCompany)
		companyGroup.DELETE("/:id", adminOnly, DeleteCompany)
	}

	favoriteGroup := r.Group("/favorites").Use(checkUserAuthentication)
	{
		favoriteGroup.GET("/user/:userID", GetFavoritesByUserID)
		favoriteGroup.GET("/user/:userID/job/:jobID", GetFavoriteByUserIDAndJobID)
		favoriteGroup.POST("/", specialistOnly, AddFavorite)
		favoriteGroup.DELETE("/", specialistOnly, RemoveFavorite)
		favoriteGroup.GET("/exists/user/:userID/job/:jobID", CheckFavoriteExists)
	}

	jobCategoryGroup := r.Group("/job-categories").Use(checkUserAuthentication)
	{
		jobCategoryGroup.GET("/", GetAllJobCategories)
		jobCategoryGroup.GET("/:id", GetJobCategoryByID)
		jobCategoryGroup.POST("/", adminOnly, CreateJobCategory)
		jobCategoryGroup.PUT("/:id", adminOnly, UpdateJobCategory)
		jobCategoryGroup.DELETE("/:id", adminOnly, DeleteJobCategory)
	}

	err := r.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun))
	if err != nil {
		return err
	}

	return nil
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
