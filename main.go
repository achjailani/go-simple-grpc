package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"github/achjailani/go-simple-grpc/client"
	"github/achjailani/go-simple-grpc/cmd"
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/domain/service"
	"github/achjailani/go-simple-grpc/infrastructure/persistence"
	"github/achjailani/go-simple-grpc/rest"
	"github/achjailani/go-simple-grpc/rest/route"
	"log"
	"os"
	"strconv"
	"time"
)

// main is a main function
func main() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.New()

	db, errConn := persistence.NewDBConnection(conf.DBConfig)
	if errConn != nil {
		log.Fatalf("unable connect to database, %v", errConn)
	}

	repo := service.NewDBService(db)

	command := cmd.NewCommand(
		cmd.WithConfig(conf),
		cmd.WithRepo(repo),
	)

	app := cmd.NewCLI()
	app.Commands = command.Build()

	clientConn, errClient := client.NewGRPCConn(conf)
	if errClient != nil {
		log.Fatalf("grpc client unable connect to server, %v", errClient)
	}

	grpcClient := client.NewGRPCClient(clientConn)
	app.Action = func(ctx *cli.Context) error {
		router := route.NewRouter(
			route.WithConfig(conf),
			route.WithRepository(repo),
			route.WithGRPCClient(grpcClient),
		).Init()

		shutdownTimeout := 10 * time.Second

		err := rest.RunHTTPServer(router, strconv.Itoa(conf.AppPort), shutdownTimeout)
		if err != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to run CLI command, err: %v", err)
	}
}
