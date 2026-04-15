package models

type Thread struct {
	Base
	Title string `json:"title"`
	Posts []Post `gorm:"constraint:OnDelete:CASCADE;" json:"posts"`
}
