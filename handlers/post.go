package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/n0rdb4hnh0f/GoBBS-API/models"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if post.Author == "" {
		post.Author = "名無しさん"
	}

	if len(post.Author) > 50 {
		http.Error(w, "Name is too long(max 100 bytes)", http.StatusBadRequest)
		return
	}

	if post.Content == "" {
		http.Error(w, "No content", http.StatusBadRequest)
		return
	}

	if len(post.Content) > 49 {
		http.Error(w, "Too long name", http.StatusBadRequest)
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result := models.DB.Create(&post)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	result := models.DB.Order("created_at desc").Find(&posts)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var post models.Post
	result := models.DB.First(&post, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
