package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/n0rdb4hnh0f/GoBBS-API/handlers"
	"github.com/n0rdb4hnh0f/GoBBS-API/models"
)

func main() {
	models.InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /threads", handlers.GetThreadsHandler)
	mux.HandleFunc("GET /threads/{id}", handlers.GetThreadDetailHandler)

	createThreadHandler := http.HandlerFunc(handlers.CreateThreadHandler)
	mux.Handle("POST /threads", handlers.RateLimitMiddleware(createThreadHandler))

	mux.HandleFunc("GET /posts", handlers.GetPostsHandler)
	mux.HandleFunc("GET /posts/{id}", handlers.GetThreadDetailHandler)

	createPostHandler := http.HandlerFunc(handlers.CreatePostHandler)
	mux.Handle("POST /posts", handlers.RateLimitMiddleware(createPostHandler))

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
