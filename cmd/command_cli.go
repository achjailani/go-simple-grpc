package cmd

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github/achjailani/go-simple-grpc/grpc/server"
	"github/achjailani/go-simple-grpc/infrastructure/persistence"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/urfave/cli/v2"
)

// newGRPCServer is a method command cli to run grpc
func (cmd *Command) newGRPCServer() *cli.Command {
	return &cli.Command{
		Name:  "grpc:start",
		Usage: "A command to run gRPC server",
		Action: func(c *cli.Context) error {
			grpcServer := server.NewGRPCServer(
				cmd.Dependency.Cfg,
				cmd.Dependency.Repo,
				cmd.Dependency.Logger,
			)
			err := grpcServer.Run(cmd.Dependency.Cfg.GRPCPort)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

// newDBMigrate is a method command cli to run db migration
func (cmd *Command) newDBMigrate() *cli.Command {
	return &cli.Command{
		Name:  "db:migrate",
		Usage: "A command to run database migration",
		Action: func(c *cli.Context) error {
			db, errConn := persistence.NewDBConnection(cmd.Dependency.Cfg)
			if errConn != nil {
				return fmt.Errorf("unable to connect to database: %w", errConn)
			}

			err := persistence.AutoMigrate(db)
			if err != nil {
				return fmt.Errorf("cannot run auto migrate: %w", err)
			}

			return nil
		},
	}
}

func (cmd *Command) newWebsocketClient() *cli.Command {
	return &cli.Command{
		Name:  "websocket:client",
		Usage: "A command to run database migration",
		Action: func(c *cli.Context) error {
			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, os.Interrupt)

			u := url.URL{Scheme: "ws", Host: "localhost:8181", Path: "/ws"}
			log.Printf("Connecting to WebSocket server: %s", u.String())

			conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Fatal("Failed to connect to WebSocket server:", err)
			}
			defer conn.Close()

			done := make(chan struct{})

			go func() {
				defer close(done)
				for {
					_, message, err := conn.ReadMessage()
					if err != nil {
						log.Println("Failed to read message from WebSocket server:", err)
						return
					}
					log.Printf("Received message from server: %s", string(message))
				}
			}()

			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-done:
					return nil
				case <-interrupt:
					log.Println("Interrupt signal received, closing connection...")
					err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					if err != nil {
						log.Println("Failed to send close message to WebSocket server:", err)
						return nil
					}
					select {
					case <-done:
					case <-time.After(time.Second):
					}
					return nil
				case <-ticker.C:
					message := []byte("Hello, server!")
					err := conn.WriteMessage(websocket.TextMessage, message)
					if err != nil {
						log.Println("Failed to send message to WebSocket server:", err)
						return nil
					}
					log.Println("Sent message to server:", string(message))
				}
			}
			return nil
		},
	}
}
