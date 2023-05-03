package logger

import (
	"go.uber.org/zap"
	"os"
)

// NewConfig is a constructor
func NewConfig() *zap.Config {
	env := os.Getenv("APP_ENV")

	cnf := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if env == "production" {
		cnf.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		cnf.Development = false
		cnf.Encoding = "json"
		cnf.EncoderConfig = zap.NewProductionEncoderConfig()
		cnf.OutputPaths = []string{"stderr"}
		cnf.ErrorOutputPaths = []string{"stderr"}
	}

	return cnf
}
