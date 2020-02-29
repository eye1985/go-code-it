package database

import (
	"github.com/jinzhu/gorm"
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

func TestGetUserAndRole(t *testing.T) {
	un := "Abu"
	u, err := GetUserAndRole(tdb, &User{
		Username: &un,
	})

	if err != nil {
		t.Fatal(err)
	}

	if u.Username != "Abu" {
		t.Fatalf("Username should be Abu, but is %v", u.Username)
	}

	if u.Role != "USER" {
		t.Fatalf("Role should be USER, but is %v", u.Role)
	}

	log.Printf("%v successfully retrieved", u.Username)
}

func TestGetUser(t *testing.T) {
	un := "Abu"
	u, err := GetUser(tdb, &User{
		Username: &un,
	})

	if err != nil {
		t.Fatal(err)
	}

	if *u.Username != "Abu" {
		t.Fatalf("Username should be Abu, but is %v", *u.Username)
	}

	log.Printf("%v successfully retrieved", u.Username)
}

func TestDeleteUser(t *testing.T) {
	un := "Abu"
	u, err := GetUser(tdb, &User{
		Username: &un,
	})

	if err != nil {
		t.Fatal(err)
	}

	dErr := DeleteUser(tdb, u)

	if dErr != nil {
		t.Fatal(dErr)
	}

	_, gErr := GetUser(tdb, &User{
		Username: &un,
	})

	if gErr != gorm.ErrRecordNotFound {
		t.Fatal(dErr)
	}
}
