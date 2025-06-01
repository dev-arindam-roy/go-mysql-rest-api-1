package controllers

import (
	"encoding/json"
	"mysql-prac1/config"
	"mysql-prac1/models"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	result := config.DB.Find(&users)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	result := config.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
