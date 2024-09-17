package models

import "time"

type VacancyApplication struct {
	ID                  uint              `json:"id" gorm:"primaryKey"`
	VacancyID           uint              `json:"vacancy_id" gorm:"not null"`
	Vacancy             Vacancy           `json:"vacancy" gorm:"foreignKey:VacancyID"`
	ApplicationID       uint              `json:"application_id" gorm:"not null"`
	Application         Application       `json:"application" gorm:"foreignKey:ApplicationID"`
	ApplicationStatusID uint              `json:"application_status_id" gorm:"not null"`
	ApplicationStatus   ApplicationStatus `json:"application_status" gorm:"foreignKey:ApplicationStatusID"`
	AppliedAt           time.Time         `json:"applied_at" gorm:"not null"`
	BaseModel
}
