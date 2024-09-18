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

	userGroup := r.Group("/user").Use(checkUserAuthentication)
	{
		userGroup.GET("/", GetAllUsers)
		userGroup.POST("/", CreateUser)
		userGroup.PUT("/:id", UpdateUser)
		userGroup.GET("/:id", GetUserByID)
		userGroup.DELETE("/:id", DeleteUser)
		userGroup.PATCH("/password", UpdateUserPassword)
		userGroup.PUT("/block/:id", BlockUser)
		userGroup.PATCH("/unblock/:id", UnblockUser)
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

	resumeGroup := r.Group("/resume").Use(checkUserAuthentication)
	{
		resumeGroup.GET("/", GetAllResumes)
		resumeGroup.GET("/:id", GetResumeByID)
		resumeGroup.POST("/", AddResume)
		resumeGroup.PUT("/:id", UpdateResume)
		resumeGroup.DELETE("/:id", DeleteResume)
		resumeGroup.PATCH("/block/:id", BlockResume)
		resumeGroup.PATCH("/unblock/:id", UnblockResume)
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

	ActivityGroup := r.Group("/activity").Use(checkUserAuthentication)
	{
		ActivityGroup.GET("/", GetSpecialistActivityReport)
		ActivityGroup.GET("/vacancy/:id", GetVacancyReport)
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
