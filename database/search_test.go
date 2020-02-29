package database

import (
	"codepocket/test"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"testing"
)

var tdb *gorm.DB

func TestMain(m *testing.M) {
	db, err := test.ConnectToDb(Connect, ClearTables, Migrate)

	if err != nil {
		log.Fatal(err)
	}
	tdb = db
	InsertDummyData(tdb)
	exitVal := m.Run()
	db.Close()
	log.Println("Search test finished")
	os.Exit(exitVal)
}

func TestSearchCodes(t *testing.T) {
	resPerPage := int16(20)
	offset := int16(21)

	count, userAndCodes, err := SearchCodes(tdb, "", offset, resPerPage)
	if err != nil {
		t.Fatalf("Should retrieve codes, but got %v", err)
	}

	if *count != 68 && len(userAndCodes) != 68 {
		t.Fatalf("Should return 68, but got %v", *count)
	}

	log.Printf("Returned result: %v", *count)
}
