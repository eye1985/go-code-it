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
		//log.Printf("Cache header: %v \n", w.Header().Get("Cache-Control"))
		log.Printf("Content-Type: %v \n", w.Header().Get("Content-Type"))

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

func authHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

//func isPublicCode (next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request){
//		params := mux.Vars(r)
//		userId := params["userId"]
//		codeId := params["codeId"]
//
//		next.ServeHTTP(w, r)
//	}
//}

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

func asJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}
