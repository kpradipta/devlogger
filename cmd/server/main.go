package main

import (
	"log"
	"net"

	pb "github.com/devlogger/kpradipta/api/proto/logger"
	"github.com/devlogger/kpradipta/internal/db"
	"github.com/devlogger/kpradipta/internal/handler"
	"github.com/devlogger/kpradipta/internal/metrics"
	"google.golang.org/grpc"
)

func main() {
	metrics.Init()
	db.Init()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLogServiceServer(s, &handler.LogServer{})

	log.Println("gRPC server running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
