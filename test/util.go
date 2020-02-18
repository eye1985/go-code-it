package test

import (
	"bytes"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"postgres/database"
	"postgres/server"
	"strconv"
	"strings"
	"testing"
)

func connectToDb() (*gorm.DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	db, err := database.Connect(dbHost, dbPort, dbName, dbUsername, dbPassword)

	if err != nil {
		return nil, err
	}

	return db, nil
}

var testUrl = "http://localhost:3000"

func request(t *testing.T, method string, path string, data interface{}, status int) *http.Response {

	var req *http.Request
	var err error

	switch data.(type) {
	case []byte:
		byteData := data.([]byte)
		req, err = http.NewRequest(method, testUrl+path, bytes.NewBuffer(byteData))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set(server.ContentType, server.AppJson)
	case string:
		req, err = http.NewRequest(method, testUrl+path, strings.NewReader(data.(string)))

		if err != nil {
			t.Fatal(err)
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.(string))))
	case nil:
		req, err = http.NewRequest(method, testUrl+path, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set(server.ContentType, server.AppJson)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != status {
		t.Error("Not status ok: ", res.StatusCode)
	}

	return res
}
