//Integration testing for helpers

package database

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/niladridas/crud/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testContactsCollection *mongo.Collection
var id string

//test model contact
var contact = model.Contact{
	Name:   "testName",
	Number: "9876543210",
}

//Initializing a duplicate database for test cases
func init() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//create DB and Collection
	testContactsDb := client.Database("contacts_test")
	testContactsCollection = testContactsDb.Collection("testContactList")
}

//Tests goes here...
func TestCreateContact(t *testing.T) {

	ctx := context.WithValue(context.Background(), "col", testContactsCollection)

	inserted, err := CreateContact(contact, ctx)
	if err != nil {
		t.Error(`CreateContact Test FAIILED! 
		err: `, err)
	}

	oid := inserted.InsertedID.(primitive.ObjectID)

	//inserting test id
	id = oid.Hex()

	t.Log("CreateContact Test PASSED. id : ", id)

}

func TestGetAllContacts(t *testing.T) {

	ctx := context.WithValue(context.Background(), "col", testContactsCollection)

	_, err := GetAllContacts(ctx)
	if err != nil {
		t.Error(`GetAllContacts Test FAILED
		err : `, err)
	}
	t.Logf("GetAllContacts Test PASSED.")

}

func TestUpdateContact(t *testing.T) {

	ctx := context.WithValue(context.Background(), "col", testContactsCollection)

	err := UpdateContact(id, contact.Name, contact.Number, ctx)
	if err != nil {
		t.Error(`UpdateContact Test FAILED!
		err: `, err)
	}
	t.Log("UpdateContact Test PASSED!")
}

func TestDeleteContact(t *testing.T) {

	ctx := context.WithValue(context.Background(), "col", testContactsCollection)

	err := DeleteContact(id, ctx)
	if err != nil {
		t.Error(`DeleteContact Test FAILED!
		err: `, err)
	}

	t.Log("DeleteContact Test PASSED.")

}

func TestDeleteAllContacts(t *testing.T) {

	ctx := context.WithValue(context.Background(), "col", testContactsCollection)

	_, err := DeleteAllContacts(ctx)
	if err != nil {
		t.Error(`DeleteAllContacts Test FAILED!
		err: `, err)
	}

	t.Log("DeleteAllContacts Test PASSED.")

}
