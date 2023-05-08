package cmd

import (
	"github/achjailani/go-simple-grpc/infrastructure/dependency"
)

// Option is an option type
type Option func(c *Command)

// WithDependency is a function option
func WithDependency(dep *dependency.Dependency) Option {
	return func(c *Command) {
		c.Dependency = dep
	}
}
