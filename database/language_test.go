package database

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var db *gorm.DB

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
	db = tdb

	if cErr != nil {
		panic(err)
	}
}

func TestCreateLanguage(t *testing.T) {
	lang, err := CreateLanguage(db, &Language{
		Language: "Javascript",
	})

	if err != nil {
		t.Fatal(err)
	}

	cLang, err2 := GetLanguage(db, &Language{
		Model: gorm.Model{
			ID: lang.ID,
		},
	})

	if err2 != nil {
		t.Fatal(err2)
	}

	if cLang.Language != "Javascript" {
		t.Error("Result should be Javascript")
	}

	log.Printf("Language %v created successfully", cLang.Language)
}

func TestCreateDuplicateLanguage(t *testing.T) {
	_, err := CreateLanguage(db, &Language{
		Language: "Javascript",
	})

	if err == nil {
		t.Error("Should throw unique error")
	}

	log.Printf("Error %v thrown successfully", err)
}

func TestCreateNilLanguage(t *testing.T) {
	_, err := CreateLanguage(db, &Language{
		Language: "",
	})

	if err == nil {
		t.Error("Should throw nil error")
	}

	log.Printf("Error %v thrown successfully", err.Error())
}

func TestUpdateLanguage(t *testing.T) {

	jsLang, getErr := GetLanguage(db, &Language{
		Language: "Javascript",
	})

	if getErr != nil {
		t.Fatal(getErr)
	}

	_, err := UpdateLanguage(db, &Language{
		Model: gorm.Model{
			ID: jsLang.ID,
		},
		Language: "Java",
	})

	if err != nil {
		t.Fatal(err)
	}

	updated, err2 := GetLanguage(db, &Language{
		Language: "Java",
	})

	if err2 != nil {
		t.Fatal(err)
	}

	if updated.Language != "Java" {
		t.Error("Language should be Java")
	}

	log.Printf("Successfully updated to %v", updated.Language)
}

func TestGetLanguages(t *testing.T) {
	langs, err := GetLanguages(db)

	if err != nil {
		t.Fatal(err)
	}

	if len(*langs) == 0 {
		t.Error("Should not be empty")
	}

	if (*langs)[0].Language != "Java" {
		t.Error("Language should be Java")
	}

	log.Printf("Successfully retrieved %v", (*langs)[0].Language)
}

func TestDeleteLanguage(t *testing.T) {
	deleted, err := DeleteLanguage(db, &Language{
		Language: "Java",
	})

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%v successfully deleted", deleted.Language)
}
