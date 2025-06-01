package routes

import (
	"mysql-prac1/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")

	router.HandleFunc("/comments", controllers.GetComments).Methods("GET")
	router.HandleFunc("/comments", controllers.CreateComment).Methods("POST")

	return router
}
