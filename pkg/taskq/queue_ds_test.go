package taskq_test

import (
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/pkg/taskq"
	"testing"
)

func TestNewQueueDS(t *testing.T) {
	max := 10
	queue := taskq.NewQueueDS(max)
	_ = queue.Push("go")
	_ = queue.Push("php")
	_ = queue.Push("java")
	_ = queue.Push("js")
	_ = queue.Push("ruby")
	_ = queue.Push("python")
	_ = queue.Push("rust")

	r, err := queue.Pop()

	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	assert.Equal(t, "go", r)
}
