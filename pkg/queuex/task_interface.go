package queuex

type TaskInterface interface {
	Execute(payload interface{}) error
}
