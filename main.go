package main

import (
	"fmt"

	"github.com/n0rdb4hnh0f/GoBBS-API/models"
)

func main() {
	models.InitDB()
	newPost := models.Post{Content: "It's test."}
	models.DB.Create(&newPost)

	fmt.Printf("生成されたID: %s\n", newPost.ID)
	fmt.Printf("投稿内容: %s\n", newPost.Content)
	fmt.Printf("作成日時: %v\n", newPost.CreatedAt)
}
