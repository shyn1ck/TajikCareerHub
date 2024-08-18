package models

type JobCategory struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"varchar(100)"`
	BaseModel
}
