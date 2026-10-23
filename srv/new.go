package srv

import (
	"and-gate/v2/pb"
	"net"

	"google.golang.org/grpc"
)

func NewServer(
	address string,
	host string,

) (*Srv, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err

	}
	server := grpc.NewServer()
	redisClient := NewRedisClient()
	srv := Srv{
		server: server, listener: listener, redis: redisClient,
		host: host,
	}
	pb.RegisterStartServiceServer(server, &srv)
	return &srv, nil

}
