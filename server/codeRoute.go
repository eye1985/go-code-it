package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"postgres/database"
	"strconv"
)

func getCodes(w http.ResponseWriter, r *http.Request) {

	res, err := database.QueryAllCodes(Db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set(contentType, appJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&res)
}

func getUserCodes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	i, err := strconv.Atoi(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := database.QueryUserCodes(Db, i)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set(contentType, appJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&user)
}

func getUserCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	codeId := params["codeId"]

	var code database.Code

	if dbs := Db.Table("codes").Where("id = ? AND user_id = ?", codeId, userId); dbs.Error != nil {
		http.Error(w, dbs.Error.Error(), http.StatusNotFound)
		return
	} else {
		dbs.Scan(&code)
	}

	w.Header().Set(contentType, appJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&code)
}
