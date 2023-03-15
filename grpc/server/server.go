package server

import (
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/domain/service"
	"github/achjailani/go-simple-grpc/grpc/handler"
	"github/achjailani/go-simple-grpc/grpc/interceptor"
	"github/achjailani/go-simple-grpc/proto/foo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server is struct to hold any dependencies used for server
type Server struct {
	config *config.Config
	repo   *service.Repositories
}

// NewGRPCServer is constructor
func NewGRPCServer(conf *config.Config, repo *service.Repositories) *Server {
	return &Server{
		config: conf,
		repo:   repo,
	}
}

// Run is a method gRPC server
func (s *Server) Run(port int) error {
	//server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.AuthorizationInterceptor))
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryLoggerServerInterceptor(),
			interceptor.UnaryAuthServerInterceptor(),
		))

	handlers := handler.NewHandler(s.config, s.repo)

	// register service server
	foo.RegisterUserServiceServer(server, handlers)
	foo.RegisterHelloServer(server, handlers)
	foo.RegisterAuthServer(server, handlers)
	foo.RegisterLogServiceServer(server, handlers)

	// register reflection
	reflection.Register(server)

	return RunGRPCServer(server, port)
}
