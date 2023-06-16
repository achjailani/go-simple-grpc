package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github/achjailani/go-simple-grpc/infrastructure/dependency"
	"github/achjailani/go-simple-grpc/internal/chat"
	"github/achjailani/go-simple-grpc/rest/handler"
	"github/achjailani/go-simple-grpc/rest/middleware"
	"log"
	"net/http"
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
	e.Use(cors.Default())
	e.Use(middleware.Logger())

	hand := handler.NewHandler(r.Dependency)

	httpLog := handler.NewRequestLogHandler(hand)
	hello := handler.NewHelloHandler(hand)

	e.GET("/ws", chat.ServeWSChat)
	e.GET("/", serveChatTemplate)

	//e.GET("/ws", websocx.NewWebSocket().Handle)

	e.GET("/api/ping", hello.Ping)
	e.POST("/api/hello", hello.SayHello)
	e.POST("/api/request-logs", httpLog.Create)

	return e
}

// serveChatTemplate is a function to serve template
func serveChatTemplate(c *gin.Context) {
	log.Println(c.Request.URL)

	if c.Request.URL.Path != "/" {
		http.Error(c.Writer, "Not found", http.StatusNotFound)
		return
	}
	if c.Request.Method != http.MethodGet {
		http.Error(c.Writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(c.Writer, c.Request, "template/chat.html")
}
