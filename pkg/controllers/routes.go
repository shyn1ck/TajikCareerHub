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

	// Изменение маршрута для имени пользователя
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

	// Роуты для резюме
	resumeGroup := r.Group("/resumes").Use(checkUserAuthentication)
	{
		resumeGroup.GET("/", GetAllResumes)
		resumeGroup.GET("/:id", GetResumeByID)
		resumeGroup.POST("/", AddResume)                         // Specialist only
		resumeGroup.PUT("/:id", specialistOnly, UpdateResume)    // Specialist only
		resumeGroup.DELETE("/:id", specialistOnly, DeleteResume) // Specialist only
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
