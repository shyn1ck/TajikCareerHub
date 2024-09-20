package service

import (
	"TajikCareerHub/configs"
	"TajikCareerHub/utils/errs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
)

type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userID uint, username, role string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(configs.AppSettings.AuthParams.JwtTtlMinutes)).Unix(),
			Issuer:    configs.AppSettings.AppParams.ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.ErrUnexpectedSigningMethod
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, errs.ErrTokenParseError
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errs.ErrInvalidToken
}

func GetUserIDFromToken(c *gin.Context) (uint, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return 0, errs.ErrAuthorizationHeaderMissing
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	claims, err := ParseToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}

func GetRoleFromToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return "", errs.ErrAuthorizationHeaderMissing
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.Role, nil
}

func GetUsernameFromToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return "", errs.ErrAuthorizationHeaderMissing
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.Username, nil
}
