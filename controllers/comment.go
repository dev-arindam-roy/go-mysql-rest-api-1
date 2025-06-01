package controllers

import (
	"encoding/json"
	"mysql-prac1/config"
	"mysql-prac1/models"
	"net/http"
)

func GetComments(w http.ResponseWriter, r *http.Request) {
	var comments []models.Comment
	result := config.DB.Find(&comments)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment

	// Decode JSON and check for errors
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validation
	if comment.UserID == 0 || comment.BlogID == 0 || comment.Comment == "" {
		http.Error(w, "All fields (user id, comment, blog id) are required", http.StatusBadRequest)
		return
	}

	// Create the post
	result := config.DB.Create(&comment)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	json.NewEncoder(w).Encode(comment)
}
