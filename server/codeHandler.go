package server

import (
	"codepocket/database"
	"codepocket/feature"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func startAndHitsPerPage(r *http.Request) (startInt int, hitPerPageInt int, err []error) {
	start := r.URL.Query().Get("start")
	hitPerPage := r.URL.Query().Get("hitPerPage")

	if hitPerPage == "" {
		hitPerPage = "10"
	}

	startInt, sErr := strconv.Atoi(start)
	hitPerPageInt, hErr := strconv.Atoi(hitPerPage)

	err = append(err, sErr)
	err = append(err, hErr)

	return startInt, hitPerPageInt, err
}

func getCodes(w http.ResponseWriter, r *http.Request) {
	startInt, hitPerPageInt, errList := startAndHitsPerPage(r)

	for _, err := range errList {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	count, userAndCodes, err := database.SearchCodes(Db, "", int16(startInt), int16(hitPerPageInt))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pagination := feature.Pagination(float64(*count), float64(hitPerPageInt), float64(startInt), userAndCodes)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&pagination)
}

func getUserCodes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	userIdInt, err := strconv.Atoi(userId)
	userIdUint := uint(userIdInt)

	startInt, hitPerPageInt, errList := startAndHitsPerPage(r)
	for _, err := range errList {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	count, userAndCodes, err := database.SearchUserCodes(Db, "", &userIdUint, int16(startInt), int16(hitPerPageInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pagination := feature.Pagination(float64(*count), float64(hitPerPageInt), float64(startInt), userAndCodes)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&pagination)
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&code)
}

func createUserCode(w http.ResponseWriter, r *http.Request) {
	var public bool
	params := mux.Vars(r)
	userId := params["userId"]

	codeTitle := r.FormValue("title")
	codeDesc := r.FormValue("description")
	code := r.FormValue("code")
	languageId := r.FormValue("languageId")
	publicStr := r.FormValue("public")

	if publicStr == "true" {
		public = true
	} else if publicStr == "false" {
		public = false
	}

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
		Public:      &public,
	})

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&deleteCode)
}
