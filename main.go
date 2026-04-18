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
	mux.HandleFunc("POST /threads", handlers.CreateThreadHandler)
	mux.HandleFunc("GET /posts", handlers.GetPostsHandler)
	mux.HandleFunc("GET /posts/{id}", handlers.GetPostDetailHandler)
	mux.HandleFunc("POST /posts", handlers.CreatePostHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fileServer)

	handlersWithCors := handlers.CorsMiddleware(mux)
	handlersWithLateLimit := handlers.RateLimitMiddleware(handlersWithCors)

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", handlersWithLateLimit); err != nil {
		log.Fatal(err)
	}
}
