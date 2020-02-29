package server

import (
	"codepocket/database"
	"codepocket/enum"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"math"
	"net/http"
	"strconv"
)

func getCodes(w http.ResponseWriter, r *http.Request) {
	var pagination database.Pagination

	start := r.URL.Query().Get("start")
	hitPerPage := r.URL.Query().Get("hitPerPage")

	startInt, err := strconv.Atoi(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hitPerPageInt, err := strconv.Atoi(hitPerPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, userAndCodes, err := database.SearchCodes(Db, "", int16(startInt), int16(hitPerPageInt))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalPage := math.Ceil(float64(*count) / float64(hitPerPageInt))
	currentPage := math.Ceil(float64(startInt) / float64(hitPerPageInt))
	next := startInt + hitPerPageInt
	prev := startInt - hitPerPageInt

	if len(userAndCodes) == 0 {
		totalPage = 0
		currentPage = 0
	}

	if prev < 0 {
		prev = 0
	}

	pagination = database.Pagination{
		Codes:       userAndCodes,
		CurrentPage: int16(currentPage),
		NextStart:   int16(next),
		PrevStart:   int16(prev),
		TotalPage:   int16(totalPage),
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&pagination)
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

	log.Printf("code public is %v", code.Public)

	if !code.Public {
		session, _ := store.Get(r, cookieName)
		u, err := database.GetUser(Db, &database.User{
			Model: gorm.Model{
				ID: uint(userIdInt),
			},
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if session.Values["auth"] != u.ID {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
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
