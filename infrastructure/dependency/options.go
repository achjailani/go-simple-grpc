package dependency

import (
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/domain/service"
	"github/achjailani/go-simple-grpc/grpc/client"
	"github/achjailani/go-simple-grpc/pkg/logger"
)

// Option is an interface
type Option interface {
	apply(*Dependency)
}

// optionFunc is a type
type optionFunc func(*Dependency)

// apply is a method implementation
func (f optionFunc) apply(dep *Dependency) {
	f(dep)
}

// WithConfig is an option function
func WithConfig(cnf *config.Config) Option {
	return optionFunc(func(dep *Dependency) {
		dep.Cfg = cnf
	})
}

// WithRepository is an option function
func WithRepository(repository *service.Repositories) Option {
	return optionFunc(func(dep *Dependency) {
		dep.Repo = repository
	})
}

// WithLogger is an option function
func WithLogger(loggr *logger.Logger) Option {
	return optionFunc(func(dep *Dependency) {
		dep.Logger = loggr
	})
}

// WithGRPCClient is an option function
func WithGRPCClient(gClient *client.GRPCClient) Option {
	return optionFunc(func(dep *Dependency) {
		dep.GRPCClient = gClient
	})
}
