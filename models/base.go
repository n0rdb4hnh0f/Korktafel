package models

import (
	"time"
)

type Base struct {
	ID        string    `gorm:"primaryKey;size:21" json:"id"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
