package models

import (
	"TajikCareerHub/errs"
	"strings"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"full_name" gorm:"type:varchar(255);not null"`
	UserName  string    `json:"username" gorm:"type:varchar(100);unique;not null"`
	BirthDate time.Time `json:"birth_date" gorm:"type:date"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Role      string    `json:"role" gorm:"type:varchar(255);not null"`
	IsBlocked bool      `json:"is_blocked" gorm:"type:bool;not null;default:false"`
	BaseModel
}

func (u User) ValidateCredentials() error {
	if strings.TrimSpace(u.UserName) == "" {
		return errs.ErrUsernameExists
	}
	if strings.TrimSpace(u.Email) == "" {
		return errs.ErrEmailExists
	}
	if strings.TrimSpace(u.Role) == "" {
		return errs.ErrRoleExist
	}
	if strings.ToLower(u.Role) == "admin" {
		return errs.ErrRoleCannotBeAdmin
	}
	if strings.TrimSpace(u.Role) != "specialist" && strings.TrimSpace(u.Role) != "employer" {
		return errs.ErrInvalidRole
	}
	if len(u.Password) < 8 {
		return errs.ErrIncorrectPasswordLength
	}
	return nil
}

type SwagUser struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type SwagInUser struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}
