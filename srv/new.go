package srv

import (
	"and-gate/v2/pb"
	"net"

	"google.golang.org/grpc"
)

func NewServer(
	port string,

) (*Srv, error) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err

	}
	server := grpc.NewServer()
	srv := Srv{server: server, listener: listener}
	pb.RegisterStartServiceServer(server, &srv)
	return &srv, nil

}
