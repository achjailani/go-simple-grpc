package logger

import (
	"go.uber.org/zap"
	"log"
)

// Logger is a type
type Logger struct {
	Log *zap.SugaredLogger
}

// New is a constructor
func New(cnf *zap.Config) *Logger {
	lg, err := cnf.Build()
	if err != nil {
		log.Fatalf("err build config: %v", err)
	}
	
	logger := lg.Sugar()

	defer lg.Sync()

	return &Logger{
		Log: logger,
	}
}
