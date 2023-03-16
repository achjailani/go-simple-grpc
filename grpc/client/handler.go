package client

import (
	"context"
	"flag"
	"fmt"
	"github/achjailani/go-simple-grpc/proto/foo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func Run() error {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Cannot connect to server :%v", err)
	}

	defer func(conn *grpc.ClientConn) {
		errClose := conn.Close()
		if errClose != nil {
			log.Fatalf("err close, %v", errClose)
		}
	}(conn)

	gClient := NewGRPCClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	md := metadata.Pairs("authorization", "hell")
	ctx = metadata.NewOutgoingContext(ctx, md)

	defer cancel()

	payloads := []*foo.SaveHttpLogRequest{
		{Ip: "1.1.1.1", Path: "/user", Method: "POST"},
		{Ip: "1.1.1.2", Path: "/user/1", Method: "GET"},
		{Ip: "1.1.1.3", Path: "/user/2", Method: "GET"},
		{Ip: "1.1.1.4", Path: "/user/3", Method: "GET"},
		{Ip: "1.1.1.5", Path: "/user", Method: "GET"},
		{Ip: "1.1.1.6", Path: "/user", Method: "GET"},
		{Ip: "1.1.1.7", Path: "/user", Method: "POST"},
		{Ip: "1.1.1.8", Path: "/user", Method: "POST"},
		{Ip: "1.1.1.9", Path: "/user", Method: "POST"},
		{Ip: "1.1.1.10", Path: "/user", Method: "POST"},
	}

	err = gClient.SaveHttpLog(ctx, payloads)
	if err != nil {
		return fmt.Errorf("Could not call: %v", err)
	}

	return nil
}
