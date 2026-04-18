package models

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Post struct {
	Base
	ThreadID  string         `gorm:"index;not null" json:"thread_id"`
	Author    string         `gorm:"type:varchar(100);not null;default:'名無し'" json:"author"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&Thread{}, &Post{})

}

func (u *Base) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := gonanoid.New()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}
