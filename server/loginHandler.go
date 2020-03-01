package server

import (
	"codepocket/database"
	"codepocket/encrypt"
	"codepocket/enum"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	username := r.FormValue("username")
	password := r.FormValue("password")

	u, err := database.GetUserAndRole(Db, &database.User{
		Username: &username,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	ok, err := encrypt.ComparePasswords(u.Password, []byte(password))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "Incorrect password", http.StatusForbidden)
		return
	}

	session.Values["auth"] = u.ID
	session.Save(r, w)

	userId := fmt.Sprint(u.ID)

	_, updateErr := database.UpdateUser(Db, userId, &database.User{
		Session: &userId,
	})

	if updateErr != nil {
		http.Error(w, updateErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	userIdStruct := struct {
		UserId string
	}{
		UserId: userId,
	}

	json.NewEncoder(w).Encode(&userIdStruct)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	if !session.Options.HttpOnly {
		http.Error(w, "Not httponly cookie", http.StatusForbidden)
		return
	}

	userIdStr := fmt.Sprint(session.Values["auth"])
	userIdUint, err := strconv.ParseUint(userIdStr, 10, 32)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u, uErr := database.GetUser(Db, &database.User{
		Model: gorm.Model{
			ID: uint(userIdUint),
		},
	})

	if uErr != nil {
		http.Error(w, uErr.Error(), http.StatusNotFound)
		return
	}

	if u.Session == nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}

	sessionStr := fmt.Sprint(session.Values["auth"])

	if sessionStr != *u.Session {
		http.Error(w, "Not same session", http.StatusUnauthorized)
		return
	}

	_, updateErr := database.UpdateUser(Db, userIdStr, &database.User{
		Session: nil,
	})

	if updateErr != nil {
		http.Error(w, updateErr.Error(), http.StatusNotFound)
		return
	}

	session.Values["auth"] = nil
	session.Options.MaxAge = -1
	session.Save(r, w)

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)
}
