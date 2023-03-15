package interceptor

import (
	"context"
	"encoding/base64"
	"fmt"
	"github/achjailani/go-simple-grpc/grpc/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

// UnaryAuthClientInterceptor is a function to attach token.
func UnaryAuthClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		var username string
		var password string

		// check if method is not protected
		if !contract.ProtectedMethods()[method] {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		if v, exist := os.LookupEnv("FAKE_USERNAME"); exist {
			username = v
		}

		if v, exist := os.LookupEnv("FAKE_PASSWORD"); exist {
			password = v
		}

		token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))

		return invoker(func(t string) context.Context {
			return metadata.AppendToOutgoingContext(ctx, "authorization", t)
		}(token), method, req, reply, cc, opts...)
	}
}

// UnaryAuthServerInterceptor is a function to handle authorization
func UnaryAuthServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// check if method is not protected
		if !contract.ProtectedMethods()[info.FullMethod] {
			return handler(ctx, req)
		}

		m, valid := metadata.FromIncomingContext(ctx)
		if !valid {
			return nil, status.Error(codes.Unauthenticated, "no metadata provided")
		}

		tokenAuth, exist := m["authorization"]
		if !exist {
			return status.Error(codes.Unauthenticated, "no token provided"), nil
		}

		log.Printf("AUTH TOKEN: %s\n", tokenAuth)

		return handler(ctx, req)
	}
}
