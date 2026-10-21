package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)
import "and-gate/v2/pb"

type srv struct {
	pb.UnimplementedStartServiceServer
}

func (s *srv) GetStartNanos(
	ctx context.Context,
	req *pb.GetStartNanosRequest,
) (*pb.GetStartNanosResponse, error) {
	return &pb.GetStartNanosResponse{StartNanos: 0}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatal("Cannot start listener")
	}
	server := grpc.NewServer()
	pb.RegisterStartServiceServer(server, &srv{})
	log.Printf("serve")
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
