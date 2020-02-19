package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"postgres/enum"
)

var Db *gorm.DB

func RoutesV1() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	var api = router.PathPrefix(enum.APIVersion).Subrouter()

	//Not found
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	// Users
	api.HandleFunc("/user", createUser).Methods("POST")
	api.HandleFunc("/user/{id}", getUser).Methods("GET")
	api.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	api.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	api.HandleFunc("/users", getUsers).Methods("GET")

	//Code
	api.HandleFunc("/codes", getCodes).Methods("GET")
	api.HandleFunc("/user/{id}/code", getUserCodes).Methods("GET")
	api.HandleFunc("/user/{id}/code", createUserCode).Methods("POST")
	api.HandleFunc("/user/{userId}/code/{codeId}", getUserCode).Methods("GET")
	api.HandleFunc("/user/{userId}/code/{codeId}", updateUserCode).Methods("PUT")
	api.HandleFunc("/user/{userId}/code/{codeId}", deleteUserCode).Methods("DELETE")

	api.Use(noCache, logger)

	return api
}

func StartServer(host string, port string) {
	router := RoutesV1()
	fmt.Println("Starting server at port " + port)
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
