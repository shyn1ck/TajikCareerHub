package models

import (
	"TajikCareerHub/errs"
	"unicode/utf8"
)

type Vacancy struct {
	ID                uint            `gorm:"primaryKey"`
	Title             string          `json:"title"`
	Description       string          `json:"description"`
	Location          string          `json:"location"`
	Salary            float64         `json:"salary"`
	CompanyID         uint            `json:"company_id"`
	Company           Company         `gorm:"foreignKey:CompanyID"`
	User              User            `gorm:"foreignKey:UserID"`
	UserID            uint            `json:"user_id"`
	VacancyCategoryID uint            `json:"vacancy_category_id"`
	VacancyCategory   VacancyCategory `gorm:"foreignKey:VacancyCategoryID"`
	IsBlocked         bool            `json:"is_blocked" gorm:"default:false"`
	BaseModel
}

func (v Vacancy) ValidateVacancy() error {
	if len(v.Title) == 0 {
		return errs.ErrTitleIsRequired
	}
	if utf8.RuneCountInString(v.Title) > 100 {
		return errs.ErrTitleMustBeLessThanDefiniteCharacters
	}
	if len(v.Description) == 0 {
		return errs.ErrDescriptionIsRequired
	}
	if utf8.RuneCountInString(v.Description) > 1000 {
		return errs.ErrDescriptionMustBeLessThanDefiniteCharacters
	}
	if v.Salary < 0 {
		return errs.ErrSalaryMustBeANonNegativeNumber
	}
	if v.CompanyID == 0 {
		return errs.ErrCompanyIDIsRequired
	}
	if v.VacancyCategoryID == 0 {
		return errs.ErrVacancyCategoryIsRequired
	}
	return nil
}
