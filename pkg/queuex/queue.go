package queuex

// QueueName is a type
type QueueName string

const (
	// RequestLogEvent is a constant
	RequestLogEvent = QueueName("request_log_event")
	// ActivityLogEvent is a constant
	ActivityLogEvent = QueueName("activity_log_event")
)
