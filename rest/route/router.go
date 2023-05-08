package route

import (
	"github.com/gin-gonic/gin"
	"github/achjailani/go-simple-grpc/infrastructure/dependency"
	"github/achjailani/go-simple-grpc/rest/handler"
	"github/achjailani/go-simple-grpc/rest/middleware"
)

// Router is a struct contains dependencies needed
type Router struct {
	*dependency.Dependency
}

// NewRouter is a constructor will initialize Router.
func NewRouter(options ...Option) *Router {
	router := &Router{}

	for _, opt := range options {
		opt(router)
	}

	return router
}

// Init is a function
func (r *Router) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.Use(middleware.Logger())

	hand := handler.NewHandler(r.Dependency)

	httpLog := handler.NewRequestLogHandler(hand)
	hello := handler.NewHelloHandler(hand)

	e.GET("/api/ping", hello.Ping)
	e.POST("/api/hello", hello.SayHello)
	e.POST("/api/request-logs", httpLog.Create)

	return e
}
