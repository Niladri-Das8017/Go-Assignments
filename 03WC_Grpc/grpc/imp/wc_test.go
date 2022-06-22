package imp

import (
	"context"
	"log"
	"net"
	"testing"
	pb "wcgrpc/grpc/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

type BookDetails struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty"`
	Author string             `json:"author,omitempty"`
}

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterWordCountServieceServer(s, &WcServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestWc(t *testing.T) {

	//Dial a connection to grpc Server
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	//Create new Client
	c := pb.NewWordCountServieceClient(conn)

	//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	//Result
	response, err := c.WordCount(ctx, &pb.TextRequest{Text: "My name is Niladri, I am Niladri."})
	if err != nil {
		t.Fatal("Could not count word: \n", err)
	}
	t.Log("WordCount:\n")
	for _, value := range response.WcList {
		t.Logf(`"Word: %s	Count: %d`, value.Word, value.Count)
	}
}
