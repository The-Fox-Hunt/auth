package main

import (
	"log"
	"net"

	"github.com/The-Fox-Hunt/auth/internal/repository/pg"
	"github.com/The-Fox-Hunt/auth/internal/rpc"
	"github.com/The-Fox-Hunt/auth/pkg/auth"
	"google.golang.org/grpc"
)

func main() {
	repo := pg.NewRepository()

	service := rpc.New(repo)

	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, service)

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)
}
