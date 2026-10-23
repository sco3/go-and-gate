package srv

import (
	"and-gate/v2/pb"
	"net"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Srv struct {
	pb.UnimplementedStartServiceServer
	server     *grpc.Server
	listener   net.Listener
	startNanos int64
	redis      *redis.Client
}
