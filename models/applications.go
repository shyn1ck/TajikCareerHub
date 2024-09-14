package models

type Application struct {
	ID        uint              `json:"id" gorm:"primaryKey"`
	UserID    uint              `json:"user_id" gorm:"not null"`
	User      User              `json:"user" gorm:"foreignKey:UserID"`
	VacancyID uint              `json:"vacancy_id" gorm:"not null"`
	Vacancy   Vacancy           `json:"vacancy" gorm:"foreignKey:VacancyID"`
	ResumeID  uint              `json:"resume_id" gorm:"not null"`
	Resume    Resume            `json:"resume" gorm:"foreignKey:ResumeID"`
	StatusID  uint              `json:"status_id" gorm:"not null"`
	Status    ApplicationStatus `json:"status" gorm:"foreignKey:StatusID"`
	BaseModel
}

type ApplicationStatus struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(50);not null"`
}
