package taskq_test

import (
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/pkg/taskq"
	"github/achjailani/go-simple-grpc/tests"
	"log"
	"testing"
)

func TestNewQueue(t *testing.T) {
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
	queue := taskq.NewQueue(client)
	_ = queue

	data := make(map[string]interface{})
	data["name"] = "John Doe"
	data["age"] = 30
	data["email"] = "johndoe@example.com"

	// generate 10mb data
	jsonStr, _ := generateJSONWithSize(data, 10*1024*1024)

	err = queue.Client.RPush(ctx, "hell_queue", jsonStr).Err()
	assert.NoError(t, err)

	r, err := queue.Client.BLPop(ctx, 0, "hell_queue").Result()

	assert.NoError(t, err)
	fmt.Println(r)
}

// generateJSONWithSize generates a JSON string with the specified approximate size
func generateJSONWithSize(data map[string]interface{}, size int) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	jsonStr := string(jsonBytes)
	for len(jsonStr) < size {
		jsonStr += jsonStr
	}

	// Truncate the JSON string to the specified size
	jsonStr = jsonStr[:size]

	return jsonStr, nil
}
