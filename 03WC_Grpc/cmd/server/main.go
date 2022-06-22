package main

import (
	"log"
	"net"
	"wcgrpc/grpc/imp"
	pb "wcgrpc/grpc/pb"

	"google.golang.org/grpc"
)

const port = ":50051"

//main function
func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen0: ", err)
	}

	//Initialize new Server
	s := grpc.NewServer()

	//Regester the server as a new grpc service
	pb.RegisterWordCountServieceServer(s, &imp.WcServer{})
	log.Println("server listening at ", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}

}
