package router

import (
	"github.com/gorilla/mux"
	"go-todo/middleware"
)

var URL = "https://todolist-backend79.herokuapp.com"

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc(URL+"/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc(URL+"/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc(URL+"/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc(URL+"/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc(URL+"/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc(URL+"/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
