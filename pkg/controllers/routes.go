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
		jobGroup.POST("/", adminOnly, AddJob)
		jobGroup.PUT("/:id", UpdateJob)
		jobGroup.DELETE("/:id", DeleteJob)
	}

	applicationGroup := r.Group("/applications").Use(checkUserAuthentication)
	{
		applicationGroup.GET("/", GetAllApplications)
		applicationGroup.GET("/:id", GetApplicationByID)
		applicationGroup.POST("/", AddApplication)
		applicationGroup.PUT("/:id", UpdateApplication)
		applicationGroup.DELETE("/:id", DeleteApplication)
		applicationGroup.GET("/user/:userID", GetApplicationsByUserID)
		applicationGroup.GET("/job/:jobID", GetApplicationsByJobID)
	}

	companyGroup := r.Group("/companies").Use(checkUserAuthentication)
	{
		companyGroup.GET("/", GetAllCompanies)
		companyGroup.GET("/:id", GetCompanyByID)
		companyGroup.POST("/", AddCompany)
		companyGroup.PUT("/:id", UpdateCompany)
		companyGroup.DELETE("/:id", DeleteCompany)
	}

	favoriteGroup := r.Group("/favorites").Use(checkUserAuthentication)
	{
		favoriteGroup.GET("/user/:userID", GetFavoritesByUserID)
		favoriteGroup.GET("/user/:userID/job/:jobID", GetFavoriteByUserIDAndJobID)
		favoriteGroup.POST("/", AddFavorite)
		favoriteGroup.DELETE("/", RemoveFavorite)
		favoriteGroup.GET("/exists/user/:userID/job/:jobID", CheckFavoriteExists)
	}

	jobCategoryGroup := r.Group("/job-categories").Use(checkUserAuthentication)
	{
		jobCategoryGroup.GET("/", GetAllJobCategories)
		jobCategoryGroup.GET("/:id", GetJobCategoryByID)
		jobCategoryGroup.POST("/", CreateJobCategory)
		jobCategoryGroup.PUT("/:id", UpdateJobCategory)
		jobCategoryGroup.DELETE("/:id", DeleteJobCategory)
	}

	resumeGroup := r.Group("/resumes").Use(checkUserAuthentication)
	{
		resumeGroup.GET("/", GetAllResumes)
		resumeGroup.GET("/:id", GetResumeByID)
		resumeGroup.POST("/", AddResume)
		resumeGroup.PUT("/:id", UpdateResume)
		resumeGroup.DELETE("/:id", DeleteResume)
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
