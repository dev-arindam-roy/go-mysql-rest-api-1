package controllers

import (
	"encoding/json"
	"mysql-prac1/config"
	"mysql-prac1/models"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	result := config.DB.Find(&posts)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	// Decode JSON and check for errors
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validation
	if post.Title == "" || post.Content == "" || post.UserID == 0 {
		http.Error(w, "All fields (title, content, user_id) are required", http.StatusBadRequest)
		return
	}

	// Create the post
	result := config.DB.Create(&post)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	json.NewEncoder(w).Encode(post)
}
