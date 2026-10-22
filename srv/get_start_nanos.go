package srv

import (
	"and-gate/v2/pb"
	"context"
)

func (s *Srv) GetStartNanos(
	_ context.Context,
	_ *pb.GetStartNanosRequest,
) (*pb.GetStartNanosResponse, error) {
	return &pb.GetStartNanosResponse{StartNanos: 314}, nil
}
