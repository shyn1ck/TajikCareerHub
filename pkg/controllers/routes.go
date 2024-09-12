package controllers

import (
	"TajikCareerHub/configs"
	_ "TajikCareerHub/docs"
	"TajikCareerHub/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

	vacancyGroup := r.Group("/vacancy").Use(checkUserAuthentication)
	{
		vacancyGroup.GET("/", GetAllVacancies)
		vacancyGroup.GET("/:id", GetVacancyByID)
		vacancyGroup.POST("/", AddVacancy)
		vacancyGroup.PUT("/:id", UpdateVacancy)
		vacancyGroup.DELETE("/:id", DeleteVacancy)
	}

	applicationGroup := r.Group("/applications").Use(checkUserAuthentication)
	{
		applicationGroup.GET("/", GetAllApplications)
		applicationGroup.GET("/:id", GetApplicationByID)
		applicationGroup.POST("/users/:user_id/jobs/:job_id/apply/:resume_id", ApplyForVacancy)
		applicationGroup.PUT("/:id", UpdateApplication)
		applicationGroup.DELETE("/:id", DeleteApplication)
		applicationGroup.GET("/user/:userID", GetApplicationsByUserID)
		applicationGroup.GET("/job/:jobID", GetApplicationsByVacancyID)
		applicationGroup.GET("/user/:userID/activity", GetUserApplicationActivity)
		applicationGroup.GET("/job/:jobID/applications", GetJobApplications)
		applicationGroup.PUT("/:id/status", UpdateApplicationStatus)
		applicationGroup.GET("/job/:jobID/report", GetJobReport)
	}

	companyGroup := r.Group("/company").Use(checkUserAuthentication)
	{
		companyGroup.GET("/", GetAllCompanies)
		companyGroup.GET("/:id", GetCompanyByID)
		companyGroup.POST("/", AddCompany)
		companyGroup.PUT("/:id", UpdateCompany)
		companyGroup.DELETE("/:id", DeleteCompany)
	}

	adminGroup := r.Group("/admin").Use(checkUserAuthentication)
	{
		adminGroup.PUT("/user/:id/block", BlockUser)
		adminGroup.PUT("/user/:id/unblock", UnblockUser)
	}

	VacancyCategoryGroup := r.Group("/category").Use(checkUserAuthentication)
	{
		VacancyCategoryGroup.GET("/", GetAllCategories)
		VacancyCategoryGroup.GET("/:id", GetCategoryByID)
		VacancyCategoryGroup.POST("/", CreateCategory)
		VacancyCategoryGroup.PUT("/:id", UpdateCategory)
		VacancyCategoryGroup.DELETE("/:id", DeleteCategory)
	}

	resumeGroup := r.Group("/resumes").Use(checkUserAuthentication)
	{
		resumeGroup.GET("/", GetAllResumes)
		resumeGroup.GET("/:id", GetResumeByID)
		resumeGroup.POST("/", AddResume)
		resumeGroup.PUT("/:id", UpdateResume)
		resumeGroup.DELETE("/:id", DeleteResume)
	}

	if err := r.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun)); err != nil {
		logger.Error.Fatalf("Error starting server: %v", err)
	}

	return r
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
