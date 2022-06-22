package database

import (
	"context"
	"errors"
	"fmt"
	pb "lms/grpc/pb"
	"lms/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateBook(ctx context.Context, newBook model.BookDetails) (string, error) {
	//Insert Data
	inserted, err := Books.InsertOne(ctx, newBook)
	if err != nil {
		return "", status.Errorf(codes.Unknown, fmt.Sprintf("Insert operation FAILED! \n %v", err))
	}

	//id assigned by mongodb
	id := inserted.InsertedID.(primitive.ObjectID)

	return id.Hex(), nil

}

func ListAllBooks() (*mongo.Cursor, error) {

	cursor, err := Books.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

func SearchBooks(title string, author string) (*mongo.Cursor, error) {

	var cursor *mongo.Cursor
	var err error

	//Validation
	if title == "" && author == "" {
		return nil, errors.New("Empty Search Key")
	}

	if title != "" {
		cursor, err = Books.Find(context.Background(), bson.M{"title": title})
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Book not found, Check title: %v", err))
		}
	}

	if author != "" {
		cursor, err = Books.Find(context.Background(), bson.M{"author": author})
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Book not found, Check author: %v", err))
		}

	}

	return cursor, nil
}

func UpdateBook(ctx context.Context, book *pb.Book) (*model.BookDetails, error) {

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(book.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert the supplied book id to a MongoDB ObjectId: %v", err))
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Convert the data to be updated into an unordered Bson document
	updates := bson.M{"$set": bson.M{"title": book.GetTitle(), "author": book.GetAuthor()}}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := Books.FindOneAndUpdate(ctx, filter, updates, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'updatedBook'
	updatedBook := model.BookDetails{}
	err = result.Decode(&updatedBook)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find book with supplied ID: %v", err),
		)
	}

	return &updatedBook, nil

}


func DeleteBook(id string) error {
	
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	filter := bson.M{"_id": oid}
	_, err = Books.DeleteOne(context.Background(), filter)
	if err != nil {
		return  status.Errorf(codes.NotFound, fmt.Sprintf(`Delete operation FAILED!
		Could not find Book with id %s: %v`, id, err))
	}

	return nil
}