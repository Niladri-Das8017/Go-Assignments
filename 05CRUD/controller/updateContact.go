package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niladridas/crud/database"
	"github.com/niladridas/crud/model"
)

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcaton/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT") //only allow PUT

	params := mux.Vars(r)

	var contact model.Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)

	ctx := r.Context()

	//update 1 record
	err := database.UpdateContact(params["id"], contact.Name, contact.Number, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Failed to update contact! \n"))
		return
	}

	fmt.Println("contact updated")
	w.Write([]byte("Contact updated.\nid: "))
	json.NewEncoder(w).Encode(params["id"])
}
