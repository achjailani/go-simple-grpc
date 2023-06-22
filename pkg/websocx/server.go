package websocx

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

// WebSocketServer is a struct
type WebSocketServer struct {
	upgrader websocket.Upgrader
	mutex    sync.Mutex
	clients  []*Client
}

// NewWebSocketServer is a constructor
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		upgrader: websocket.Upgrader{
			HandshakeTimeout: 10 * time.Second,
			ReadBufferSize:   readBufferSize,
			WriteBufferSize:  writeBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		mutex:   sync.Mutex{},
		clients: make([]*Client, 0),
	}
}

// Serve is a method to upgrade to websocket protocol
func (ws *WebSocketServer) Serve(writer http.ResponseWriter, request *http.Request) {
	// Allow connections from all origins (CORS)
	responseHeaders := http.Header{}
	responseHeaders.Set("Access-Control-Allow-Origin", "*")

	conn, err := ws.upgrader.Upgrade(writer, request, responseHeaders)
	if err != nil {
		log.Println("Failed to upgrade connection to WebSocket:", err)
		return
	}

	id := ws.genUID()

	ws.registerClient(id, conn)

	go ws.handleClientConn(id, conn)
}

// registerClient is a method to register new client connection
func (ws *WebSocketServer) registerClient(id string, conn *websocket.Conn) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	// initialize new client
	client := &Client{
		ID:   id,
		Conn: conn,
	}
	ws.clients = append(ws.clients, client)

	log.Printf("New client connected ID: %s, Remote addres: %s\n", client.ID, conn.RemoteAddr().String())
	log.Printf("Current connections: %d\n", len(ws.clients))
}

// unregisterClient is a method to unregister client
func (ws *WebSocketServer) unregisterClient(id string) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	idx := ws.getIndexByID(id)
	ws.clients = append(ws.clients[:idx], ws.clients[idx+1:]...)
}

// getIndexByID is a method to get index
func (ws *WebSocketServer) getIndexByID(id string) int {
	for i, v := range ws.clients {
		if id == v.ID {
			return i
		}
	}

	return 0
}

// genUID is a method to generate uuid
func (ws *WebSocketServer) genUID() string {
	q := "client"
	return fmt.Sprintf("%s:%s", q, strconv.Itoa(len(ws.clients)+1))
}

// handleClientConn is a method to handle
// client connection
func (ws *WebSocketServer) handleClientConn(id string, conn *websocket.Conn) {
	defer func() {
		// Unregister the client when the connection is closed
		ws.unregisterClient(id)
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			switch {
			case websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway):
				log.Printf("client ID: %s disconnected, err: %v", id, err)
			case websocket.IsUnexpectedCloseError(err, websocket.CloseNoStatusReceived):
				log.Printf("connection already closed")
			default:
				log.Printf("error reading message from WebSocket :%v", err)
			}

			break
		}

		// Process the received message
		// Add your custom logic here
		log.Printf("Received message from client: %s, message: %s", id, string(message))

		// Write response back to the WebSocket client
		err = conn.WriteMessage(websocket.TextMessage, []byte("Received your message"))
		if err != nil {
			log.Println("Error writing message to WebSocket:", err)
			break
		}
	}
}
