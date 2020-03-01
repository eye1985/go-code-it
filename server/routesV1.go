package server

import (
	"codepocket/enum"
	"github.com/gorilla/mux"
	"net/http"
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
	//api.HandleFunc("/search", search).Methods("GET")
	api.HandleFunc("/user", createUser).Methods("POST")
	api.HandleFunc("/users", getUsers).Methods("GET")
	api.HandleFunc("/codes", getCodes).Methods("GET")
	api.Use(cors)

	// Logout
	authPath := api.PathPrefix("/logout").Subrouter()
	authPath.HandleFunc("", logout).Methods("POST")

	// Users
	users := api.PathPrefix("/user").Subrouter()
	users.HandleFunc("/{userId}", getUser).Methods("GET")
	users.HandleFunc("/{userId}", updateUser).Methods("PUT")
	users.HandleFunc("/{userId}", deleteUser).Methods("DELETE")
	users.Use(auth)

	//Code
	codes := api.PathPrefix("/user/{userId}/code").Subrouter()
	codes.HandleFunc("", authHandle(getUserCodes)).Methods("GET")
	codes.HandleFunc("", authHandle(createUserCode)).Methods("POST")
	codes.HandleFunc("/{codeId}", getUserCode).Methods("GET")
	codes.HandleFunc("/{codeId}", authHandle(updateUserCode)).Methods("PUT")
	codes.HandleFunc("/{codeId}", authHandle(deleteUserCode)).Methods("DELETE")

	api.Use(logger)

	return router
}
