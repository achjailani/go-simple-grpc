package event

import (
	"fmt"
	"sync"
	"time"
)

var (
	// dispatcher is a global variable
	dispatcher *Dispatcher
	// onceDispatcher is a variable to define once instance
	onceDispatcher sync.Once
)

// Dispatcher is a struct
type Dispatcher struct {
	data     []string // temporary example
	events   map[Name]Listener
	shutdown chan struct{}
	wg       *sync.WaitGroup
}

// NewDispatcher is constructor
func NewDispatcher(listener Listener, names ...Name) *Dispatcher {
	onceDispatcher.Do(func() {
		dispatcher = &Dispatcher{
			data:     make([]string, 0),
			events:   make(map[Name]Listener),
			shutdown: make(chan struct{}),
			wg:       &sync.WaitGroup{},
		}
	})

	// register listener & events
	dispatcher.register(listener, names...)

	return dispatcher
}

// dispatch is a method to set new event
func (dpc *Dispatcher) dispatch(name Name, param interface{}) error {
	if _, ok := dpc.events[name]; !ok {
		return fmt.Errorf("the '%s' event is not registered", name)
	}

	// Dispatching process should be here
	// let say, you're using redis queue mechanism, you need to perform
	// push operation here, and should be using sync process
	// the data below is just and example!
	dpc.data = append(dpc.data, param.(string))

	return nil
}

// consume is a method
func (dpc *Dispatcher) consume() {
	for event, _ := range dpc.events {
		go dpc.listen(event)
	}
}

// listen is a method
func (dpc *Dispatcher) listen(event Name) {
	defer dpc.wg.Done()
	for {
		select {
		case <-dpc.shutdown:
			fmt.Printf("Event %s shutting down\n", event)
			return
		default:
			// Listening process should be here
			// let say, you're using redis queue mechanism, you need to perform
			// pop operation here
			dpc.events[event].Listen(event)
			fmt.Printf("done here: %s\n", event)
			time.Sleep(1 * time.Second)
		}
	}
}
