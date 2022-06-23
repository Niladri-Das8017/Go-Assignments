package imp

import (
	"context"
	"lms/database"
	pb "lms/grpc/pb"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	lis = bufconn.Listen(bufSize)
	server := grpc.NewServer()
	pb.RegisterLmsServiceServer(server, &LmsServiceServer{})
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestLmsImp(t *testing.T) {

	//Dial a connection to grpc Server
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	//Create new Client
	client := pb.NewLmsServiceClient(conn)

	//Test CreateBook
	newBook := &pb.Book{
		Title:  "test title",
		Author: "test author",
	}

	response, err := client.CreateBook(ctx, &pb.CreateBookReq{Book: newBook})
	if err != nil {
		t.Error("Test CreateBook FAILED!\nerr: ", err)
	}
	t.Log("Test CreateBook PASSED.")

	//get id
	id := response.Book.Id

	_, err = client.ListAllBooks(ctx, &pb.ListAllBooksReq{})
	if err != nil {
		t.Error("Test ListAllBooks FAILED!\nerr: ", err)
	}
	t.Log("Test ListAllBooks PASSED.")

	//Test SearchBooks
	_, err = client.SearchBooks(ctx, &pb.SearchBooksReq{
		Search: &pb.SearchBooksReq_Title{Title: newBook.Title},
	})
	if err != nil {
		t.Error("Test SearchBooks FAILED!\nerr: ", err)
	}

	t.Log("Test SearchBooks PASSED.")

	//Test Updatebook
	updateBook := &pb.Book{
		Id:     id,
		Title:  "updated test title",
		Author: "updated test title",
	}
	_, err = client.UpdateBook(ctx, &pb.UpdateBookReq{Book: updateBook})
	if err != nil {
		t.Error("Test UpdateBooks FAILED!\nerr: ", err)
	}

	t.Log("Test UpdateBooks PASSED.")

	//Test DeleteBook
	_, err = client.DeleteBook(ctx, &pb.DeleteBookReq{Id: id})
	if err != nil {
		t.Error("Test DeleteBooks FAILED!\nerr: ", err)
	}

	t.Log("Test DeleteBooks PASSED.")

}
