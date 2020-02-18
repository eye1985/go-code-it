package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"postgres/database"
	"postgres/dummyData"
	"testing"
)

func init() {
	db, err := connectToDb()
	if err != nil {
		panic(err)
	}
	database.ClearTables(db)
	database.Migrate(db)
	dummyData.InsertDummyData(db)
}

func printResult(response *http.Response) {
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Response body: ", string(body))
}

func TestCreateUser(t *testing.T) {
	user := &database.User{
		Username: "Test user",
		Email:    "test@test.com",
	}

	jsonUser, _ := json.Marshal(&user)

	res := request(t, "POST", "/user", jsonUser, http.StatusCreated)
	printResult(res)
}

//router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

func TestGetUser(t *testing.T) {
	res := request(t, "GET", "/user/1", nil, http.StatusOK)
	printResult(res)
}

func TestGetUsers(t *testing.T) {
	res := request(t, "GET", "/users", nil, http.StatusOK)
	printResult(res)
}

func TestUpdateUser(t *testing.T) {
	form := url.Values{}
	form.Set("username", "new user name")
	form.Set("email", "newEmail@gmail.com")

	res := request(t, "PUT", "/user/1", form.Encode(), http.StatusOK)
	printResult(res)
}

func TestDeleteUser(t *testing.T) {
	res := request(t, "DELETE", "/user/1", nil, http.StatusOK)
	printResult(res)
}
