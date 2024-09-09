package models

type Vacancy struct {
	ID                uint            `gorm:"primaryKey"`
	Title             string          `json:"title"`
	Description       string          `json:"description"`
	Location          string          `json:"location"`
	Salary            float64         `json:"salary"`
	CompanyID         uint            `json:"company_id"`
	Company           Company         `gorm:"foreignKey:CompanyID"` // Foreign key
	VacancyCategoryID uint            `json:"vacancy_category_id"`
	VacancyCategory   VacancyCategory `gorm:"foreignKey:VacancyCategoryID"` // Foreign key
	BaseModel
}
