package server

import (
	"codepocket/database"
	"encoding/json"
	"net/http"
)

func getLanguage(w http.ResponseWriter, r *http.Request) {
	languages, err := database.GetLanguages(Db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&languages)
}
