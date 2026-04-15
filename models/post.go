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
	Base
	Author    string         `gorm:"type:varchar(100);not null;default:'名無し'" json:"author"`
	ThreadID  string         `gorm:"index" json:"thread_id"`
	CreatedAt time.Time      `gorm:"index" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // DeletedAtは隠すのが一般的
	Content   string         `gorm:"column:content;type:text;not null" json:"content"`
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
