package event

import "context"

// Name is a type
type Name string

// Event is a contract
type Event interface {
	Handle(ctx context.Context) error
}
