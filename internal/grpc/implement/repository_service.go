package impl

import (
	"context"

	pb "github.com/ssoda/grpc-playground/internal/grpc/proto"
)

// PlayServer struct
type PlayServer struct{}

// Ping 測試連線
func (p *PlayServer) Ping(ctx context.Context, req *pb.PingRequest) (res *pb.PongResponse, err error) {
	return &pb.PongResponse{
		Message: "Pong!",
	}, nil
}
