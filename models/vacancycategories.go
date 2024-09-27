package models

type VacancyCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	BaseModel
}

type SwagVacancyCategories struct {
	Name string `json:"name"`
}
