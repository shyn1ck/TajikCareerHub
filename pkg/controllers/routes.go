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
		userGroup.GET("/:id", GetUserByID)
		userGroup.DELETE("/:id", DeleteUser)
		userGroup.PATCH("/password", UpdateUserPassword)
	}

	vacancyGroup := r.Group("/vacancy").Use(checkUserAuthentication)
	{
		vacancyGroup.GET("/", GetAllVacancies)
		vacancyGroup.GET("/:id", GetVacancyByID)
		vacancyGroup.POST("/", AddVacancy)
		vacancyGroup.PUT("/:id", UpdateVacancy)
		vacancyGroup.DELETE("/:id", DeleteVacancy)
		vacancyGroup.GET("/report", GetVacancyReport)
	}

	resumeGroup := r.Group("/resumes").Use(checkUserAuthentication)
	{
		resumeGroup.GET("/", GetAllResumes)
		resumeGroup.GET("/:id", GetResumeByID)
		resumeGroup.POST("/", AddResume)
		resumeGroup.PUT("/:id", UpdateResume)
		resumeGroup.DELETE("/:id", DeleteResume)
		resumeGroup.PUT("/block/:id", BlockResume)
		resumeGroup.PUT("/unblock/:id", UnblockResume)
	}

	companyGroup := r.Group("/company").Use(checkUserAuthentication)
	{
		companyGroup.GET("/", GetAllCompanies)
		companyGroup.GET("/:id", GetCompanyByID)
		companyGroup.POST("/", AddCompany)
		companyGroup.PUT("/:id", UpdateCompany)
		companyGroup.DELETE("/:id", DeleteCompany)
	}

	applicationGroup := r.Group("/application").Use(checkUserAuthentication)
	{
		applicationGroup.GET("/", GetAllApplications)
		applicationGroup.GET("/:id", GetApplicationByID)
		applicationGroup.POST("/", AddApplication)
		applicationGroup.PUT("/:id", UpdateApplication)
		applicationGroup.DELETE("/:id", DeleteApplication)
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
