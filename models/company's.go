package models

type Company struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"varchar(100);unique"`
	Description string `json:"description" gorm:"text"`
	BaseModel
}
