package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"postgres/enum"
)

func RoutesV1() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	var api = router.PathPrefix(enum.APIVersion).Subrouter()

	//Not found
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	//Public
	api.HandleFunc("/login", login).Methods("POST")

	api.HandleFunc("/user", createUser).Methods("POST")
	api.HandleFunc("/users", getUsers).Methods("GET")
	api.HandleFunc("/codes", getCodes).Methods("GET")

	// Logout
	authPath := api.PathPrefix("/logout").Subrouter()
	authPath.HandleFunc("", logout).Methods("POST")
	authPath.Use(logoutAuth)

	// Users
	users := api.PathPrefix("/user").Subrouter()
	users.HandleFunc("/{userId}", getUser).Methods("GET")
	users.HandleFunc("/{userId}", updateUser).Methods("PUT")
	users.HandleFunc("/{userId}", deleteUser).Methods("DELETE")
	users.Use(auth)

	//Code
	codes := api.PathPrefix("/user/{userId}/code").Subrouter()
	codes.HandleFunc("", getUserCodes).Methods("GET")
	codes.HandleFunc("", createUserCode).Methods("POST")
	codes.HandleFunc("/{codeId}", getUserCode).Methods("GET")
	codes.HandleFunc("/{codeId}", updateUserCode).Methods("PUT")
	codes.HandleFunc("/{codeId}", deleteUserCode).Methods("DELETE")
	codes.Use(auth)

	api.Use(noCache, logger)

	return router
}
