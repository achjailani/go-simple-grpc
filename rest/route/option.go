package route

import (
	"github/achjailani/go-simple-grpc/infrastructure/dependency"
)

// Option return Router with RouterOption to fill up the dependencies
type Option func(*Router)

// WithDependency is an option
func WithDependency(dep *dependency.Dependency) Option {
	return func(r *Router) {
		r.Dependency = dep
	}
}
