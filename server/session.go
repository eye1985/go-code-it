package server

import (
	"github.com/gorilla/sessions"
)

func CreateSessionStore(secret string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(secret))
}
