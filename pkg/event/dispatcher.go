package event

import (
	"fmt"
	"sync"
)

// Dispatcher is a type
type Dispatcher struct {
	jobs     chan job
	events   map[Name]Listener
	shutdown chan struct{}   // channel to shut down
	wg       *sync.WaitGroup // waitGroup to wait all goroutines done
	workers  int
}

// NewDispatcher is a constructor
func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		jobs:   make(chan job),
		events: make(map[Name]Listener),
	}

	go d.consume()

	return d
}

// Register is a method
func (d *Dispatcher) Register(listener Listener, names ...Name) error {
	for _, name := range names {
		if _, ok := d.events[name]; ok {
			return fmt.Errorf("the '%s' event is already registered", name)
		}

		d.events[name] = listener
	}

	return nil
}

// Dispatch is method
func (d *Dispatcher) Dispatch(name Name, event interface{}) error {
	if _, ok := d.events[name]; !ok {
		return fmt.Errorf("the '%s' event is not registered", name)
	}

	d.jobs <- job{eventName: name, eventType: event}

	return nil
}

// consume is a method
func (d *Dispatcher) consume() {
	for j := range d.jobs {
		d.events[j.eventName].Listen(j.eventType)
	}
}
