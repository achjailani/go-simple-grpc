package route

import (
	"github.com/gin-gonic/gin"
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/domain/service"
	"github/achjailani/go-simple-grpc/grpc/client"
	"github/achjailani/go-simple-grpc/pkg/logger"
	"github/achjailani/go-simple-grpc/rest/handler"
	"github/achjailani/go-simple-grpc/rest/middleware"
)

// WithConfig is function
func WithConfig(config *config.Config) RouterOption {
	return func(r *Router) {
		r.config = config
	}
}

// WithRepository is function
func WithRepository(repo *service.Repositories) RouterOption {
	return func(r *Router) {
		r.repo = repo
	}
}

// WithGRPCClient is function
func WithGRPCClient(gClient *client.GRPCClient) RouterOption {
	return func(r *Router) {
		r.client = gClient
	}
}

// WithLogger is a function
func WithLogger(loggr *logger.Logger) RouterOption {
	return func(r *Router) {
		r.logger = loggr
	}
}

// Init is a function
func (r *Router) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.Use(middleware.Logger())

	hand := handler.NewHandler(r.repo, r.client, r.logger)

	httpLog := handler.NewRequestLogHandler(hand)
	hello := handler.NewHelloHandler(hand)

	e.GET("/api/ping", hello.Ping)
	e.POST("/api/hello", hello.SayHello)
	e.POST("/api/request-logs", httpLog.Create)

	return e
}
