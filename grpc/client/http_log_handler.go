package client

import (
	"context"
	"fmt"
	"github/achjailani/go-simple-grpc/proto/foo"
)

// SaveHttpLog is a method
func (r GRPCClient) SaveHttpLog(ctx context.Context, payloads []*foo.SaveHttpLogRequest) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stream, err := r.httpLog.SaveStreamHttpLog(ctx)
	if err != nil {
		return fmt.Errorf("create stream, %w", err)
	}

	for _, row := range payloads {
		if errSend := stream.Send(row); errSend != nil {
			return fmt.Errorf("send stream: %w", errSend)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("close & receive %w", err)
	}

	fmt.Println("Response: ")
	fmt.Println(resp)

	return nil
}
