package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var Db *gorm.DB

const (
	appJson     = "application/json"
	contentType = "Content-Type"
)

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)

	// Users
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/users", getUsers).Methods("GET")

	//Code
	router.HandleFunc("/codes", getCodes).Methods("GET")
	router.HandleFunc("/user/{id}/code", getUserCodes).Methods("GET")
	router.HandleFunc("/user/{userId}/code/{codeId}", getUserCode).Methods("GET")
	router.HandleFunc("/user/{userId}/code/{codeId}", updateUserCode).Methods("PUT")
	router.HandleFunc("/user/{userId}/code/{codeId}", deleteUserCode).Methods("DELETE")

	router.Use(logger)

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
