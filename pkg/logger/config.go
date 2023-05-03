package logger

import (
	"go.uber.org/zap"
	"os"
)

// NewConfig is a constructor
func NewConfig() *zap.Config {
	env := os.Getenv("APP_ENV")

	cnf := &zap.Config{
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if env == "production" {
		cnf.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		cnf.Development = false
		cnf.Encoding = "json"
		cnf.EncoderConfig = zap.NewProductionEncoderConfig()
	} else {
		cnf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		cnf.Development = true
		cnf.Encoding = "console"
		cnf.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	return cnf
}
