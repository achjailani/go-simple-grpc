package websocx

import "github.com/gorilla/websocket"

// Client is a struct
type Client struct {
	ID         string // unique identifier client connection
	RemoteAddr string
	Conn       *websocket.Conn
}
