package main

import (
	"github.com/cesc1802/go_training/controllers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/user/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/tasks", controllers.GetTask).Methods("GET")
	router.HandleFunc("/api/task/new", controllers.CreateTask).Methods("POST")
	http.ListenAndServe(":8000",router)
}
