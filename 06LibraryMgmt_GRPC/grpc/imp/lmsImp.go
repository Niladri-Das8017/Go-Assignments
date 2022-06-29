package imp

import (
	"context"
	"errors"
	"fmt"
	"lms/database"
	pb "lms/grpc/pb"
	"lms/model"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Implementation of proto interfaces
type LmsServiceServer struct {
	pb.UnimplementedLmsServiceServer
}

//Create Book
func (s *LmsServiceServer) CreateBook(ctx context.Context, req *pb.CreateBookReq) (*pb.CreateBookRes, error) {

	book := req.Book
	//Request Validation
	if book.Title == "" || book.Author == "" {
		return nil, errors.New("Invalid Input")
	}

	//convert message CreateBookRequest into a BookDetails type to convert into BSON
	data := model.BookDetails{
		Title:  book.Title,
		Author: book.Author,
	}

	//Insert Data
	//Call CreateBook
	id, err := database.CreateBook(ctx, data)
	if err != nil {
		return nil, err
	}

	book.Id = id

	return &pb.CreateBookRes{Book: book}, nil

}

//List All Books
func (s *LmsServiceServer) ListAllBooks(req *pb.ListAllBooksReq, stream pb.LmsService_ListAllBooksServer) error {

	//data contains structure of book
	data := &model.BookDetails{}

	//Initializing Cursor by Calling database interactor
	cursor, err := database.ListAllBooks()
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())

	// Looping through Database
	for cursor.Next(context.Background()) {

		//whenever we decode, we pass on a referrence like, "if you decode use my structure to decode that"
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		// If no error is found, send book over stream
		err = stream.Send(&pb.ListAllBooksRes{
			Book: &pb.Book{
				Id:     data.ID.Hex(),
				Title:  data.Title,
				Author: data.Author,
			},
		})

		//Check if client still connected
		if err != nil {
			log.Fatal("Client Disconnected : ", err) //Client Diisconnected thatts why fatal error
		}

	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}

	return nil
}

//Search Book
func (s *LmsServiceServer) SearchBooks(req *pb.SearchBooksReq, stream pb.LmsService_SearchBooksServer) error {

	//data contains structure of book
	data := &model.BookDetails{}

	title := req.GetTitle()
	author := req.GetAuthor()

	//Request validation
	if title == "" && author == "" {
		return errors.New("Empty Search Key")
	}

	//Finding book and initializing cursor
	cursor, err := database.SearchBooks(title, author)
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	// Looping through Database
	for cursor.Next(context.Background()) {

		//whenever we decode, we pass on a referrence like, "if you decode use my structure to decode that"
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		// If no error is found, send book over stream
		err = stream.Send(&pb.SearchBooksRes{
			Book: &pb.Book{
				Id:     data.ID.Hex(),
				Title:  data.Title,
				Author: data.Author,
			},
		})

		//Check if client still connected
		if err != nil {
			log.Fatal("Client Disconnected : ", err) //Client Diisconnected thatts why fatal error
		}

	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

//Update Book
func (s *LmsServiceServer) UpdateBook(ctx context.Context, req *pb.UpdateBookReq) (*pb.UpdateBookRes, error) {

	updateReq := req.GetBook()
	//UpdateReq validation
	if updateReq.GetId() == "" || updateReq.GetTitle() == "" || updateReq.GetAuthor() == "" {
		return nil, errors.New("Invalid Update Request")
	}

	//Call Interactor UpdateBook
	updatedBook, err := database.UpdateBook(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateBookRes{Book: &pb.Book{Id: updatedBook.ID.Hex(), Title: updatedBook.Title, Author: updatedBook.Author}}, nil
}

//Delete Book
func (s *LmsServiceServer) DeleteBook(ctx context.Context, req *pb.DeleteBookReq) (*pb.DeleteBookRes, error) {

	//DeleteReq Validation
	if req.GetId() == "" {
		return nil, errors.New("Invald Delete Reequest")

	}

	//Call Database Interactor DeeleteBook
	err := database.DeleteBook(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.DeleteBookRes{Success: true}, nil
}
