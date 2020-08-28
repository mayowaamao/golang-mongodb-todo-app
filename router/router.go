package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/tasks/{id}/undo", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/tasks/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
