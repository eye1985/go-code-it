package server

import (
	"codepocket/database"
	"codepocket/enum"
	"codepocket/validate"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func validateForm(username string, password string, rPassword string, email string) string {
	var err error
	var errStr string

	err = validate.Username(username)
	if err != nil {
		errStr += fmt.Sprintf("%v \n", err.Error())
	}
	err = validate.Password(password)

	if password != rPassword {
		errStr += fmt.Sprintf("%v \n", "Repeated password not the same")
	}

	if err != nil {
		errStr += fmt.Sprintf("%v \n", err.Error())
	}

	err = validate.Email(email)
	if err != nil {
		errStr += fmt.Sprintf("%v \n", err.Error())
	}

	return errStr
}

func createUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	repeatPassword := r.FormValue("repeatPassword")
	email := r.FormValue("email")

	errorMsg := validateForm(username, password, repeatPassword, email)
	if len(errorMsg) > 0 {
		http.Error(w, errorMsg, http.StatusUnprocessableEntity)
		return
	}

	user, err := database.CreateUser(Db, &database.User{
		Username: &username,
		Password: &password,
		Email:    &email,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]

	var user database.User
	err := database.QueryOne(Db, "id = ?", id, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	repeatPassword := r.FormValue("repeatPassword")

	errorMsg := validateForm(username, password, repeatPassword, email)
	if len(errorMsg) > 0 {
		http.Error(w, errorMsg, http.StatusUnprocessableEntity)
		return
	}

	user, err := database.UpdateUser(Db, userId, username, email, password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]

	var user database.User
	var err error
	err = database.QueryOne(Db, "id = ?", id, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = database.DeleteUser(Db, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []database.User
	err := database.Query(Db, &users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&users)
}
