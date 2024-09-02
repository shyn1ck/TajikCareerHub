package controllers

import (
	"TajikCareerHub/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
	userRoleCtx         = "userRole"
)

func checkUserAuthentication(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}
	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return
	}
	accessToken := headerParts[1]
	claims, err := service.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("UserID:", claims.UserID, "Role:", claims.Role)
	c.Set(userIDCtx, claims.UserID)
	c.Set(userRoleCtx, claims.Role)
	c.Next()
}
func adminOnly(c *gin.Context) {
	role, exists := c.Get(userRoleCtx)
	if !exists || role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "access forbidden: admins only",
		})
		return
	}
	c.Next()
}
func employerOnly(c *gin.Context) {
	role, exists := c.Get(userRoleCtx)
	if !exists || role != "Employer" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "access forbidden: employers only",
		})
		return
	}
	c.Next()
}
func specialistOnly(c *gin.Context) {
	role, exists := c.Get(userRoleCtx)
	if !exists || role != "Specialist" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "access forbidden: specialists only",
		})
		return
	}
	c.Next()
}
