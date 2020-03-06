package server

import (
	"codepocket/database"
	"codepocket/enum"
	"encoding/json"
	"net/http"
)

func getLanguage(w http.ResponseWriter, r *http.Request) {
	languages, err := database.GetLanguages(Db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set(enum.ContentType, enum.AppJson)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&languages)
}
