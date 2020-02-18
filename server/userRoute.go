package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"postgres/database"
)

type TestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var post database.User
	var err error
	err = json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.CreateUser(Db, post.Username, post.Email, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&post)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user database.User
	err := database.QueryOne(Db, "id = ?", id, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(ContentType, AppJson)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	username := r.FormValue("username")
	email := r.FormValue("email")

	if len(username) == 0 || len(username) == 0 {
		http.Error(w, "Username or email missing", http.StatusUnprocessableEntity)
		return
	}

	user, err := database.UpdateUser(Db, userId, username, email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(ContentType, AppJson)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user database.User
	var err error
	err = database.QueryOne(Db, "id = ?", id, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = database.SoftDeleteUser(Db, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(ContentType, AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []database.User
	err := database.Query(Db, &users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set(ContentType, AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&users)
}
