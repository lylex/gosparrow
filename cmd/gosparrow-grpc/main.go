package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "gosparrow/pkg/gosparrow"
)

// TODO make the const in the somewhere else
const (
	port = "8081"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// GetName implements gosparrow.GetName
func (s *server) GetName(ctx context.Context, in *pb.GetNameReq) (*pb.GetNameResp, error) {
	return &pb.GetNameResp{Name: in.Prefix + "gosparrow"}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGosparrowServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
