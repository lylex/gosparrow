package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "gosparrow/pkg/gosparrow"
)

const (
	address = "localhost:8081"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGosparrowClient(conn)

	// Contact the server and print out its response.
	r, err := c.GetName(context.Background(), &pb.GetNameReq{Prefix: "client-test-"})
	if err != nil {
		log.Fatalf("could not get the resp: %v", err)
	}
	log.Printf("The name is: %s", r.Name)
}
