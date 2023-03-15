package client

import (
	"context"
	"github/achjailani/go-simple-grpc/proto/foo"
)

// SayHello is a method
func (r GRPCClient) SayHello(ctx context.Context, text string) (*foo.HelloReply, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	reply, err := r.hello.SayHello(ctx, &foo.HelloRequest{Text: text})
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Ping is a method
func (r GRPCClient) Ping(ctx context.Context) (*foo.PingReply, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ping, err := r.hello.Ping(ctx, &foo.PingRequest{})
	if err != nil {
		return nil, err
	}

	return ping, nil
}
