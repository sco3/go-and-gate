package srv

import (
	"and-gate/v2/pb"
	"net"

	"google.golang.org/grpc"
)

type Srv struct {
	pb.UnimplementedStartServiceServer
	server   *grpc.Server
	listener net.Listener
}
