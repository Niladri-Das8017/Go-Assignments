package controller

import (
	"encoding/json"
	"net/http"

	"github.com/niladridas/crud/database"
)

func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcaton/json")

	ctx := r.Context()

	allContacts, err := database.GetAllContacts(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Failed to get contacts! \n"))
		return
	}

	w.Write([]byte("Contact List : \n"))
	json.NewEncoder(w).Encode(allContacts)
}
