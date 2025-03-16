package infra

import (
	"context"
	"fmt"

	"github.com/The-Fox-Hunt/auth/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "failed to get metadata")
	}

	username := md.Get("username")

	if len(username) == 0 {
		return nil, fmt.Errorf("username is empty")
	}

	ctx = context.WithValue(ctx, model.Username, username[0])

	return handler(ctx, req)
}
