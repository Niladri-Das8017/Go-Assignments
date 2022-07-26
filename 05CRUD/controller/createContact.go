package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/niladridas/crud/database"
	"github.com/niladridas/crud/model"
)

func CreateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcaton/json")

	var contact model.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened"))
	}

	if contact.Name == "" || contact.Number == "" || len(contact.Number) != 10 {
		if len(contact.Number) != 10 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Invalid Mobile No."))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Invalid Details"))
			return
		}
	}

	ctx := r.Context()

	//insert  1 contact
	inserted, err := database.CreateContact(contact, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Contact cannot be created"))
		return
	}

	// oid := inserted.InsertedID.(primitive.ObjectID)
	// id := oid.Hex()
	fmt.Println("contact created")
	w.Write([]byte("Contact created : \n"))
	json.NewEncoder(w).Encode(inserted.InsertedID)
	json.NewEncoder(w).Encode(contact.Name)
	json.NewEncoder(w).Encode(contact.Number)

}
