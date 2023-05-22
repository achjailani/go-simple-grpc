package taskq

import "fmt"

// QueueDS is a method
type QueueDS struct {
	max  int
	data []string
}

// NewQueueDS is a constructor
func NewQueueDS(max int) *QueueDS {
	return &QueueDS{
		max: max,
	}
}

// IsFull is a method to check if queue is full
func (q *QueueDS) IsFull() bool {
	return len(q.data) >= q.max
}

// IsEmpty is a method to check if queue is empty
func (q *QueueDS) IsEmpty() bool {
	return len(q.data) < 1
}

// Push / Enqueue is a method to push element
func (q *QueueDS) Push(str string) error {
	if q.IsFull() {
		return fmt.Errorf("queue is full")
	}

	q.data = append(q.data, str)

	return nil
}

// Pop is a method retrieve first item and remove it
func (q *QueueDS) Pop() (string, error) {
	if q.IsEmpty() {
		return "", fmt.Errorf("queue is empty")
	}

	r := q.data[0]
	cp := q.data[1:]
	q.data = cp

	return r, nil
}

// List is a method to retrieve all
func (q *QueueDS) List() []string {
	return q.data
}
