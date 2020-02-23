package main

import (
	"codepocket/database"
	"codepocket/server"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
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
	serverPort := os.Getenv("SERVER_PORT")
	serverHost := os.Getenv("SERVER_HOST")
	secret := os.Getenv("SESSION_SECRET")

	db, cErr := database.Connect(dbHost, dbPort, dbName, dbUsername, dbPassword)

	if cErr != nil {
		log.Println("Cannot connect to database: " + cErr.Error())
		os.Exit(3)
	}

	server.Db = db

	defer cleanup(db)

	database.ClearTables(db)
	database.Migrate(db)
	database.InsertDummyData(db)

	server.InitSession(secret)
	server.StartServer(serverHost, serverPort)
}
