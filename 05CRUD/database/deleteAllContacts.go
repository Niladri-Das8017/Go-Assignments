package database

import (
	"context"
	"fmt"


	"go.mongodb.org/mongo-driver/bson"
)

func DeleteAllContacts(ctx context.Context) (int64, error) {

	//parsing it to type dbIface.CollectionIface, so it will expect an interface.
	//now we are allowed to send any structure that implements this interface.
	//In a word,  we can now send a mock database insted of original database
	collection := ctx.Value("col").(CollectionIface)

	deleteResult, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		return -1, err
	}

	fmt.Println("Number  of  Contacts  deleted : ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount, nil
}
