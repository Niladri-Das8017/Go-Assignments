package database

import (
	"context"
	"fmt"

	"github.com/niladridas/crud/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateContact(contact model.Contact, ctx context.Context) (*mongo.InsertOneResult, error) {

	//parsing it to type dbIface.CollectionIface, so it will expect an interface.
	//now we are allowed to send any structure that impleements this interface.
	//In a word,  we can now send a mock database insted of orriginal database
	collection := ctx.Value("col").(CollectionIface)
	
	inserted, err := collection.InsertOne(context.Background(), contact)

	if err != nil {
		return nil, err
	}

	fmt.Println("Inserted 1 Contact in db with id: ", inserted.InsertedID)
	return inserted, nil
}
