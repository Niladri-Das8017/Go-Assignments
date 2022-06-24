package database

import (
	"context"
	"reflect"

	pb "lms/grpc/pb"
	"lms/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDb(t *testing.T) {

	ctx := context.Background()

	book := model.BookDetails{
		Title:  "test title",
		Author: "test author",
	}

	//Initiialize and Test Database
	err := InitDB()
	assert.Nil(t, err)

	//Test CreateBook
	id, err := CreateBook(ctx, book)
	assert.Nil(t, err)
	//Is id a String?
	if reflect.ValueOf(id).Kind() != reflect.String {
		t.Errorf("Id is not type string!!")
	}

	//Test ListAllBooks
	cursor, err := ListAllBooks()
	assert.Nil(t, err)
	assert.IsType(t, &mongo.Cursor{}, cursor)

	//Test SeearchBooks
	cursor, err = SearchBooks(book.Title, book.Author)
	assert.Nil(t, err)
	assert.IsType(t, &mongo.Cursor{}, cursor)

	//Test UpdateBook
	updateBoookDetails := &pb.Book{
		Id:     id,
		Title:  "Updated Title",
		Author: "Updated Author",
	}

	updatedBook, err := UpdateBook(ctx, updateBoookDetails)
	assert.Nil(t, err)
	assert.IsType(t, &model.BookDetails{}, updatedBook)

	//Test DeleteBook
	err = DeleteBook(id)
	assert.Nil(t, err)

}
