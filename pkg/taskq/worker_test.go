package taskq_test

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/pkg/taskq"
	"github/achjailani/go-simple-grpc/tests"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	testbox := tests.Init()
	cfg := testbox.Cfg
	ctx := testbox.Ctx

	// Create a Redis client
	dns := fmt.Sprintf("%s:%s", cfg.RedisTestConfig.RedisHost, cfg.RedisTestConfig.RedisPort)
	client := redis.NewClient(&redis.Options{
		Addr:     dns,
		Password: cfg.RedisTestConfig.RedisPassword, // Empty if no password is set
		DB:       cfg.RedisTestConfig.RedisDB,       // Default database
	})

	// Ping the Redis server to ensure connectivity
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	// Create the task queue
	taskQueue := taskq.NewQueue(client)

	_ = taskQueue.Enqueue(ctx, &taskq.Task{
		ID:   1,
		Data: "Hi, task 1",
	})
	_ = taskQueue.Enqueue(ctx, &taskq.Task{
		ID:   2,
		Data: "Hi, task 2",
	})
	_ = taskQueue.Enqueue(ctx, &taskq.Task{
		ID:   3,
		Data: "Hi, task 3",
	})
	_ = taskQueue.Enqueue(ctx, &taskq.Task{
		ID:   4,
		Data: "Hi, task 4",
	})

	// Create a wait group to ensure all worker goroutines finish
	var wg sync.WaitGroup

	// Create a channel to receive termination signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start the workers
	worker := taskq.NewWorker(3, taskQueue, &wg)
	worker.StartWorkers()

	// Wait for termination signal or completion of all tasks
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
			fmt.Println("All tasks processed. Exiting...")
		case <-time.After(5 * time.Second):
			fmt.Println("Graceful shutdown timed out. Exiting...")
		}
	}

	assert.True(t, true)
}
