package main

import (
	"github.com/The-Fox-Hunt/auth/internal/rpc"
	"github.com/The-Fox-Hunt/auth/pkg/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	service := rpc.New()

	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, service)

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)
}
