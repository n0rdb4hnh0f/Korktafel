package main

import (
	"fmt"
	"log"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Post struct {
	ID        string    `gorm:"primaryKey;size:21"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Content   string         `gorm:"type:text;not null"`
}

func (u *Post) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := gonanoid.New()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Post{})

	newPost := Post{Content: "It's test."}
	db.Create(&newPost)

	fmt.Printf("生成されたID: %s\n", newPost.ID)
	fmt.Printf("投稿内容: %s\n", newPost.Content)
	fmt.Printf("作成日時: %v\n", newPost.CreatedAt)
}
