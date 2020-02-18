package test

import (
	"net/http"
	"testing"
)

//router.HandleFunc("/user/{id}/code", getUserCodes).Methods("GET")
//router.HandleFunc("/user/{userId}/code/{codeId}", getUserCode).Methods("GET")
//router.HandleFunc("/user/{userId}/code/{codeId}", updateUserCode).Methods("PUT")
//router.HandleFunc("/user/{userId}/code/{codeId}", deleteUserCode).Methods("DELETE")

func TestGetCodes(t *testing.T) {
	res := request(t, "GET", "/codes", nil, http.StatusOK)
	printResult(res)
}
