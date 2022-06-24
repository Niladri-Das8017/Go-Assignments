package database

import (
	"context"
	"fmt"

	"github.com/niladridas/crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
