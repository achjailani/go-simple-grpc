package event

import (
	"context"
)

// Name is a string type for event name
type Name string

// Event is an event interface
type Event interface {
	Handle(ctx context.Context)
}
