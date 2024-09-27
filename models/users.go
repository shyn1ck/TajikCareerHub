package models

import (
	"TajikCareerHub/utils/errs"
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
	RoleID    uint      `json:"role_id" gorm:"not null"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleID"`
	IsBlocked bool      `json:"-" gorm:"type:bool;not null;default:false"`
	BaseModel
}

func (u User) ValidateCredentials() (err error) {
	if strings.TrimSpace(u.UserName) == "" {
		return errs.ErrUsernameExists
	}

	if strings.TrimSpace(u.Email) == "" {
		return errs.ErrEmailExists
	}

	if u.RoleID == 0 {
		return errs.ErrRoleExist
	}

	if len(u.Password) < 8 {
		return errs.ErrIncorrectPasswordLength
	}

	switch u.RoleID {
	case 1:
		return errs.ErrRoleCannotBeAdmin
	case 2, 3:
	default:
		return errs.ErrInvalidRole
	}
	return nil
}

type SwagUser struct {
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password" gorm:"not null"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	RoleID    uint      `json:"role_id"`
	BirthDate time.Time `json:"birth_date"`
}

type SwagInUser struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}

type UpdateUser struct {
	FullName  string    `json:"full_name"`
	UserName  string    `json:"username"`
	BirthDate time.Time `json:"birth_date"`
	Email     string    `json:"email"`
}
