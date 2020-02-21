package server

import (
	"net/http"
	"postgres/database"
	"postgres/enum"
)

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	username := r.FormValue("username")
	password := r.FormValue("password")

	u, err := database.GetUser(Db, &database.User{
		Username: &username,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if u.Password != password {
		http.Error(w, "Incorrect password", http.StatusForbidden)
		return
	}

	session.Values["auth"] = u.ID
	session.Save(r, w)

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	session.Values["auth"] = 0
	session.Save(r, w)

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)
}
