package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunRoutes() error {
	r := gin.Default()
	r.GET("/ping", PingPong)
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", GetAllUsers)
		userGroup.GET("/:id", GetUserByID)
		userGroup.POST("/", CreateUser)
		userGroup.PUT("/:id", UpdateUser)
		userGroup.DELETE("/:id", DeleteUser)
	}
	passwordGroup := r.Group("/users/:id/password")
	{
		passwordGroup.PUT("/", UpdateUserPassword)
	}
	existenceGroup := r.Group("/users/existence")
	{
		existenceGroup.GET("/", CheckUserExists)
	}
	jobGroup := r.Group("/jobs")
	{
		jobGroup.GET("/", GetAllJobs)
		jobGroup.GET("/:id", GetJobByID)
		jobGroup.POST("/", AddJob)
		jobGroup.PUT("/:id", UpdateJob)
		jobGroup.DELETE("/:id", DeleteJob)
		jobGroup.GET("/filter", FilterJobs)
		jobGroup.GET("/salary-range", GetJobsBySalaryRange)
		jobGroup.PUT("/:id/salary", UpdateJobSalary)
	}

	port := ":8181"
	err := r.Run(port)
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
