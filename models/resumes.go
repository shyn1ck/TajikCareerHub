package models

import (
	"TajikCareerHub/utils/errs"
	"strings"
)

type Resume struct {
	ID                uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	Title             string          `json:"title" gorm:"not null"`
	UserID            uint            `json:"user_id" gorm:"not null"`
	FullName          string          `json:"full_name" gorm:"type:varchar(255);not null"`
	Summary           string          `json:"summary" gorm:"type:text"`
	Skills            string          `json:"skills" gorm:"type:text"`
	ExperienceYears   uint            `json:"experience_years"`
	Education         string          `json:"education" gorm:"type:text"`
	Certifications    string          `json:"certifications" gorm:"type:text"`
	Location          string          `json:"location" gorm:"type:varchar(255)"`
	VacancyCategoryID uint            `json:"vacancy_category_id" gorm:"not null"`
	VacancyCategory   VacancyCategory `gorm:"foreignKey:VacancyCategoryID"`
	IsBlocked         bool            `json:"-" gorm:"default:false"`
	BaseModel
}

func (r Resume) ValidateResume() error {
	if strings.TrimSpace(r.FullName) == "" {
		return errs.ErrFullNameIsRequired
	}
	if r.VacancyCategoryID == 0 {
		return errs.ErrVacancyCategoryIsRequired
	}
	if r.ExperienceYears < 0 {
		return errs.ExperienceYearsCannotBeNegative
	}
	if len(r.Summary) > 1000 {
		return errs.SummaryCannotExceedDefiniteCharacters
	}
	return nil
}

type SwagResume struct {
	ID                uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName          string `json:"full_name" gorm:"type:varchar(255);not null"`
	Skills            string `json:"skills" gorm:"type:text"`
	Summary           string `json:"summary" gorm:"type:text"`
	ExperienceYears   uint   `json:"experience_years" gorm:"not null"`
	Location          string `json:"location" gorm:"type:varchar(255)"`
	VacancyCategoryID uint   `json:"vacancy_category_id" gorm:"not null"`
	Title             string `json:"title" gorm:"type:varchar(255)"`
}

type ResumeReport struct {
	ResumeID          uint   `json:"resume_id"`
	ResumeTitle       string `json:"resume_title"`
	ViewsCount        int    `json:"views_count"`
	ApplicationsCount int    `json:"applications_count"`
}
