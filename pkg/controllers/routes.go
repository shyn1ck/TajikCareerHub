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
		userGroup.POST("/", SignUp)
		userGroup.PUT("/", UpdateUser)
		userGroup.GET("/:id", GetUserByID)
		userGroup.DELETE("/:id", DeleteUser)
		userGroup.PATCH("/password", UpdateUserPassword)
		userGroup.PATCH("/block/:id", BlockUser)
		userGroup.PATCH("/unblock/:id", UnblockUser)
	}

	vacancyGroup := r.Group("/vacancies").Use(checkUserAuthentication)
	{
		vacancyGroup.GET("/", GetAllVacancies)
		vacancyGroup.GET("/:vacancyID", GetVacancyByID)
		vacancyGroup.POST("/", AddVacancy)
		vacancyGroup.PUT("/:id", UpdateVacancy)
		vacancyGroup.DELETE("/:id", DeleteVacancy)
		vacancyGroup.DELETE("/block/:id", BlockVacancy)
		vacancyGroup.PATCH("/unblock/:id", UnblockVacancy)
	}

	resumeGroup := r.Group("/resumes").Use(checkUserAuthentication)
	{
		resumeGroup.GET("/", GetAllResumes)
		resumeGroup.GET("/:id", GetResumeByID)
		resumeGroup.POST("/", AddResume)
		resumeGroup.PUT("/:id", UpdateResume)
		resumeGroup.DELETE("/:id", DeleteResume)
		resumeGroup.PATCH("/block/:id", BlockResume)
		resumeGroup.PATCH("/unblock/:id", UnblockResume)
	}

	companyGroup := r.Group("/companies").Use(checkUserAuthentication)
	{
		companyGroup.GET("/", GetAllCompanies)
		companyGroup.GET("/:id", GetCompanyByID)
		companyGroup.POST("/", AddCompany)
		companyGroup.PUT("/:id", UpdateCompany)
		companyGroup.DELETE("/:id", DeleteCompany)
	}

	applicationGroup := r.Group("/applications").Use(checkUserAuthentication)
	{
		applicationGroup.GET("/", GetAllApplications)
		applicationGroup.GET("/:application_id", GetApplicationByID) // Измените :id на :application_id
		applicationGroup.POST("/", AddApplication)
		applicationGroup.PUT("/:application_id", UpdateApplication)    // Измените :id на :application_id
		applicationGroup.DELETE("/:application_id", DeleteApplication) // Измените :id на :application_id
	}

	statusGroup := r.Group("/applications/:application_id/status")
	{
		statusGroup.PUT("/:status_id", UpdateApplicationStatus)
	}

	activityGroup := r.Group("/activities").Use(checkUserAuthentication)
	{
		activityGroup.GET("/", GetSpecialistActivityReportByUser)
		activityGroup.GET("/vacancy/:id", GetVacancyReportByID)
		activityGroup.GET("/resume/:id", GetResumeReportByID)
	}

	VacancyCategoryGroup := r.Group("/categories").Use(checkUserAuthentication)
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
