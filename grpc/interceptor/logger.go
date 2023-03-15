package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

// UnaryLoggerServerInterceptor is a method for logger server unary interceptor
func UnaryLoggerServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Printf("[%s]", info.FullMethod)

		return handler(ctx, req)
	}
}
