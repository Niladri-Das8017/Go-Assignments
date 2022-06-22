package main

import (
	"fmt"
	"lms/database"
	"lms/grpc/imp"
	pb "lms/grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":50051"

func main() {

	fmt.Println("\nWelcome to LMS GRPC")

	//Creatiung Database Connection
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen0: ", err)
	}

	//Initialize new Server
	server := grpc.NewServer()

	//Regester the server as a new grpc service
	pb.RegisterLmsServiceServer(server, &imp.LmsServiceServer{})
	log.Println("server listening at ", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}

}
