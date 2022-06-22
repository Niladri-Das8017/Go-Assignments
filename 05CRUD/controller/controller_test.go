package controller

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//test model contact
var id string

var testContactsCollection *mongo.Collection

const connectionString = "mongodb+srv://Niladri:12345@cluster0.wvsw2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

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

func TestSeverHome(t *testing.T) {

	//creating request
	req, err := http.NewRequest("GET", "localhost:4040/api/allcontacts", nil)
	if err != nil {
		t.Fatal("Could not create request for ServerHome.\n", err)
	}
	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()
	ServeHome(rec, req)

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
		ServerHome Passed:
		Status Code : "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}
	t.Log("\nTest Result : ", string(result))
}

func TestCreateCcontact(t *testing.T) {

	newContact := []byte(`
		{
			"name" : "testName",
			"number" : "1234567890"
		}
	`)

	//creating request
	req, err := http.NewRequest("POST", "localhost:4040/api/allcontacts", strings.NewReader(string(newContact)))
	if err != nil {
		t.Fatal("Could not create request for CreateContact.\n", err)
	}
	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()

	//passing test database using context
	ctx := context.WithValue(req.Context(), "col", testContactsCollection)

	CreateContact(rec, req.WithContext(ctx))

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
	CreateContact Passed:
	Status Code: "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}

	t.Log("\nTest Result : ", string(result))

	contactSlice := strings.Split(string(result), "\n")

	id = contactSlice[1]
	id = id[1 : len(id)-1]
	t.Log("id : ", id)

}

func TestGetAllContacts(t *testing.T) {

	//creating request
	req, err := http.NewRequest("GET", "localhost:4040/api/allcontacts", nil)
	if err != nil {
		t.Fatal("Could not create request for GetAllContacts.\n", err)
	}
	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()

	//passing test database using context
	ctx := context.WithValue(req.Context(), "col", testContactsCollection)

	GetAllContacts(rec, req.WithContext(ctx))

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
		GetAllContacts Passed:
		Status Code: "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}
	t.Log("\nTest Result : ", string(result))
}

func TestUpdateContact(t *testing.T) {
	updateContact := []byte(`
		{
			"name" : "updatedtestName",
			"number" : "9874563210"
		}
	`)

	req, err := http.NewRequest("PUT", "localhost:4040/api/updatecontact", strings.NewReader(string(updateContact)))
	if err != nil {
		t.Fatal("Could not create request for UpdateContact.\n", err)
	}
	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()

	//fake gorilla/mux var
	vars := map[string]string{
		"id": id,
	}

	// setting var to the req
	req = mux.SetURLVars(req, vars)

	//passing test database using context
	ctx := context.WithValue(req.Context(), "col", testContactsCollection)
	UpdateContact(rec, req.WithContext(ctx))

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
	UpdateContact Passed:
	Status Code: "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}

	t.Log("\nTest Result : ", string(result))

}

func TestDeleteContact(t *testing.T) {


	req, err := http.NewRequest("DELETE", "localhost:4040/api/deletecontact", nil)
	if err != nil {
		t.Fatal("Could not create request for DeleteContact.\n", err)
	}
	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()

	//fake gorilla/mux var
	vars := map[string]string{
		"id": id,
	}

	// setting var to the req
	req = mux.SetURLVars(req, vars)

	//passing test database using context
	ctx := context.WithValue(req.Context(), "col", testContactsCollection)
	DeleteContact(rec, req.WithContext(ctx))

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
	DeleteContact Passed:
	Status Code: "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}

	t.Log("\nTest Result : ", string(result))

}

func TestDeleteAll(t *testing.T) {

	req, err := http.NewRequest("DELETE", "localhost:4040/api/deleteallcontacts", nil)
	if err != nil {
		t.Fatal("Could not create request for DeleteAllContacts.\n", err)
	}
	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()

	//passing test database using context
	ctx := context.WithValue(req.Context(), "col", testContactsCollection)
	DeleteAllContacts(rec, req.WithContext(ctx))

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
	DeleteAllContacts Passed:
	Status Code: "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}

	t.Log("\nTest Result : ", string(result))
}
