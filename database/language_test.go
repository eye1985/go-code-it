package database

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var ldb *gorm.DB

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
	ldb = tdb

	if cErr != nil {
		panic(err)
	}
}

func TestCreateLanguage(t *testing.T) {
	lang, err := CreateLanguage(ldb, &Language{
		Language: "Javascript",
	})

	if err != nil {
		t.Fatal(err)
	}

	cLang, err2 := GetLanguage(ldb, &Language{
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

	log.Printf("Language %v created successfully \n", cLang.Language)
}

func TestCreateDuplicateLanguage(t *testing.T) {
	_, err := CreateLanguage(ldb, &Language{
		Language: "Javascript",
	})

	if err == nil {
		t.Error("Should throw unique error")
	}

	log.Printf("Error %v thrown successfully \n", err)
}

func TestCreateNilLanguage(t *testing.T) {
	_, err := CreateLanguage(ldb, &Language{
		Language: "",
	})

	if err == nil {
		t.Error("Should throw nil error")
	}

	log.Printf("Error %v thrown successfully \n", err.Error())
}

func TestUpdateLanguage(t *testing.T) {

	_, goGetErr := CreateLanguage(ldb, &Language{
		Language: "Go",
	})

	if goGetErr != nil {
		t.Fatal(goGetErr)
	}

	goLang, getErr := GetLanguage(ldb, &Language{
		Language: "Go",
	})

	if getErr != nil {
		t.Fatal(getErr)
	}

	_, err := UpdateLanguage(ldb, &Language{
		Model: gorm.Model{
			ID: goLang.ID,
		},
		Language: "Java",
	})

	if err != nil {
		t.Fatal(err)
	}

	updated, err2 := GetLanguage(ldb, &Language{
		Language: "Java",
	})

	if err2 != nil {
		t.Fatal(err)
	}

	if updated.Language != "Java" {
		t.Errorf("Language should be Java, but is %v", updated.Language)
	}

	log.Printf("Successfully updated to %v \n", updated.Language)
}

func TestGetLanguages(t *testing.T) {
	langs, err := GetLanguages(ldb)

	if err != nil {
		t.Fatal(err)
	}

	if len(*langs) == 0 {
		t.Error("Should not be empty")
	}

	if (*langs)[0].Language != "Go" {
		t.Errorf("Language should be Go %v", (*langs)[0].Language)
	}

	if (*langs)[1].Language != "Java" {
		t.Errorf("Language should be Java, but is %v", (*langs)[1].Language)
	}

	log.Printf("Successfully retrieved %v \n", (*langs)[0].Language)
}

func TestDeleteLanguage(t *testing.T) {
	deleted, err := DeleteLanguage(ldb, &Language{
		Language: "Java",
	})

	if err != nil {
		t.Fatal(err)
	}

	langs, err2 := GetLanguages(ldb)

	if err2 != nil {
		t.Fatal(err2)
	}

	if len(*langs) != 1 {
		t.Errorf("Length of Languages should be 1, but is %v", len(*langs))
	}

	log.Printf("%v successfully deleted \n", deleted.Language)
	ldb.Close()
}
