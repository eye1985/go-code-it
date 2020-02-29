package database

import (
	"fmt"
	"log"
	"testing"
)

func TestCreateUserCode(t *testing.T) {
	uname := "Mikkel"
	email := "m@m.com"
	password := "abcdefghiklkmn"

	user, cuErr := CreateUser(tdb, &User{
		Username: &uname,
		Email:    &email,
		Password: &password,
	})

	if cuErr != nil {
		t.Fatal(cuErr)
	}

	langId := uint(1)
	code := "bla bla bla"

	uc, ucErr := CreateUserCode(tdb, fmt.Sprint(user.ID), &Code{
		Title:       "My code",
		Description: "None",
		Code:        &code,
		LanguageID:  &langId,
		UserID:      &user.ID,
	})

	if ucErr != nil {
		t.Fatal(ucErr)
	}

	log.Printf("Successfully created code %v", uc.Title)
}

func TestQueryUserCode(t *testing.T) {
	code, err := QueryUserCode(tdb, 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	if code.Language != "Html" {
		t.Fatalf("Should be HTML, but is %v", code.Language)
	}

	if code.Title != "Private Alert code" {
		t.Fatalf("Should be Private Alert code, but is %v", code.Title)
	}
}
