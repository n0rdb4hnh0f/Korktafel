package models

import (
	"log"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Post struct {
	ID        string    `gorm:"primaryKey;size:21"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Content   string         `gorm:"type:text;not null"`
}

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&Post{})

}

func (u *Post) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := gonanoid.New()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}
