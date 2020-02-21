package database

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var udb *gorm.DB

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
	udb = tdb

	if cErr != nil {
		panic(err)
	}
}

func TestCreateUser(t *testing.T) {
	u, err := CreateUser(udb, "Abu", "bla@bla.com", "password")

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%v successfully created", u.Username)
}

func TestGetUser(t *testing.T) {
	u, err := GetUser(udb, &User{
		Username: "Abu",
	})

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%v successfully retrieved", u.Username)
}
