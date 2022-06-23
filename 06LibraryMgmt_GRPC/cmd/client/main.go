package main

import (
	"bufio"
	"context"
	"fmt"
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

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Failed to take Input : ", err)
		}
		option, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
		if err != nil {
			log.Fatal("Failed to convert string into int")
		}

		switch option {
		case 1:

			err := createBook(client, ctx)
			if err != nil {
				fmt.Println(err)
			}

			continue
		case 2:

			err := listAllBooks(client, ctx)
			if err != nil {
				fmt.Println(err)
			}

			continue
		case 3:

			err := searchBooks(client, ctx)
			if err != nil {
				fmt.Println(err)
			}

			continue
		case 4:

			err := updateBook(client, ctx)
			if err != nil {
				fmt.Println(err)
			}

			continue
		case 5:

			err := deleteBook(client, ctx)
			if err != nil {
				fmt.Println(err)
			}

			continue
		default:
			os.Exit(0)
		}
	}

}
