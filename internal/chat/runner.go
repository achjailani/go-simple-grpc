package chat

import "github.com/gin-gonic/gin"

// ServeWSChat is a function
func ServeWSChat(c *gin.Context) {
	hub := newHub()
	go hub.run()

	serveWs(hub, c.Writer, c.Request)
}
