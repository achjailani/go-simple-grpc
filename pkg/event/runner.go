package event

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run is a method to start listening event
func (dpc *Dispatcher) Run() {
	// channel for termination signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// start consuming event
	dpc.consume()

	// listening signal
	select {
	case <-sigChan:
		// Termination signal received, initiate graceful shutdown
		fmt.Println("Termination signal received. Initiating graceful shutdown...")

		closeSigChan := make(chan struct{})
		go func() {
			close(closeSigChan)
		}()

		select {
		case <-closeSigChan:
			fmt.Println("dispatchers processed, existing...")
		case <-time.After(5 * time.Second):
			fmt.Println("Graceful shutdown timed out. Exiting...")
		}
	}
}
