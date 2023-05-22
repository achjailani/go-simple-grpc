package taskq

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Worker is a struct
type Worker struct {
	queue    *Queue          // queue used by worker
	shutdown chan struct{}   // channel to shut down
	wg       *sync.WaitGroup // waitGroup to wait all goroutines done
	workers  int
}

// NewWorker is a constructor
func NewWorker(workers int, queue *Queue, wg *sync.WaitGroup) *Worker {
	return &Worker{
		workers:  workers,
		queue:    queue,
		shutdown: make(chan struct{}),
		wg:       wg,
	}
}

// Stop is a method
func (w *Worker) stop() {
	close(w.shutdown)
}

// run is a method
func (w *Worker) run(worker int) {
	defer w.wg.Done()
	for {
		select {
		case <-w.shutdown:
			fmt.Printf("Worker %d shutting down\n", worker)
			return
		default:
			task, err := w.queue.Dequeue(context.Background())
			if err != nil {
				log.Println("Failed to dequeue task:", err)
				continue
			}

			w.process(worker, task)
		}
	}
}

// process is a method to process the job
func (w *Worker) process(worker int, task *Task) {
	fmt.Printf("Worker %d processing task (ID: %d, Data: %s)\n", worker, task.ID, task.Data)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d completed task (ID: %d, Data: %s)\n", worker, task.ID, task.Data)
}

// StartWorkers is a method to run all workers
func (w *Worker) StartWorkers() {
	for i := 0; i < w.workers; i++ {
		w.wg.Add(1)
		go w.run(i + 1)
	}

	w.wg.Wait()
	w.stop()
}
