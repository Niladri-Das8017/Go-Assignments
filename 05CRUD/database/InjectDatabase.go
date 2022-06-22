package database

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

var contactList *mongo.Collection

//Middleware to inject database
func InjectDatabase(controller http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {


		//passing database using context
		ctx := context.WithValue(r.Context(), "col", ContactsCollection)

		controller.ServeHTTP(w, r.WithContext(ctx))
	}
}
