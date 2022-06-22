package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	pb "lms/grpc/pb"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

const address = ":50051"

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

//main function
func main() {

	//Dial a connection to grpc Server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Dial failed: ", err)
	}
	defer conn.Close()

	//Create new Client
	client := pb.NewLmsServiceClient(conn)

	//initialize context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	//Input option
	for {
		fmt.Println(`Welcome to the library...
	 
		**MENU**
		1. Upload a New Book
		2. All Books
		3. Search a Book
		4. Update Book
		5. Delete Book 
		Choose other number to exit!`)

		fmt.Printf("Choose Option: ")

		input, _ := reader.ReadString('\n')
		option, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
		if err != nil {
			log.Fatal("Failed to convert string into int")
		}

		switch option {
		case 1:
			createBook(client, ctx)
			continue
		case 2:
			listAllBooks(client, ctx)
			continue
		case 3:
			searchBooks(client, ctx)
			continue
		case 4:
			updateBook(client, ctx)
			continue
		case 5:
			deleteBook(client, ctx)
			continue
		default:
			os.Exit(0)
		}
	}

}

//Crud methods goes here...

//Create Book
func createBook(client pb.LmsServiceClient, ctx context.Context) {

	//Input
	fmt.Println("Insert Details: ")

	fmt.Print("Title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Wrong input: Title", err)
	}
	title = strings.TrimSpace(title)

	fmt.Print("Author: ")
	author, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Wrong input: Title", err)
	}
	author = strings.TrimSpace(author)

	if title == "" || author == "" {
		log.Fatal("Empty input")
	}

	//Creating book to send as request
	newBook := &pb.Book{
		Title:  title,
		Author: author,
	}
	//Call CreateBook that returns a book as response
	response, err := client.CreateBook(ctx, &pb.CreateBookReq{Book: newBook})
	if err != nil {
		log.Fatal("Could not Create Book: \n", err)
	}
	//print
	log.Printf(`New Book Uploaded:
	Book Id: %s
	Title: %s
	Author: %s`, response.Book.Id, response.Book.Title, response.Book.Author)
}

//All Books / Bi-Directional Streaming
func listAllBooks(client pb.LmsServiceClient, ctx context.Context) {

	fmt.Print("Page no: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	pageNo, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Fatal("Wrong input: Title", err)
	}

	fmt.Print("How many books you want in a page? ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	noOfItems, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Fatal("Wrong input: Title", err)
	}

	dontPrintUpto := (pageNo - 1) * noOfItems

	// Call ListAllBook that returns a stream
	stream, err := client.ListAllBooks(ctx, &pb.ListAllBooksReq{})
	// Check for errors
	if err != nil {
		log.Fatal("ListAllBooks did not return stream: ", err)
	}

	fmt.Println("List of Books: ")

	// Start iterating
	for i := 0; i < int(pageNo*noOfItems); i++ {

		// stream.Recv returns a pointer to a book in a current iteration
		response, err := stream.Recv()
		// If end of stream, break the loop
		if err == io.EOF {
			break
		}
		// if err, print error
		if err != nil {
			log.Fatal("Stream error: ", err)
		}

		// If everything went well use the generated getter to print the Book Details
		if i >= int(dontPrintUpto) {

			fmt.Printf("%d %v\n", i+1, response.GetBook())

		}

	}
}

//Search Books / Bi-Directional Streaming
func searchBooks(client pb.LmsServiceClient, ctx context.Context) {

	//Initializing stream
	var stream pb.LmsService_SearchBooksClient
	//Input
	fmt.Println(`Search Books from Library: 

	Chose an option...
	1. Search By Title
	2. Search By Author`)

	fmt.Printf("Your Decition : ")
	input, _ := reader.ReadString('\n')
	decition, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		log.Fatal("Failed to convert string into int")
	}

	switch decition {
	case 1:
		fmt.Print("Title: ")
		title, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Wrong input: Title", err)
		}
		title = strings.TrimSpace(title)

		// Call SearchBook that returns a stream
		stream, err = client.SearchBooks(ctx, &pb.SearchBooksReq{
			Search: &pb.SearchBooksReq_Title{Title: title},
		})
		// Check for errors
		if err != nil {
			fmt.Println("SearchBook did not return stream: ", err)
		}

	case 2:

		fmt.Print("Author: ")
		author, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Wrong input: Author", err)
		}
		author = strings.TrimSpace(author)

		// Call ListBook that returns a stream
		stream, err = client.SearchBooks(ctx, &pb.SearchBooksReq{
			Search: &pb.SearchBooksReq_Author{Author: author},
		})

		// Check for errors
		if err != nil {
			fmt.Println("SearchBook did not return stream: ", err)
		}

	default:
		fmt.Println("Invalid Option")
	}

	//Print Result
	fmt.Println("Books we have found...")

	// Start iterating
	for i := 0; i < 3; i++ {

		// stream.Recv returns a pointer to a book in a current iteration
		responseStream, err := stream.Recv()
		// If end of stream, break the loop
		if err == io.EOF {
			break
		}
		// if err, print error
		if err != nil {
			log.Fatal("Stream error: ", err)
		}

		// If everything went well use the generated getter to print the Book Details
		fmt.Println(responseStream.GetBook())

	}
}

//Update Book
func updateBook(client pb.LmsServiceClient, ctx context.Context) {

	//Input
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert Details to update: ")

	fmt.Print("Id: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Wrong input: Title", err)
	}
	id = strings.TrimSpace(id)

	fmt.Print("Title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Wrong input: Title", err)
	}
	title = strings.TrimSpace(title)

	fmt.Print("Author: ")
	author, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Wrong input: Title", err)
	}
	author = strings.TrimSpace(author)

	if id == "" || title == "" || author == "" {
		log.Fatal("Empty input")
	}

	//Creating Request
	updateBook := &pb.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}
	//Call UpdateBook that returns a Book as response
	response, err := client.UpdateBook(ctx, &pb.UpdateBookReq{Book: updateBook})
	if err != nil {
		log.Fatal("Could not update book: \n", err)
	}
	//print
	log.Printf(`Book Updated:
	Book Id: %s
	Title: %s
	Author: %s`, response.Book.Id, response.Book.Title, response.Book.Author)
}

//Delete Book
func deleteBook(client pb.LmsServiceClient, ctx context.Context) {
	//Input
	//fmt.Println("Input Any details to delete a book")
	fmt.Printf("Book Id: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	id := strings.TrimSpace(input)
	//Call DeleteBook
	_, err = client.DeleteBook(ctx, &pb.DeleteBookReq{Id: id})
	if err != nil {
		log.Fatal(err)
	}

	//Print Result
	fmt.Print("\nDeleted book with id: ", id)
}
