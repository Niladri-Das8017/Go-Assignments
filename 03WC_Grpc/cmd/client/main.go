package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
	pb "wcgrpc/grpc/pb"

	"google.golang.org/grpc"
)

const address = ":50051"

func main() {
	//Input
	fmt.Print("Input String: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	//Error Handling
	if err != nil {
		panic(err)
	}

	//Dial a connection to grpc Server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Dial failed: ", err)
	}
	defer conn.Close()

	//Create new Client
	c := pb.NewWordCountServieceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Result
	response, err := c.WordCount(ctx, &pb.TextRequest{Text: text})
	if err != nil {
		log.Fatal("Could not count word: \n", err)
	}
	//priint
	log.Println("WordCount:")
	for _, value := range response.WcList {
		log.Printf(`	Word: %s	Count: %d`, value.Word, value.Count)
	}

}
