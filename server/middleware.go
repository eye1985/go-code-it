package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("URL: %v \n", r.RequestURI)
		log.Printf("Cache header: %v \n", w.Header().Get("Cache-Control"))

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, no-transform, must-revalidate, private, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("X-Accel-Expires", "0")
		next.ServeHTTP(w, r)
	})
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userId := params["userId"]
		session, _ := store.Get(r, cookieName)

		uid, ok := session.Values["auth"]
		uidStr := fmt.Sprintf("%v", uid)

		if !ok || uidStr != userId {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func logoutAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, cookieName)
		uid, ok := session.Values["auth"]

		if !ok || uid == 0 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
