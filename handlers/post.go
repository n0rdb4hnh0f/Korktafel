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

	models.DB.Order("created_at desc").Find(&posts)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
