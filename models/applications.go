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

type SpecialistActivityReport struct {
	UserID           uint   `json:"-"`
	UserName         string `json:"full_name"`
	ApplicationCount uint   `json:"application_count"`
}

type SwaggerApplication struct {
	UserID    uint `json:"user_id" example:"1"`
	VacancyID uint `json:"vacancy_id" example:"1"`
	ResumeID  uint `json:"resume_id" example:"1"`
	StatusID  uint `json:"status_id" example:"1"`
}

type ResumeView struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint   `json:"user_id"`
	ResumeID uint   `json:"resume_id"`
	User     User   `gorm:"foreignKey:UserID"`
	Resume   Resume `gorm:"foreignKey:ResumeID"`
	Count    int    `json:"count" gorm:"default:0"`
}
