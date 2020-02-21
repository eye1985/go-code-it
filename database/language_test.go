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
		Language: "Java",
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

	if cLang.Language != "Java" {
		t.Fatalf("Result should be Java, but is %v", cLang.Language)
	}

	log.Printf("Language %v created successfully \n", cLang.Language)
}

func TestCreateDuplicateLanguage(t *testing.T) {
	_, err := CreateLanguage(ldb, &Language{
		Language: "Java",
	})

	if err == nil {
		t.Fatal("Should throw unique error")
	}

	log.Printf("Error %v thrown successfully \n", err)
}

func TestCreateNilLanguage(t *testing.T) {
	_, err := CreateLanguage(ldb, &Language{
		Language: "",
	})

	if err == nil {
		t.Fatal("Should throw nil error")
	}

	log.Printf("Error %v thrown successfully \n", err.Error())
}

func TestUpdateLanguage(t *testing.T) {

	_, goGetErr := CreateLanguage(ldb, &Language{
		Language: "Python",
	})

	if goGetErr != nil {
		t.Fatal(goGetErr)
	}

	pyLang, getErr := GetLanguage(ldb, &Language{
		Language: "Python",
	})

	if getErr != nil {
		t.Fatal(getErr)
	}

	_, err := UpdateLanguage(ldb, &Language{
		Model: gorm.Model{
			ID: pyLang.ID,
		},
		Language: "Kotlin",
	})

	if err != nil {
		t.Fatal(err)
	}

	updated, err2 := GetLanguage(ldb, &Language{
		Language: "Kotlin",
	})

	if err2 != nil {
		t.Fatal(err)
	}

	if updated.Language != "Kotlin" {
		t.Fatalf("Language should be Kotlin, but is %v", updated.Language)
	}

	log.Printf("Successfully updated to %v \n", updated.Language)
}

func TestGetLanguages(t *testing.T) {
	langs, err := GetLanguages(ldb)

	if err != nil {
		t.Fatal(err)
	}

	if len(*langs) == 0 {
		t.Fatal("Should not be empty")
	}

	var createdLang []string

	for _, lang := range *langs {
		if lang.Language == "Kotlin" || lang.Language == "Python" {
			createdLang = append(createdLang, lang.Language)
		}
	}

	if len(createdLang) != 2 {
		t.Fatalf("Language list should be length of 2, but is %v", createdLang)
	}

	if createdLang[0] != "Python" {
		t.Fatalf("Should be Python, but is %v", createdLang[1])
	}

	if createdLang[1] != "Kotlin" {
		t.Fatalf("Should be Kotlin, but is %v", createdLang[0])
	}

	log.Printf("Successfully retrieved %v \n", (*langs)[0].Language)
}

func TestDeleteLanguage(t *testing.T) {
	deleted, err := DeleteLanguage(ldb, &Language{
		Language: "Kotlin",
	})

	if err != nil {
		t.Fatal(err)
	}

	langs, err2 := GetLanguages(ldb)

	if err2 != nil {
		t.Fatal(err2)
	}

	if len(*langs) != 11 {
		t.Fatalf("Length of Languages should be 11, but is %v", len(*langs))
	}

	log.Printf("%v successfully deleted \n", deleted.Language)
	ldb.Close()
}
