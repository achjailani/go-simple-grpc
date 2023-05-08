package cmd

import (
	"fmt"
	"github/achjailani/go-simple-grpc/grpc/server"
	"github/achjailani/go-simple-grpc/infrastructure/persistence"

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
			db, errConn := persistence.NewDBConnection(cmd.Dependency.Cfg.DBConfig)
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
