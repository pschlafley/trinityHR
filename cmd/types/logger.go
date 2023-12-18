package types

import (
	"log"
	"os"

	"github.com/gofor-little/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	NewProductionLog() (*zap.Logger, error)
	NewDevelopmentLog() (*zap.Logger, error)
}

func InitializeLogger() *zap.Logger {
	if err := env.Load(".env"); err != nil {
		log.Fatal(err)
	}

	environment := env.Get("ENVIRONMENT", "")

	var checkIfDevEnv bool

	if environment == "production" {
		checkIfDevEnv = true
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       checkIfDevEnv,
		DisableCaller:     checkIfDevEnv,
		DisableStacktrace: checkIfDevEnv,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	logger := zap.Must(config.Build())

	defer logger.Sync()

	return logger
}

func NewProductionLog() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	defer logger.Sync()

	return logger, err
}

func NewDevelopmentLog() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()

	return logger, err
}
