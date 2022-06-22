package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateContact(contactId string, name string, number string, ctx context.Context) error {

	//parsing it to type dbIface.CollectionIface, so it will expect an interface.
	//now we are allowed to send any structure that implements this interface.
	//In a word,  we can now send a mock database insted of original database
	collection := ctx.Value("col").(CollectionIface)

	id, err := primitive.ObjectIDFromHex(contactId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	updates := bson.M{"$set": bson.M{"name": name, "number": number}}

	result, err := collection.UpdateOne(context.Background(), filter, updates)
	if err != nil {
		return err
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
	return nil
}
