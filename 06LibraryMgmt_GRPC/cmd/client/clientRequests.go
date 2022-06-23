package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	pb "lms/grpc/pb"
	"log"
	"strconv"
	"strings"
)

//Create Book
func createBook(client pb.LmsServiceClient, ctx context.Context) error {

	//Input
	fmt.Println("Insert Details: ")

	fmt.Print("Title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return errors.New(fmt.Sprint("Wrong input: Title", err))
	}
	title = strings.TrimSpace(title)

	fmt.Print("Author: ")
	author, err := reader.ReadString('\n')
	if err != nil {
		return errors.New(fmt.Sprint("Wrong input: Title", err))
	}
	author = strings.TrimSpace(author)

	if title == "" || author == "" {
		return errors.New(("Empty input"))
	}

	//Creating book to send as request
	newBook := &pb.Book{
		Title:  title,
		Author: author,
	}
	//Call CreateBook that returns a book as response
	response, err := client.CreateBook(ctx, &pb.CreateBookReq{Book: newBook})
	if err != nil {
		return errors.New(fmt.Sprint("Could not Create Book: \n", err))
	}
	//print
	log.Printf(`New Book Uploaded:
	Book Id: %s
	Title: %s
	Author: %s`, response.Book.Id, response.Book.Title, response.Book.Author)

	return nil
}

//All Books / Bi-Directional Streaming
func listAllBooks(client pb.LmsServiceClient, ctx context.Context) error {

	fmt.Print("Page no: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	input = strings.TrimSpace(input)
	pageNo, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return errors.New("Wrong input: Page No")
	}

	fmt.Print("How many books you want in a page? ")
	input, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	input = strings.TrimSpace(input)
	noOfItems, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return errors.New("Wrong input: No. of Items")
	}

	dontPrintUpto := (pageNo - 1) * noOfItems

	// Call ListAllBook that returns a stream
	stream, err := client.ListAllBooks(ctx, &pb.ListAllBooksReq{})
	// Check for errors
	if err != nil {
		return errors.New(fmt.Sprint("ListAllBooks did not return stream: ", err))
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
			return errors.New(fmt.Sprint("Stream error: ", err))
		}

		// If everything went well use the generated getter to print the Book Details
		if i >= int(dontPrintUpto) {

			fmt.Printf("%d %v\n", i+1, response.GetBook())

		}

	}

	return nil
}

//Search Books / Bi-Directional Streaming
func searchBooks(client pb.LmsServiceClient, ctx context.Context) error {

	//Initializing stream
	var stream pb.LmsService_SearchBooksClient
	//Input
	fmt.Println(`Search Books from Library: 

	Chose an option...
	1. Search By Title
	2. Search By Author`)

	fmt.Printf("Your Decition : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	decition, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		return errors.New(fmt.Sprint("Failed to convert string into int"))
	}

	switch decition {
	case 1:
		fmt.Print("Title: ")
		title, err := reader.ReadString('\n')
		if err != nil {
			return errors.New(fmt.Sprint("Wrong input: Title", err))
		}
		title = strings.TrimSpace(title)

		if title == "" {
			return errors.New("Empty Title")
		}

		// Call SearchBook that returns a stream
		stream, err = client.SearchBooks(ctx, &pb.SearchBooksReq{
			Search: &pb.SearchBooksReq_Title{Title: title},
		})
		// Check for errors
		if err != nil {
			return errors.New(fmt.Sprint("SearchBook did not return stream: ", err))
		}

	case 2:

		fmt.Print("Author: ")
		author, err := reader.ReadString('\n')
		if err != nil {
			return errors.New(fmt.Sprint("Wrong input: Author", err))
		}
		author = strings.TrimSpace(author)

		if author == "" {
			return errors.New("Empty Author")
		}

		// Call ListBook that returns a stream
		stream, err = client.SearchBooks(ctx, &pb.SearchBooksReq{
			Search: &pb.SearchBooksReq_Author{Author: author},
		})

		// Check for errors
		if err != nil {
			return errors.New(fmt.Sprint("SearchBook did not return stream: ", err))
		}

	default:
		return errors.New("Invalid Option")
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
			return errors.New(fmt.Sprint("Stream error: ", err))
		}

		// If everything went well use the generated getter to print the Book Details
		fmt.Println(responseStream.GetBook())

	}
	return nil
}

//Update Book
func updateBook(client pb.LmsServiceClient, ctx context.Context) error {

	//Input
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert Details to update: ")

	fmt.Print("Id: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		return errors.New(fmt.Sprint("Wrong input: ID", err))
	}
	id = strings.TrimSpace(id)

	fmt.Print("Title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return errors.New(fmt.Sprint("Wrong input: Title", err))
	}
	title = strings.TrimSpace(title)

	fmt.Print("Author: ")
	author, err := reader.ReadString('\n')
	if err != nil {
		return errors.New(fmt.Sprint("Wrong input: Author", err))
	}
	author = strings.TrimSpace(author)

	if id == "" || title == "" || author == "" {
		return errors.New(fmt.Sprint("Empty input"))
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
		return errors.New(fmt.Sprint("Could not update book: \n", err))
	}
	//print
	log.Printf(`Book Updated:
	Book Id: %s
	Title: %s
	Author: %s`, response.Book.Id, response.Book.Title, response.Book.Author)

	return nil
}

//Delete Book
func deleteBook(client pb.LmsServiceClient, ctx context.Context) error {
	//Input
	//fmt.Println("Input Any details to delete a book")
	fmt.Printf("Book Id: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	id := strings.TrimSpace(input)
	//Call DeleteBook
	_, err = client.DeleteBook(ctx, &pb.DeleteBookReq{Id: id})
	if err != nil {
		return err
	}

	//Print Result
	fmt.Print("\nDeleted book with id: ", id)
	return nil
}
