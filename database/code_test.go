package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var cdb *gorm.DB

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	tdb, cErr := Connect(dbHost, dbPort, dbName, dbUsername, dbPassword)
	ClearTables(tdb)
	Migrate(tdb)
	cdb = tdb

	if cErr != nil {
		panic(err)
	}
}

func TestCreateUserCode(t *testing.T) {
	uname := "Mikkel"
	email := "m@m.com"
	password := "abcdefghiklkmn"

	user, cuErr := CreateUser(cdb, &User{
		Username: &uname,
		Email:    &email,
		Password: &password,
	})

	if cuErr != nil {
		t.Fatal(cuErr)
	}

	langId := uint(1)
	code := "bla bla bla"

	uc, ucErr := CreateUserCode(cdb, fmt.Sprint(user.ID), &Code{
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
	cdb.Close()
}
