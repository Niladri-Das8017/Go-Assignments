package database

import (
	"context"
	//pb "lms/grpc/pb"
	"lms/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDbOT(t *testing.T) {

	ctx := context.Background()
	newBook := model.BookDetails{
		Title:  "test title",
		Author: "test author",
	}

	//Test CreateBook
	id, err := CreateBook(ctx, newBook)
	assert.Nil(t, err)
	t.Log(id)
}
