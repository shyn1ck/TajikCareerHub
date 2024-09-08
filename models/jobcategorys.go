package models

type JobCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	BaseModel
}
