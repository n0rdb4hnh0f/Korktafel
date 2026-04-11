package main

import (
	"fmt"
	"net/http"

	"github.com/n0rdb4hnh0f/GoBBS-API/handlers"
	"github.com/n0rdb4hnh0f/GoBBS-API/models"
)

func main() {
	models.InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /posts", handlers.GetPostsHandler)
	mux.HandleFunc("GET /posts/{id}", handlers.GetPostHandler)

	createHandler := http.HandlerFunc(handlers.CreatePostHandler)
	mux.Handle("POST /posts", handlers.RateLimitMiddleware(createHandler))

	mux.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
