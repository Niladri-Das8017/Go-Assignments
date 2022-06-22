package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niladridas/crud/database"
)

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcaton/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") //only allow DELETE

	ctx := r.Context()
	//initializing Params
	params := mux.Vars(r)
	//calling handler
	err := database.DeleteContact(params["id"], ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Failed to delete contacts! \n"))
		return
	}
	w.Write([]byte("Contacts Deleted! \nDeleted id: "))
	json.NewEncoder(w).Encode(params["id"])
}
