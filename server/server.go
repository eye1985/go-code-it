package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var Db *gorm.DB

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)

	// Users
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/users", getUsers).Methods("GET")

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
