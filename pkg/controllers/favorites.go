package controllers

import (
	"TajikCareerHub/logger"
	"TajikCareerHub/models"
	"TajikCareerHub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetFavoritesByUserID(c *gin.Context) {
	ip := c.ClientIP()
	userIDStr := c.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	favorites, err := service.GetFavoritesByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve favorites"})
		return
	}
	logger.Info.Printf("[controllers.GetFavoritesByUserID] Client IP: %s - Successfully retrieved favorites for user ID %v\n", ip, userID)
	c.JSON(http.StatusOK, favorites)
}

func GetFavoriteByUserIDAndJobID(c *gin.Context) {
	ip := c.ClientIP()
	userIDStr := c.Param("userID")
	jobIDStr := c.Param("jobID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	jobID, err := strconv.ParseUint(jobIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	favorite, err := service.GetFavoriteByUserIDAndJobID(uint(userID), uint(jobID))
	if err != nil {
		logger.Info.Printf("[controllers.GetFavoriteByUserIDAndJobID] Client IP: %s - Client requested favorite for user ID %v and job ID %v. Error retrieving favorite.\n", ip, userID, jobID)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetFavoriteByUserIDAndJobID] Client IP: %s - Successfully retrieved favorite for user ID %v and job ID %v\n", ip, userID, jobID)
	c.JSON(http.StatusOK, favorite)
}

func AddFavorite(c *gin.Context) {
	ip := c.ClientIP()
	var favorite models.Favorite
	if err := c.ShouldBindJSON(&favorite); err != nil {
		handleError(c, err)
		return
	}
	if err := service.AddFavorite(favorite); err != nil {
		logger.Info.Printf("[controllers.AddFavorite] Client IP: %s - Client attempted to add favorite with data %v. Error adding favorite.\n", ip, favorite)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.AddFavorite] Client IP: %s - Successfully added favorite with data %v\n", ip, favorite)
	c.JSON(http.StatusCreated, gin.H{"message": "Favorite added successfully"})
}

func RemoveFavorite(c *gin.Context) {
	ip := c.ClientIP()
	userIDStr := c.Param("userID")
	jobIDStr := c.Param("jobID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	jobID, err := strconv.ParseUint(jobIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	if err := service.RemoveFavorite(uint(userID), uint(jobID)); err != nil {
		logger.Info.Printf("[controllers.RemoveFavorite] Client IP: %s - Client attempted to remove favorite for user ID %v and job ID %v. Error removing favorite.\n", ip, userID, jobID)
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.RemoveFavorite] Client IP: %s - Successfully removed favorite for user ID %v and job ID %v\n", ip, userID, jobID)
	c.JSON(http.StatusOK, gin.H{"message": "Favorite removed successfully"})
}

func CheckFavoriteExists(c *gin.Context) {
	ip := c.ClientIP()
	userIDStr := c.Param("userID")
	jobIDStr := c.Param("jobID")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}
	jobID, err := strconv.ParseUint(jobIDStr, 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	exists, err := service.CheckFavoriteExists(uint(userID), uint(jobID))
	if err != nil {
		logger.Info.Printf("[controllers.CheckFavoriteExists] Client IP: %s - Client checked if job ID %v is in favorites for user ID %v. Error checking favorite.\n", ip, jobID, userID)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.CheckFavoriteExists] Client IP: %s - Checked if job ID %v is in favorites for user ID %v. Exists: %v\n", ip, jobID, userID, exists)
	c.JSON(http.StatusOK, gin.H{"exists": exists})
}
