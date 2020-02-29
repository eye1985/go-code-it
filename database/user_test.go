package database

import (
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {

	un := "Abu"
	p := "somepassxyz"
	e := "asdasd@asdasd.com"

	u, err := CreateUser(tdb, &User{
		Username: &un,
		Password: &p,
		Email:    &e,
	})

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%v successfully created", u.Username)
}

func TestGetUser(t *testing.T) {
	un := "Abu"
	u, err := GetUser(tdb, &User{
		Username: &un,
	})

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%v successfully retrieved", u.Username)
}
