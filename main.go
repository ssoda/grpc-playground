package main

import (
	pb "github.com/ssoda/grpc-playground/internal/grpc/proto"
	impl "github.com/ssoda/grpc-playground/internal/grpc/implement"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9487")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := &impl.PlayServer{}
	pb.RegisterPlayServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
