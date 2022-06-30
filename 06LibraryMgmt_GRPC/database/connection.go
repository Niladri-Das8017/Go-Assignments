package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Niladri:12345@cluster0.wvsw2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

//Initiialize ContactCollection
var Books *mongo.Collection

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
	library := client.Database("library")
	Books = library.Collection("books")
	fmt.Println("Database Connected")

}
