package server

import (
	"codepocket/database"
	"codepocket/enum"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getCodes(w http.ResponseWriter, r *http.Request) {

	res, err := database.QueryAllCodes(Db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&res)
}

func getUserCodes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	i, err := strconv.Atoi(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, uerr := database.QueryUserCodes(Db, i)

	if uerr != nil {
		http.Error(w, uerr.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&user)
}

func getUserCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	codeId := params["codeId"]

	codeIdInt, err := strconv.Atoi(codeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userIdInt, err2 := strconv.Atoi(userId)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

	code, err3 := database.QueryUserCode(Db, userIdInt, codeIdInt)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&code)
}

func createUserCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	codeTitle := r.FormValue("title")
	codeDesc := r.FormValue("description")
	code := r.FormValue("code")
	languageId := r.FormValue("languageId")

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userIdUint := uint(userIdInt)

	languageIdInt, err := strconv.Atoi(languageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	languageIdUint := uint(languageIdInt)

	createdCode, err2 := database.CreateUserCode(Db, userId, &database.Code{
		UserID:      &userIdUint,
		LanguageID:  &languageIdUint,
		Title:       codeTitle,
		Description: codeDesc,
		Code:        &code,
	})

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&createdCode)
}

func updateUserCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	codeId := params["codeId"]

	codeTitle := r.FormValue("title")
	codeDesc := r.FormValue("description")
	code := r.FormValue("code")

	codeIdInt, err := strconv.Atoi(codeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedCode, err := database.UpdateUserCode(Db, codeIdInt, userIdInt, &database.Code{
		Title:       codeTitle,
		Description: codeDesc,
		Code:        &code,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&updatedCode)
}

func deleteUserCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	codeId := params["codeId"]

	codeIdInt, err := strconv.Atoi(codeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	deleteCode, err := database.DeleteUserCode(Db, codeIdInt, userIdInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&deleteCode)
}
