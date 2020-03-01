package server

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var Db *gorm.DB
var store *sessions.CookieStore

const cookieName = "auth"

func InitSession(secret string) {
	store = CreateSessionStore(secret)
	store.Options = &sessions.Options{
		MaxAge:   60 * 30,
		HttpOnly: true,
	}
}

func StartServer(host string, port string) {
	router := RoutesV1()
	fmt.Println("Starting server at port " + port)
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
