package logger_test

import (
	"github/achjailani/go-simple-grpc/pkg/logger"
	"testing"
)

func TestNew(t *testing.T) {
	loggr := logger.New(logger.NewConfig())
	loggr.Log.Info("hello there!")
}
