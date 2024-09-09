package models

type VacancyCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	BaseModel
}
