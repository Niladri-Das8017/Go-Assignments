package controller

import (
	"encoding/json"
	"net/http"

	"github.com/niladridas/crud/database"
)

func DeleteAllContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcaton/json")

	ctx := r.Context()

	//calling handler
	deleteCount, err := database.DeleteAllContacts(ctx)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Failed to Delete contact! \n"))
		return
	}

	w.Write([]byte("All Contacts Deleted! \nDelete Count: "))
	json.NewEncoder(w).Encode(deleteCount)
}
