package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
	"postgres/database"
	"postgres/server"
)

func cleanup(db *gorm.DB) {
	if r := recover(); r != nil {
		fmt.Println("Recovered from: ", r)
		return
	}

	db.Close()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	db, cErr := database.Connect(dbHost, dbPort, dbName, dbUsername, dbPassword)

	if cErr != nil {
		panic(err)
	}

	server.Db = db

	defer cleanup(db)

	//database.ClearTables(db)
	database.Migrate(db)
	//dummyData.InsertDummyData(db)

	server.StartServer()
}
