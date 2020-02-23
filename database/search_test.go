package database

import (
	"codepocket/test"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"testing"
)

var sdb *gorm.DB

func TestMain(m *testing.M) {
	db, err := test.ConnectToDb(Connect, ClearTables, Migrate)

	if err != nil {
		log.Fatal(err)
	}
	sdb = db
	InsertDummyData(sdb)
	exitVal := m.Run()
	db.Close()
	log.Println("Search test finished")
	os.Exit(exitVal)
}

func TestSearchCodes(t *testing.T) {
	resPerPage := int16(20)
	offset := int16(21)

	count, userAndCodes, err := SearchCodes(sdb, offset, resPerPage)
	if err != nil {
		t.Fatal(err)
	}

	if *count != 66 && len(userAndCodes) != 66 {
		t.Fatal(err)
	}
}
