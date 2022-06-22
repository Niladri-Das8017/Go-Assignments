package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllContacts(ctx context.Context) ([]primitive.M, error) {

	//parsing it to type dbIface.CollectionIface, so it will expect an interface.
	//now we are allowed to send any structure that implements this interface.
	//In a word,  we can now send a mock database insted of original database
	collection := ctx.Value("col").(CollectionIface)

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	//array of movies
	var contacts []primitive.M //since we are using "M" primitive package

	for cursor.Next(context.Background()) {

		var contact bson.M
		//if cursor.Decode(&movie) worked,we have movie fieled up, else we will have an err.
		err := cursor.Decode(&contact) //whenever we decode, we pass on a referrence like, "if you decode use my structure to decode that"
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	defer cursor.Close(context.Background())
	return contacts, nil
}
