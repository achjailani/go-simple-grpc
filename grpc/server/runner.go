package server

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// RunGRPCServer is a function to run grpc server
func RunGRPCServer(s *grpc.Server, port int) error {
	lis, errListen := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if errListen != nil {
		return fmt.Errorf("failed to listen: %v", errListen)
	}

	serverError := make(chan error, 1)

	go func() {
		log.Printf("GRPC running on port %d\n", port)
		serverError <- s.Serve(lis)
	}()

	shutdownListenerChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownListenerChannel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverError:
		if err != nil {
			log.Printf("GRPC server can not started, %v\n", err)
			return err
		}
	case sig := <-shutdownListenerChannel:
		log.Printf("GRPC shutdown by signal %v\n", sig)
		s.GracefulStop()
		log.Println("GRPC stopped gracefully")
	}

	return nil
}
