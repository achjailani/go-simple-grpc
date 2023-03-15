package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

// UnaryLoggerServerInterceptor is a function for logger server unary interceptor
func UnaryLoggerServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Printf("--> Unary interceptor: [%s]", info.FullMethod)

		return handler(ctx, req)
	}
}

// StreamLoggerServerInterceptor is a function for logger server stream interceptor
func StreamLoggerServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Printf("--> Stream interceptor: [%s]", info.FullMethod)

		return handler(srv, stream)
	}
}
