package websocx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	writeBufferSize = 1024
	readBufferSize  = 1024
)

// WebSocket is a struct
type WebSocket struct {
}

// NewWebSocket is a constructor
func NewWebSocket() *WebSocket {
	return &WebSocket{}
}

// Handle is a method
func (ws *WebSocket) Handle(c *gin.Context) {
	// Allow connections from all origins (CORS)
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

	// Allow connections from all origins (CORS)
	responseHeaders := http.Header{}
	responseHeaders.Set("Access-Control-Allow-Origin", "*")

	upgrader := websocket.Upgrader{
		HandshakeTimeout: 10 * time.Second,
		ReadBufferSize:   readBufferSize,
		WriteBufferSize:  writeBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, responseHeaders)
	if err != nil {
		log.Printf("Failed to upgrade connection to WebSocket: %v\n", err)
		return
	}

	defer conn.Close()

	// Handle WebSocket connection
	for {
		// Read message from client
		messageType, message, er := conn.ReadMessage()
		if er != nil {
			log.Println("Error reading message from WebSocket:", er)
			break
		}

		if string(message) == "ping" {
			message = []byte("PONG!!")
		}

		log.Printf("RECIEVED: %s, message type: %d\n", message, messageType)

		// Write response back to the client
		err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("SVR: %s", string(message))))
		if err != nil {
			log.Println("Error writing message to WebSocket:", err)
			break
		}
	}
}
