package interceptor

import (
	"context"
	"encoding/base64"
	"fmt"
	"github/achjailani/go-simple-grpc/grpc/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
)

// UnaryAuthClientInterceptor is a function to attach token.
func UnaryAuthClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// check if method is not protected
		if !contract.ProtectedMethods()[method] {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		return invoker(func(t string) context.Context {
			return metadata.AppendToOutgoingContext(ctx, "authorization", t)
		}(clientAttachToken()), method, req, reply, cc, opts...)
	}
}

// StreamAuthClientInterceptor is a function to attach token for stream interceptor
func StreamAuthClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		if !contract.ProtectedMethods()[method] {
			return streamer(ctx, desc, cc, method, opts...)
		}

		return streamer(func(t string) context.Context {
			return metadata.AppendToOutgoingContext(ctx, "authorization", t)
		}(clientAttachToken()), desc, cc, method, opts...)
	}
}

// clientAttachToken is a private function to attach token
func clientAttachToken() string {
	var username string
	var password string

	if v, exist := os.LookupEnv("FAKE_USERNAME"); exist {
		username = v
	}

	if v, exist := os.LookupEnv("FAKE_PASSWORD"); exist {
		password = v
	}

	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
}
