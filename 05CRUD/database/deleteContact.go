package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteContact(contactID string, ctx context.Context) error {

	//parsing it to type dbIface.CollectionIface, so it will expect an interface.
	//now we are allowed to send any structure that implements this interface.
	//In a word,  we can now send a mock database insted of original database
	collection := ctx.Value("col").(CollectionIface)

	id, err := primitive.ObjectIDFromHex(contactID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	fmt.Println("Contact got deleted with delete count: ", deleteCount)

	return nil
}
