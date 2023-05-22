package taskq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const (
	// queueName is a queue name
	queueName = "task_queue"
	// dequeueTimeout is a default queue timeout
	dequeueTimeout = 0
)

// Queue is a struct
type Queue struct {
	client *redis.Client
}

// NewQueue is a constructor
func NewQueue(client *redis.Client) *Queue {
	return &Queue{
		client: client,
	}
}

// Enqueue is a method to push item to list
func (q *Queue) Enqueue(ctx context.Context, task *Task) error {
	v, _ := json.Marshal(task)
	err := q.client.RPush(ctx, queueName, string(v)).Err()
	if err != nil {
		return fmt.Errorf("failed to enqueue: %w", err)
	}

	return nil
}

// Dequeue is a method to pop item from list
func (q *Queue) Dequeue(ctx context.Context) (*Task, error) {
	r, err := q.client.BLPop(ctx, dequeueTimeout, queueName).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to dequeue: %w", err)
	}

	if len(r) < 2 {
		return nil, fmt.Errorf("invalid task format")
	}

	var task Task
	err = json.Unmarshal([]byte(r[1]), &task)
	if err != nil {
		return nil, fmt.Errorf("error parsing task: %w", err)
	}

	return &task, nil
}
