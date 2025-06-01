package main

import (
	"log"
	"mysql-prac1/config"
	"mysql-prac1/routes"
	"net/http"
)

func main() {
	config.Connect()                  // Connect to MySQL
	router := routes.RegisterRoutes() // Register routes

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router)) // Use the router here
}
