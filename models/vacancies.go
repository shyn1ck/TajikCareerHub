package models

import (
	"TajikCareerHub/utils/errs"
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
	VacancyViews      []VacancyView   `gorm:"foreignKey:VacancyID"`
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

type VacancyReport struct {
	VacancyID         uint   `json:"vacancy_id"`
	VacancyTitle      string `json:"vacancy_title"`
	ViewsCount        int64  `json:"views_count"`
	ApplicationsCount int64  `json:"applications_count"`
}

type SwagVacancy struct {
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	Location          string  `json:"location"`
	Salary            float64 `json:"salary"`
	CompanyID         uint    `json:"company_id"`
	VacancyCategoryID uint    `json:"vacancy_category_id"`
}

type VacancyView struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `json:"user_id"`
	VacancyID uint    `json:"vacancy_id"`
	User      User    `gorm:"foreignKey:UserID"`
	Vacancy   Vacancy `gorm:"foreignKey:VacancyID"`
	Count     int     `json:"count" gorm:"default:0"`
}
