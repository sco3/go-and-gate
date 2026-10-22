package srv

import (
	"and-gate/v2/pb"
	"context"
)

func (s *Srv) StopServer(
	_ context.Context,
	_ *pb.StopServerRequest,
) (*pb.StopServerResponse, error) {

	go func() {
		s.server.GracefulStop()
	}()
	return &pb.StopServerResponse{}, nil

}
