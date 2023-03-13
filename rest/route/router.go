package route

import (
	"github.com/gin-gonic/gin"
	"github/achjailani/go-simple-grpc/client"
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/domain/service"
	"github/achjailani/go-simple-grpc/rest/handler"
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

// Init is a function
func (r *Router) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()

	hand := handler.NewHandler(r.repo, r.client)

	httpLog := handler.NewRequestLogHandler(hand)

	e.POST("/api/request-logs", httpLog.Create)

	return e
}
