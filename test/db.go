package test

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

func ConnectToDb(Connect func(host string, port string, dbname string, user string, pass string) (*gorm.DB, error), ClearTables func(db *gorm.DB), Migrate func(db *gorm.DB)) (*gorm.DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("TEST_DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	ndb, err := Connect(dbHost, dbPort, dbName, dbUsername, dbPassword)
	ClearTables(ndb)
	Migrate(ndb)

	if err != nil {
		return nil, err
	}

	return ndb, nil
}
